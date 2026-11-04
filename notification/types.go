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

package notification

import (
	"bytes"
	"encoding/json"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/notification/mock_notification.go -package=notification github.com/tq-systems/go-dbus/notification Client

// Client is the interface for a notification client that can send over dbus
type Client interface {
	Send(Severity, string, string) error
	GetPerformance() (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, error)
	SendServiceLog(string, string) error
}

// Severity indicates the severity of an error
type Severity int

const (
	// Info is the lowest severity
	Info Severity = iota
	// Warning is the middle severity
	Warning
	// Error is the highest severity
	Error
)

var toString = map[Severity]string{
	Info:    "info",
	Warning: "warning",
	Error:   "error",
}

var toID = map[string]Severity{
	"info":    Info,
	"warning": Warning,
	"error":   Error,
}

// MarshalJSON provides string representations for json export of the Severity type
func (s Severity) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON converts string representations into Severity type
func (s *Severity) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*s = toID[j]
	return nil
}
