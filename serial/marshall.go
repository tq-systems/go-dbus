/*
 * Copyright (c) 2020-2026 TQ-Systems GmbH <license@tq-group.com>, D-82229
 * Seefeld, Germany. All rights reserved.
 * Author: Alexander Poegelt and the Energy Manager development team
 *
 * This software is licensed under the TQ-Systems Product Software License
 * Agreement Version 1.0.3 or any later version.
 * You can obtain a copy of the License Agreement in the TQS (TQ-Systems
 * Software Licenses) folder on the following website:
 * https://www.tq-group.com/en/support/downloads/tq-software-license-conditions/
 * In case of any license issues please contact license@tq-group.com.
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
