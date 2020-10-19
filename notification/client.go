/*
 * utils - notification - client.go
 * Copyright (c) 2019, TQ-Systems GmbH
 * All rights reserved. For further information see LICENSE.
 */

package notification

// Go get dbus-codegen-go:
//
//   go get -u vgitlab01.tq-net.de/tq-em/tools/dbus-codegen-go.git/cmd/dbus-codegen-go
//
// To get an updated com.tq_group.tq_em.health_check1.xml after changes to the exported interface:
//
//   dbus-send --system --type=method_call --print-reply=literal \
//     --dest=com.tq_group.tq_em.health_check1 \
//     /com/tq_group/tq_em/health_check1 \
//     org.freedesktop.DBus.Introspectable.Introspect \
//     > com.tq_group.tq_em.health_check1.xml
//
//go:generate sh -c "go get -u vgitlab01.tq-net.de/tq-em/tools/dbus-codegen-go.git/cmd/dbus-codegen-go@switch-2-gomod"
//go:generate sh -c "dbus-codegen-go -prefix com.tq_group.tq_em.health_check1 -package health_check com.tq_group.tq_em.health_check1.xml > health_check/health_check1.go"

import (
	"sync"
	"time"

	"github.com/godbus/dbus/v5"

	"vgitlab01.tq-net.de/tq-em/libs/dbus.git/notification/health_check"
)

const (
	serviceName = "com.tq_group.tq_em.health_check1"
	pathName    = "/com/tq_group/tq_em/health_check1"
)

type client struct {
	notification *health_check.Notification
	performance  *health_check.Performance
	service      *health_check.Service
	app          string
	mutex        *sync.Mutex
}

// NewClient returns a client that can send notifications over the dbus interface
func NewClient(app string) (Client, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	notification := health_check.NewNotification(
		conn.Object(serviceName, dbus.ObjectPath(pathName)),
	)

	performance := health_check.NewPerformance(
		conn.Object(serviceName, dbus.ObjectPath(pathName)),
	)

	service := health_check.NewService(
		conn.Object(serviceName, dbus.ObjectPath(pathName)),
	)

	mutex := &sync.Mutex{}

	return &client{notification, performance, service, app, mutex}, nil
}

func (c *client) Send(severity Severity, msgCode string, msgText string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.notification.Send(c.app, byte(severity), msgCode, msgText, time.Now().Unix())
}

func (c *client) GetPerformance() (out0 uint64, out1 uint64, out2 uint64, out3 uint64, out4 uint64, out5 uint64, out6 uint64, out7 uint64, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.performance.GetPerformance()
}

func (c *client) SendServiceLog(app string, message string) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.service.SendServiceLog(app, message)
}
