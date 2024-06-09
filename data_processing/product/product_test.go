package product

import "testing"

func TestNewProductError(t *testing.T) {
	productTests := []struct {
		name       string
		ean        int
		shortname  string
		pricecents int
		want       interface{}
	}{
		{"negative ean", -123, "a", 123, ErrNegativeEAN},
		{"empty short name", 123, "", 123, ErrEmptyShortName},
		{"zero ean", 0, "a", 123, ErrZeroEAN},
	}

	for _, tt := range productTests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := NewProduct(tt.ean, tt.shortname, tt.pricecents)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestPriceEuro(t *testing.T) {
	priceTests := []struct {
		product Product
		want    float64
	}{
		{Product{123, "a", 123}, 1.23},
	}

	for _, tt := range priceTests {
		got := tt.product.PriceEuro()

		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
