package types

import (
	"encoding/json"
)

func (enc OrderWrap) MarshalJson() ([]byte,error) {
	return json.Marshal(enc)
}

func (ord *OrderWrap) UnMarshalJson(input []byte) error {
	return json.Unmarshal(input, ord)

}