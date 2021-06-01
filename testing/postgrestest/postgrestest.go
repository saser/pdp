package postgrestest

import (
	"context"
	"database/sql"
	"net/url"
	"testing"

	"github.com/Saser/pdp/postgres"
	"github.com/cenkalti/backoff/v4"
	"github.com/ory/dockertest/v3"
)

const (
	repository = "postgres"
	tag        = "13.3"

	user        = "postgrestest-user"
	userKey     = "POSTGRES_USER"
	password    = "postgrestest-password"
	passwordKey = "POSTGRES_PASSWORD"
	dbname      = "postgrestest-dbname"
	dbnameKey   = "POSTGRES_DB"
)

var (
	env = map[string]string{
		userKey:     user,
		passwordKey: password,
		dbnameKey:   dbname,
	}
)

func NewInDocker(tb testing.TB) *postgres.Database {
	tb.Helper()

	// Set up the dockertest pool.
	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf(`dockertest.NewPool("") err = %v; want nil`, err)
	}

	// Start the container, and arrange for its cleanup.
	var envStrings []string
	for k, v := range env {
		envStrings = append(envStrings, k+"="+v)
	}
	res, err := pool.Run(repository, tag, envStrings)
	if err != nil {
		tb.Fatalf("pool.Run(%q, %q, %q) err = %v; want nil", repository, tag, envStrings, envStrings)
	}
	tb.Cleanup(func() {
		if err := pool.Purge(res); err != nil {
			tb.Errorf("pool.Purge(%v) = %v; want nil", res, err)
		}
	})

	// Fetch the host-port combination from the container.
	portID := "5432/tcp"
	host := res.GetHostPort(portID)
	if host == "" {
		tb.Fatalf(`res.GetHostPort(%q) = ""; want non-empty`, portID)
	}

	// Build up the connection string as an URI.
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     host,
		Path:     url.PathEscape(dbname),
		RawQuery: (url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	connString := u.String()

	// We're ready to try and open a connection to the database in
	// the container.
	db, err := postgres.New(connString)
	if err != nil {
		tb.Errorf("postgres.New(%q) err = %v; want nil", connString, err)
	}
	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Errorf("db.Close() = %v; want nil", err)
		}
	})
	return db
}

func BackoffGet(ctx context.Context, db *postgres.Database) (*sql.DB, error) {
	var sqldb *sql.DB
	op := func() error {
		var err error
		sqldb, err = db.Get(ctx)
		return err
	}
	bo := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	if err := backoff.Retry(op, bo); err != nil {
		return nil, err
	}
	return sqldb, nil
}
