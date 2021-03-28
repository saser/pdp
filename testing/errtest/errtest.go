package errtest

import (
	"strings"
	"testing"
)

type TestFunc func(tb testing.TB, err error)

func All(fs ...TestFunc) TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		for _, f := range fs {
			f(tb, err)
		}
	}
}

func IsNil() TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		if err != nil {
			tb.Errorf("err = %v; want nil", err)
		}
	}
}

func IsNonNil() TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		if err == nil {
			tb.Errorf("err = %v; want non-nil", err)
		}
	}
}

func ErrorContains(s string) TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		if !strings.Contains(err.Error(), s) {
			tb.Errorf("err.Error() = %q; want it to contain %q", err.Error(), s)
		}
	}
}
