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

package dbus

import (
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/dbus/mock_dbus.go -package=dbus github.com/tq-systems/go-dbus/dbus Service

// A Service provides a simple interface to make interfaces available over D-Bus
type Service interface {
	GetConnection() *dbus.Conn
	Export(pathName string, interfaceName string, obj interface{})
	Serve() error
	Stop() error
}

type exportIface struct {
	name string
	obj  interface{}
}

type service struct {
	conn *dbus.Conn

	exports map[string][]exportIface
}

// NewService opens a D-Bus system connection and registers under a given service name
func NewService(serviceName string) (Service, error) {
	conn, err := dbus.SystemBusPrivate()
	if err != nil {
		return nil, err
	}

	if err = conn.Auth(nil); err != nil {
		_ = conn.Close()
		return nil, err
	}

	if err = conn.Hello(); err != nil {
		_ = conn.Close()
		return nil, err
	}

	reply, err := conn.RequestName(serviceName,
		dbus.NameFlagDoNotQueue)
	if err != nil {
		return nil, err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return nil, errors.New("name already taken")
	}

	return &service{
		conn:    conn,
		exports: make(map[string][]exportIface),
	}, nil
}

func (srv *service) GetConnection() *dbus.Conn {
	return srv.conn
}

// Export exports an object as a D-Bus interface, providing all necessary introspection information
//
// All public methods which have a *dbus.Error as their last return value will be exported.
func (srv *service) Export(pathName string, interfaceName string, obj interface{}) {
	exports := srv.exports[pathName]

	exports = append(exports, exportIface{
		name: interfaceName,
		obj:  obj,
	})

	srv.exports[pathName] = exports
}

func (srv *service) Serve() error {
	for pathName, exports := range srv.exports {
		ifaces := []introspect.Interface{
			introspect.IntrospectData,
		}

		for i := range exports {
			iface := &exports[i]

			introspectInterface := introspect.Interface{
				Name:    iface.name,
				Methods: introspect.Methods(iface.obj),
			}
			ifaces = append(ifaces, introspectInterface)

			err := srv.conn.Export(iface.obj, dbus.ObjectPath(pathName), iface.name)
			if err != nil {
				return err
			}
		}

		introspectNode := introspect.Node{
			Name:       pathName,
			Interfaces: ifaces,
		}

		err := srv.conn.Export(introspect.NewIntrospectable(&introspectNode), dbus.ObjectPath(pathName),
			"org.freedesktop.DBus.Introspectable")
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop may be called in a defer after instantiating the dbus server
// improving the shut down process of programs using the dbus server
func (srv *service) Stop() error {
	return srv.conn.Close()
}
