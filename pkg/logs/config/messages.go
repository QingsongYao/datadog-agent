// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://wwm.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

import (
	"strings"
	"sync"
)

// Messages holds a message message
type Messages struct {
	messages map[string]string
	lock     *sync.Mutex
}

// NewMessages initialize Messages with the default values
func NewMessages() *Messages {
	return &Messages{
		messages: make(map[string]string),
		lock:     &sync.Mutex{},
	}
}

// AddMessage create a message
func (m *Messages) AddMessage(key string, message string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.messages[key] = message
}

// AddWarning create a warning
func (m *Messages) AddWarning(key string, warning string) {
	m.AddMessage("warning_"+key, warning)
}

// GetMessages returns all the messages for an array of keys
func (m *Messages) GetMessages() []string {
	m.lock.Lock()
	defer m.lock.Unlock()

	messages := make([]string, 0)
	for key, message := range m.messages {
		if !strings.HasPrefix(key, "warning_") {
			messages = append(messages, message)
		}
	}
	return messages
}

// GetWarnings returns all the warnings for an array of keys
func (m *Messages) GetWarnings() []string {
	m.lock.Lock()
	defer m.lock.Unlock()

	warnings := make([]string, 0)
	for key, warning := range m.messages {
		if strings.HasPrefix(key, "warning_") {
			warnings = append(warnings, warning)
		}
	}
	return warnings
}

// RemoveMessage marks a RaisedMessage as solved
func (m *Messages) RemoveMessage(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.messages, key)
}

// RemoveWarning marks a RaisedWarning as solved
func (m *Messages) RemoveWarning(key string) {
	m.RemoveMessage("warning_" + key)
}
