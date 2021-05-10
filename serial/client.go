/*
 * Serial interface config RPC - client.go
 * Copyright (c) 2019, TQ-Systems GmbH
 * All rights reserved. For further information see LICENSE.txt
 */

package serial

// Go get dbus-codegen-go:
//
//   go get -u github.com/tq-systems/go-dbus-codegen/cmd/dbus-codegen-go
//
// To get an updated com.tq_group.tq_em.device_settings1.xml after changes to the exported interface:
//
//   dbus-send --system --type=method_call --print-reply=literal \
//     --dest=com.tq_group.tq_em.device_settings1 \
//     /com/tq_group/tq_em/device_settings1 \
//     org.freedesktop.DBus.Introspectable.Introspect \
//     > com.tq_group.tq_em.device_settings1.xml
//
//go:generate sh -c "go get -u github.com/tq-systems/go-dbus-codegen/cmd/dbus-codegen-go"
//go:generate sh -c "dbus-codegen-go -prefix com.tq_group.tq_em.device_settings1 -package device_settings com.tq_group.tq_em.device_settings1.xml > device_settings/device_settings1.go"

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"

	"github.com/tq-systems/go-dbus/serial/device_settings"
)

const (
	serviceName = "com.tq_group.tq_em.device_settings1"
	pathName    = "/com/tq_group/tq_em/device_settings1"
)

type configHandler struct {
	serialControl *device_settings.SerialControl
	app           string
}

// NewClient returns a new Config client for the app with the given name
func NewClient(app string) (Client, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	serialControl := device_settings.NewSerialControl(
		conn.Object(serviceName, dbus.ObjectPath(pathName)),
	)

	return &configHandler{serialControl, app}, nil
}

func (h *configHandler) ListInterfaces() (map[string]BindState, error) {
	ifaces, err := h.serialControl.ListInterfaces(h.app)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]BindState)
	for iface, state := range ifaces {
		ret[iface] = BindState(state)
	}
	return ret, nil
}

func (h *configHandler) GetInterfaceRaw(name string) ([]byte, error) {
	jsonConfig, err := h.serialControl.GetInterface(h.app, name)
	if err != nil {
		return nil, err
	}

	return []byte(jsonConfig), nil
}

func (h *configHandler) GetInterface(name string, config interface{}) error {
	jsonConfig, err := h.GetInterfaceRaw(name)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonConfig, config)
}

func (h *configHandler) BindInterface(name string, config interface{}, rebind bool) error {
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return h.serialControl.BindInterface(h.app, name, string(jsonConfig), rebind)
}

func (h *configHandler) UnbindInterface(name string) error {
	return h.serialControl.UnbindInterface(h.app, name)
}

func (h *configHandler) Reset(name string) error {
	return h.serialControl.Reset(h.app, name)
}
func (h *configHandler) SetOperState(name string, state OperState) error {
	return h.serialControl.SetOperState(h.app, name, byte(state))
}
