// Product contains the business entity of a product, specifying a contract
// which source as well as target systems must fulfil.
package product

import "errors"

var ErrNegativeEAN = errors.New("negative product EAN")
var ErrEmptyShortName = errors.New("empty product shortname")
var ErrZeroEAN = errors.New("zero EAN")

type Product struct {
	EAN        int    `xml:"EAN" json:"EAN"`
	ShortName  string `xml:"ShortName" json:"shortName"`
	PriceCents int    `xml:"PriceCents" json:"priceCents"`
}

type Stock int

func New(ean int, shortname string, pricecents int) (*Product, error) {
	if ean < 0 {
		return nil, ErrNegativeEAN
	}
	if ean == 0 {
		return nil, ErrZeroEAN
	}
	if shortname == "" {
		return nil, ErrEmptyShortName
	}

	return &Product{
		EAN:        ean,
		ShortName:  shortname,
		PriceCents: pricecents,
	}, nil
}

func (p *Product) PriceEuro() float64 {
	return float64(p.PriceCents) / 100
}
