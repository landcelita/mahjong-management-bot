package tonpuorhanchan

import (
	"testing"
)

func TestNewTonpuOrHanchan(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			
		}
	})
}

func TestNewMoney(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount, err := decimal.NewFromString("100")
		if err != nil {
			t.Fatal(err)
		}
		currency := "JPY"
		got, err := NewMoney(amount, currency)
		if err != nil {
			t.Fatal(err)
		}
		want := &Money{amount: amount, currency: currency}
		if diff := cmp.Diff(got, want, cmp.AllowUnexported(Money{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
}