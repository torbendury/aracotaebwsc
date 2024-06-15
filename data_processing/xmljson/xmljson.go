package xmljson

import (
	"encoding/json"
	"encoding/xml"
)

func XmlToStruct(in []byte, out interface{}) error {
	err := xml.Unmarshal(in, out)
	if err != nil {
		return err
	}
	return nil
}

func XmlToJson(in []byte, i interface{}) ([]byte, error) {
	err := XmlToStruct(in, i)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return res, nil
}
