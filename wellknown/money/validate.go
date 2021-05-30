package money

import (
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/type/money"
)

var (
	ErrOutOfRange = errors.New("money: out of range")
)

type CurrencyCodeError struct {
	CurrencyCode string
}

func (e *CurrencyCodeError) Error() string {
	return fmt.Sprintf(`invalid currency code %q, want %q`, e.CurrencyCode, "SEK")
}

type SignsError struct {
	UnitsPositive bool
	NanosPositive bool
}

func (e *SignsError) Error() string {
	unitsSign := "negative"
	if e.UnitsPositive {
		unitsSign = "positive"
	}
	nanosSign := "negative"
	if e.NanosPositive {
		nanosSign = "positive"
	}
	return fmt.Sprintf("mismatched signs: `units` is %s and `nanos` is %s", unitsSign, nanosSign)
}

func Validate(m *money.Money) error {
	if got, want := m.GetCurrencyCode(), "SEK"; got != want {
		return &CurrencyCodeError{
			CurrencyCode: got,
		}
	}
	units := m.GetUnits()
	nanos := m.GetNanos()
	if (units > 0 && nanos < 0) || (units < 0 && nanos > 0) {
		return &SignsError{
			UnitsPositive: units > 0,
			NanosPositive: nanos > 0,
		}
	}
	if nanos > +999_999_999 || nanos < -999_999_999 {
		return ErrOutOfRange
	}
	return nil
}
