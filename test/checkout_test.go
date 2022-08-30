package test

import (
	"testing"

	"github.com/mghifariyusuf/fita/cli/checkout"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	cases := []struct {
		name          string
		cart          []checkout.Cart
		expectedTotal float64
	}{
		{"MacBook Pro free Raspberry Pi B", []checkout.Cart{{"43N23P", 1}, {"234234", 1}}, 5399.99},
		{"Buy 3 Google Homes", []checkout.Cart{{"120P90", 3}}, 99.98},
		{"Discount Alexa Speakers", []checkout.Cart{{"A304SD", 3}}, 295.65},
	}

	inventory, err := checkout.LoadInventory()
	assert.NoError(t, err)

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			total := inventory.Checkout(tc.cart)
			assert.Equal(t, tc.expectedTotal, total)
		})
	}
}
