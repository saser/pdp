package money

import (
	"errors"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/type/money"
)

func asCurrencyCodeError(f func(tb testing.TB, e *CurrencyCodeError)) errtest.TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		var e *CurrencyCodeError
		if !errors.As(err, &e) {
			tb.Fatalf("errors.As(%v, %T) failed", err, &e)
		}
		f(tb, e)
	}
}

func asSignsError(f func(tb testing.TB, e *SignsError)) errtest.TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		var e *SignsError
		if !errors.As(err, &e) {
			tb.Fatalf("errors.As(%v, %T) failed", err, &e)
		}
		f(tb, e)
	}
}

func TestValidate(t *testing.T) {
	for _, m := range []*money.Money{
		{
			CurrencyCode: "SEK",
			Units:        0,
			Nanos:        0,
		},
		{
			CurrencyCode: "SEK",
			Units:        100,
			Nanos:        0,
		},
		{
			CurrencyCode: "SEK",
			Units:        -100,
			Nanos:        0,
		},
		{
			CurrencyCode: "SEK",
			Units:        0,
			Nanos:        100,
		},
		{
			CurrencyCode: "SEK",
			Units:        0,
			Nanos:        -100,
		},
		{
			CurrencyCode: "SEK",
			Units:        -100,
			Nanos:        -100,
		},
		{
			CurrencyCode: "SEK",
			Units:        100,
			Nanos:        100,
		},
		{
			CurrencyCode: "SEK",
			Units:        0,
			Nanos:        999_999_999,
		},
		{
			CurrencyCode: "SEK",
			Units:        0,
			Nanos:        -999_999_999,
		},
	} {
		if err := Validate(m); err != nil {
			t.Errorf("Validate(%v) = %v; want nil", m, err)
		}
	}
}

func TestValidate_Errors(t *testing.T) {
	for _, tt := range []struct {
		name string
		m    *money.Money
		tf   errtest.TestFunc
	}{
		{
			name: "WrongCurrencyCode",
			m: &money.Money{
				CurrencyCode: "USD",
			},
			tf: asCurrencyCodeError(func(tb testing.TB, e *CurrencyCodeError) {
				tb.Helper()
				want := &CurrencyCodeError{
					CurrencyCode: "USD",
				}
				if diff := cmp.Diff(want, e); diff != "" {
					tb.Errorf("unexpected CurrencyCodeError (-want +got)\n%s", diff)
				}
			}),
		},
		{
			name: "MismatchedSigns_UnitsPositive_NanosNegative",
			m: &money.Money{
				CurrencyCode: "SEK",
				Units:        +10,
				Nanos:        -10,
			},
			tf: asSignsError(func(tb testing.TB, e *SignsError) {
				tb.Helper()
				want := &SignsError{
					UnitsPositive: true,
					NanosPositive: false,
				}
				if diff := cmp.Diff(want, e); diff != "" {
					tb.Errorf("unexpected SignsError (-want +got)\n%s", diff)
				}
			}),
		},
		{
			name: "MismatchedSigns_UnitsNegative_NanosPositive",
			m: &money.Money{
				CurrencyCode: "SEK",
				Units:        -10,
				Nanos:        +10,
			},
			tf: asSignsError(func(tb testing.TB, e *SignsError) {
				tb.Helper()
				want := &SignsError{
					UnitsPositive: false,
					NanosPositive: true,
				}
				if diff := cmp.Diff(want, e); diff != "" {
					tb.Errorf("unexpected SignsError (-want +got)\n%s", diff)
				}
			}),
		},
		{
			name: "OutOfRange_PositiveNanos",
			m: &money.Money{
				CurrencyCode: "SEK",
				Nanos:        +1_000_000_000,
			},
			tf: errtest.Is(ErrOutOfRange),
		},
		{
			name: "OutOfRange_NegativeNanos",
			m: &money.Money{
				CurrencyCode: "SEK",
				Nanos:        -1_000_000_000,
			},
			tf: errtest.Is(ErrOutOfRange),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.m)
			if err == nil {
				t.Fatalf("Validate(%v) = nil; want non-nil", tt.m)
			}
			tt.tf(t, err)
		})
	}
}
