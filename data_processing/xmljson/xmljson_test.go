package xmljson

import (
	"strings"
	"testing"

	"github.com/torbendury/aracotaebwsc/data_processing/product"
)

func MustEqual(want interface{}, got interface{}, t *testing.T) {
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func MinifyString(in string) string {
	return strings.ReplaceAll(strings.ReplaceAll(in, "\n", ""), " ", "")
}

func TestXmlToStruct(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
<Product>
    <EAN>123</EAN>
    <ShortName>USB Plug</ShortName>
    <PriceCents>123</PriceCents>
</Product>`)

	want := &product.Product{
		EAN:        123,
		ShortName:  "USB Plug",
		PriceCents: 123,
	}

	got := &product.Product{}

	err := XmlToStruct(xmlData, got)

	MustEqual(nil, err, t)

	MustEqual(want.EAN, got.EAN, t)
	MustEqual(want.ShortName, got.ShortName, t)
	MustEqual(want.PriceCents, got.PriceCents, t)
}

func TestXmlToJson(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
<Product>
    <EAN>123</EAN>
    <ShortName>USB Plug</ShortName>
    <PriceCents>123</PriceCents>
</Product>`)

	want := []byte(`{
"EAN": 123,
"shortName": "USB Plug",
"priceCents": 123
}`)

	intermediateProduct := &product.Product{}

	got, err := XmlToJson(xmlData, intermediateProduct)

	MustEqual(nil, err, t)
	MustEqual(MinifyString(string(want[:])), MinifyString(string(got[:])), t)
}
