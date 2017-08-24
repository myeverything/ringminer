package types

import (
	"encoding/json"
	"errors"
	"reflect"
	"math/big"
)

func (ord Order) MarshalJson() ([]byte,error) {
	type order struct {
		PeerId      []byte  `json:"peerId"`
		Id          []byte  `json:"id"`
		Protocol    []byte  `json:"protocol"`
		Owner       []byte  `json:"owner"`
		OutToken    []byte  `json:"outToken"`
		InToken     []byte  `json:"inToken"`
		OutAmount   uint64  `json:"outAmount"`
		InAmount    uint64  `json:"inAmount"`
		Expiration  uint64  `json:"expiration"`
		Fee         uint64  `json:"fee"`
		SavingShare uint64  `json:"savingShare"`
		V           uint8   `json:"v"`
		R           []byte  `json:"r"`
		S           []byte  `json:"s"`
	}

	var enc order
	// TODO(fukun): set the locate ipfs peerid
	enc.PeerId = []byte("")
	enc.Id = []byte(ord.Id)
	enc.Protocol = []byte(ord.Protocol)
	enc.Owner = []byte(ord.Owner)
	enc.OutToken = []byte(ord.OutToken)
	enc.InToken = []byte(ord.InToken)
	enc.Expiration = ord.Expiration
	enc.Fee = ord.Fee.Uint64()
	enc.SavingShare = ord.SavingShare.Uint64()
	enc.V = ord.V
	enc.R = []byte(ord.R)
	enc.S = []byte(ord.S)
	return json.Marshal(enc)
}

func (ord *Order) UnMarshalJson(input []byte) error {
	type order struct {
		PeerId      []byte  `json:"peerId"`
		Id          []byte  `json:"id"`
		Protocol    []byte  `json:"protocol"`
		Owner       []byte  `json:"owner"`
		OutToken    []byte  `json:"outToken"`
		InToken     []byte  `json:"inToken"`
		OutAmount   uint64  `json:"outAmount"`
		InAmount    uint64  `json:"inAmount"`
		Expiration  uint64  `json:"expiration"`
		Fee         uint64  `json:"fee"`
		SavingShare uint64  `json:"savingShare"`
		V           uint8   `json:"v"`
		R           []byte  `json:"r"`
		S           []byte  `json:"s"`
	}

	var dec order
	err := json.Unmarshal(input, &dec)
	if err != nil {
		return err
	}

	// TODO(fukun): create order id
	ord.Id = Hash("")

	if dec.Protocol == nil {
		return errors.New("missing required field 'Protocol' for order")
	}
	ord.Protocol = Address(dec.Protocol)

	if dec.Owner == nil {
		return errors.New("missing required field 'Owner' for order")
	}
	ord.Owner = Address(dec.Owner)

	if dec.OutToken == nil {
		return errors.New("missing required field 'OutToken' for order")
	}
	ord.OutToken = Address(dec.OutToken)

	if dec.InToken == nil {
		return errors.New("missing required field 'InToken' for order")
	}
	ord.InToken = Address(dec.InToken)

	if !reflect.ValueOf(dec.OutAmount).IsValid() {
		return errors.New("missing required field 'OutAmount' for order")
	}
	ord.OutAmount = big.NewInt(int64(dec.OutAmount))

	if !reflect.ValueOf(dec.InAmount).IsValid() {
		return errors.New("missing required field 'InAmount' for order")
	}
	ord.InAmount = big.NewInt(int64(dec.InAmount))

	if !reflect.ValueOf(dec.Expiration).IsValid() {
		return errors.New("missing required field 'Expiration' for order")
	}
	ord.Expiration = dec.Expiration

	if !reflect.ValueOf(dec.Fee).IsValid() {
		return errors.New("missing required field 'Fee' for order")
	}
	ord.Fee = big.NewInt(int64(dec.Fee))

	if !reflect.ValueOf(dec.SavingShare).IsValid() {
		return errors.New("missing required field 'SavingShare' for order")
	}
	ord.SavingShare = big.NewInt(int64(dec.SavingShare))

	if !reflect.ValueOf(dec.V).IsValid() {
		return errors.New("missing required field 'ECDSA.V' for order")
	}
	ord.V = dec.V

	if dec.S == nil {
		return errors.New("missing required field 'ECDSA.S' for order")
	}
	ord.S = Sign(dec.S)

	if dec.R == nil {
		return errors.New("missing required field 'ECSA.R' for order")
	}
	ord.R = Sign(dec.R)

	return nil
}