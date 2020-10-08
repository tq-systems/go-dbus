/*
 * Serial interface config RPC - types.go
 * Copyright (c) 2019, TQ-Systems GmbH
 * All rights reserved. For further information see LICENSE.txt
 */

package serial

// BindState describes the bind state of a serial interface (in the context of a specific app)
type BindState byte

const (
	// BindUnbound is the state of an interface that is currently unbound
	BindUnbound BindState = iota
	//BindBound is the state of an interface that is currently bound to this app
	BindBound
	//BindUnavailable is the state of an interface that is currently bound to another app
	BindUnavailable
)

// OperState describes the operating state of a serial interface
type OperState byte

const (
	// OperOff switches the serial LED off
	OperOff OperState = iota
	// OperIdle sets the serial LED to indicate that no communication is taking place
	OperIdle
	// OperComm sets the serial LED to indicate that communication is taking place over the serial interfaces
	OperComm
	// OperScan sets the serial LED to indicate that a sensor scan is in progress
	OperScan
	// OperTimeout sets the serial LED to show a timeout warning
	OperTimeout
	// OperError sets the serial LED to show an error
	OperError
)

// Parity is a serial interface parity parameter
type Parity byte

const (
	// ParityNone is no parity
	ParityNone Parity = iota
	// ParityEven is even parity
	ParityEven
	// ParityOdd is odd parity
	ParityOdd
)

// Client is the public interface of the serial config and state client
type Client interface {
	// ListInterfaces lists all existing serial interfaces, with the information
	// whether the interfaces are unbound (available), bound the this app, or
	// unavailable (bound to another app)
	ListInterfaces() (map[string]BindState, error)

	// GetInterfaceRaw returns the JSON-encoded configuration for the serial
	// interface `name` (for example "APP0") that was stored using BindInterface
	// Trying to retrieve the configuration for an interface that is not bound
	// to this app is an error.
	GetInterfaceRaw(name string) ([]byte, error)

	// GetInterface retrieves the configuration for the serial interface
	// `name` and unmarshalls it into the given interface.
	// Trying to retrieve the configuration for an interface that is not bound
	// to this app is an error.
	GetInterface(name string, config interface{}) error

	// BindInterface tries to bind the given interface `name` to this app,
	// optionally storing additional app-specific information (which will be
	// marshalled as JSON)
	// For the initial bind of an interface, `rebind` must be false. To update
	// the stored configuration without unbinding, BindInterface can be called
	// again with `rebind` set to true.
	BindInterface(name string, config interface{}, rebind bool) error

	// UnbindInterface releases an interface bound using BindInterface.
	UnbindInterface(name string) error

	// Reset attempts to reset the RS485 interface `name`.
	//
	// Caveats:
	// - It is unspecified whether other interfaces are also reset (with our
	//   current hardware, resetting individual interfaces is not possible
	// - In the future we may to decide to refuse resetting the interfaces
	//   when the other interface is currently bound to another app
	Reset(name string) error

	// Sets the operating state of interface `name`, so the correct state
	// is signalled using the RS485 LED.
	SetOperState(name string, state OperState) error
}

// PortParameters contains the configurable serial interface parameters that are independent of the used protocol
type PortParameters struct {
	Baudrate uint32 `json:"baudrate"`
	Databits uint8  `json:"databits"`
	Parity   Parity `json:"parity"`
	Stopbits uint8  `json:"stopbits"`
}

// LegacyConfig contains app- and protocol-specific values that have been migrated from config version 0
type LegacyConfig struct {
	PresetID uint8  `json:"presetID,omitempty"`
	Address  uint8  `json:"address,omitempty"`
	BoundID  string `json:"boundID,omitempty"`
}

// SimpleConfig contains just the protocol-independent serial parameters
// Serial parameters from migrated v0 configs will be stored in this struct
type SimpleConfig struct {
	Serial *PortParameters `json:"serial,omitempty"`
}

// SimpleConfigLegacy extends SimpleConfig with protocol-specific legacy fields
// Legacy config from migrated v0 configs will be stored in this struct
type SimpleConfigLegacy struct {
	SimpleConfig
	Legacy *LegacyConfig `json:"legacy,omitempty"`
}
