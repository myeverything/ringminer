package types

import (
	"encoding/json"
	"errors"
	"reflect"
	"math/big"
)

func (ord Order) MarshalJson() ([]byte,error) {
	type order struct {
		PeerId      string  `json:"peerId"`
		Id          string  `json:"id"`
		Protocol    string  `json:"protocol"`
		Owner       string  `json:"owner"`
		OutToken    string  `json:"outToken"`
		InToken     string  `json:"inToken"`
		OutAmount   uint64  `json:"outAmount"`
		InAmount    uint64  `json:"inAmount"`
		Expiration  uint64  `json:"expiration"`
		Fee         uint64  `json:"fee"`
		ShareRate int  `json:"shareRate"`
		V           uint8   `json:"v"`
		R           string  `json:"r"`
		S           string  `json:"s"`
	}

	var enc order
	// TODO(fukun): set the locate ipfs peerid
	enc.PeerId = ""
	//enc.Id = ord.Id.Str()
	enc.Protocol = ord.Protocol.Str()
	//enc.Owner = ord.Owner.Str()
	enc.OutToken = ord.TokenS.Str()
	enc.InToken = ord.TokenB.Str()
	enc.Expiration = ord.Expiration
	enc.Fee = ord.LrcFee.Uint64()
	enc.ShareRate = ord.SavingSharePercentage
	enc.V = ord.V
	enc.R = ord.R.Str()
	enc.S = ord.S.Str()
	return json.Marshal(enc)
}

func (ord *Order) UnMarshalJson(input []byte) error {
	type order struct {
		PeerId      string  `json:"peerId"`
		Id          string  `json:"id"`
		Protocol    string  `json:"protocol"`
		Owner       string  `json:"owner"`
		OutToken    string  `json:"outToken"`
		InToken     string  `json:"inToken"`
		OutAmount   uint64  `json:"outAmount"`
		InAmount    uint64  `json:"inAmount"`
		Expiration  uint64  `json:"expiration"`
		Fee         uint64  `json:"fee"`
		ShareRate int  `json:"shareRate"`
		V           uint8   `json:"v"`
		R           string  `json:"r"`
		S           string  `json:"s"`
	}

	var dec order
	err := json.Unmarshal(input, &dec)
	if err != nil {
		return err
	}

	// TODO(fukun): create order id
	//ord.Id = BytesToHash([]byte(""))

	if !reflect.ValueOf(dec.Protocol).IsValid() {
		return errors.New("missing required field 'Protocol' for order")
	}
	ord.Protocol = StringToAddress(dec.Protocol)

	if !reflect.ValueOf(dec.Owner).IsValid() {
		return errors.New("missing required field 'Owner' for order")
	}
	//ord.Owner = StringToAddress(dec.Owner)

	if !reflect.ValueOf(dec.OutToken).IsValid() {
		return errors.New("missing required field 'OutToken' for order")
	}
	ord.TokenS = StringToAddress(dec.OutToken)

	if !reflect.ValueOf(dec.InToken).IsValid() {
		return errors.New("missing required field 'InToken' for order")
	}
	ord.TokenB = StringToAddress(dec.InToken)

	if !reflect.ValueOf(dec.OutAmount).IsValid() {
		return errors.New("missing required field 'OutAmount' for order")
	}
	ord.AmountS = big.NewInt(int64(dec.OutAmount))

	if !reflect.ValueOf(dec.InAmount).IsValid() {
		return errors.New("missing required field 'InAmount' for order")
	}
	ord.AmountB = big.NewInt(int64(dec.InAmount))

	if !reflect.ValueOf(dec.Expiration).IsValid() {
		return errors.New("missing required field 'Expiration' for order")
	}
	ord.Expiration = dec.Expiration

	if !reflect.ValueOf(dec.Fee).IsValid() {
		return errors.New("missing required field 'Fee' for order")
	}
	ord.LrcFee = big.NewInt(int64(dec.Fee))

	if !reflect.ValueOf(dec.ShareRate).IsValid() {
		return errors.New("missing required field 'SavingShare' for order")
	}
	ord.SavingSharePercentage = int(dec.ShareRate)

	if !reflect.ValueOf(dec.V).IsValid() {
		return errors.New("missing required field 'ECDSA.V' for order")
	}
	ord.V = dec.V

	if !reflect.ValueOf(dec.S).IsValid() {
		return errors.New("missing required field 'ECDSA.S' for order")
	}
	ord.S = StringToSign(dec.S)

	if  !reflect.ValueOf(dec.R).IsValid() {
		return errors.New("missing required field 'ECSA.R' for order")
	}
	ord.R = StringToSign(dec.R)

	return nil
}