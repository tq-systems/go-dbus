/*
 * Serial interface config RPC - marshall.go
 * Copyright (c) 2019, TQ-Systems GmbH
 * All rights reserved. For further information see LICENSE.txt
 */

package serial

import (
	"encoding/json"
	"errors"
)

// MarshalJSON is the custom marshalling implementation for BindState
func (s BindState) MarshalJSON() ([]byte, error) {
	var str string

	switch s {
	case BindUnbound:
		str = "unbound"
	case BindBound:
		str = "bound"
	case BindUnavailable:
		str = "unavailable"
	default:
		return nil, errors.New("invalid bind state")
	}

	return json.Marshal(str)
}

// UnmarshalJSON is the custom unmarshalling implementation for BindState
func (s *BindState) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	switch str {
	case "unbound":
		*s = BindUnbound
	case "bound":
		*s = BindBound
	case "unavailable":
		*s = BindUnavailable
	default:
		return errors.New("invalid bind state")
	}

	return nil
}

// MarshalJSON is the custom marshalling implementation for Parity
func (p Parity) MarshalJSON() ([]byte, error) {
	var str string

	switch p {
	case ParityNone:
		str = "none"
	case ParityEven:
		str = "even"
	case ParityOdd:
		str = "odd"
	default:
		return nil, errors.New("invalid parity")
	}

	return json.Marshal(str)
}

// UnmarshalJSON is the custom unmarshalling implementation for Parity
func (p *Parity) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	switch str {
	case "none":
		*p = ParityNone
	case "even":
		*p = ParityEven
	case "odd":
		*p = ParityOdd
	default:
		return errors.New("invalid parity")
	}

	return nil
}
