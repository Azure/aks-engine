//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package event

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

// List contains zero or more events.
type List struct {
	Events []Event `json:"items"`
}

// Event represents the basic fields of a Kubernetes event.
type Event struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Type    string `json:"type"`
}

// GetAll returns all events.
func GetAll() (*List, error) {
	cmd := exec.Command("k", "get", "events", "-o", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error getting event:\n")
		util.PrintCommand(cmd)
		return nil, err
	}
	hl := List{}
	if err = json.Unmarshal(out, &hl); err != nil {
		log.Printf("Error unmarshalling events json:%s\n", err)
		return nil, err
	}
	return &hl, nil
}

// GetResult is a return struct for GetAsync.
type GetResult struct {
	Event *Event
	Err   error
}

// GetAsync wraps Get with a struct response for goroutine + channel usage.
func GetAsync(reason string) GetResult {
	event, err := Get(reason)
	return GetResult{
		Event: event,
		Err:   err,
	}
}

// Get returns an event with the given message.
func Get(message string) (*Event, error) {
	list, err := GetAll()
	if err != nil {
		return nil, err
	}
	for _, event := range list.Events {
		if event.Message == message {
			return &event, nil
		}
	}
	return nil, errors.Errorf("Event with message %s not found", message)
}

// GetWithRetry gets an event, retrying on failure.
func GetWithRetry(reason string, sleep, timeout time.Duration) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetResult)
	var mostRecentGetWithRetryError error
	var event *Event
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetAsync(reason)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.Err
			event = result.Event
			if mostRecentGetWithRetryError == nil {
				if event != nil {
					return event, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}
