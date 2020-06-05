// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	return func() {
		server.Close()
	}
}

func TestGetTenantID(t *testing.T) {

	tearDown := setup()
	defer tearDown()

	expectedTenantID := "96fe9d1-6171-40aa-945b-4c64b63bf655"
	mux.HandleFunc("/subscriptions/foobarsubscription", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `authorization_uri="https://login.windows.net/`+expectedTenantID+`"`)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("Unauthorized"))
	})

	tenantID, err := GetTenantID(server.URL, "foobarsubscription")

	if err != nil {
		t.Error("Did not expect error")
	}

	if tenantID != expectedTenantID {
		t.Errorf("expected tenant Id : %s, but got %s", expectedTenantID, tenantID)
	}
}

func TestGetTenantID_UnexpectedResponse(t *testing.T) {

	tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/subscriptions/foobarsubscription", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := GetTenantID(server.URL, "foobarsubscription")

	expectedMsg := "Unexpected response from Get Subscription: 400"

	if err == nil || err.Error() != expectedMsg {
		t.Errorf("expected error with msg : %s to be thrown", expectedMsg)
	}
}

func TestGetTenantID_InvalidHeader(t *testing.T) {

	tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/subscriptions/foobarsubscription", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("fookey", "bazvalue")
	})

	_, err := GetTenantID(server.URL, "foobarsubscription")

	expectedMsg := "Header WWW-Authenticate not found in Get Subscription response"

	if err == nil || err.Error() != expectedMsg {
		t.Errorf("expected error with msg : %s to be thrown", expectedMsg)
	}
}

func TestGetTenantID_InvalidHeaderValue(t *testing.T) {

	tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/subscriptions/foobarsubscription", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `sample_invalid_auth_uri`)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("Unauthorized"))
	})

	_, err := GetTenantID(server.URL, "foobarsubscription")

	expectedMsg := "Could not find the tenant ID in header: WWW-Authenticate \"sample_invalid_auth_uri\""

	if err == nil || err.Error() != expectedMsg {
		t.Errorf("expected error with msg : %s to be thrown", expectedMsg)
	}
}
