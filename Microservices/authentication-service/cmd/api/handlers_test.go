package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authenticate(t *testing.T) {
	postBody := map[string]any{
		"email":    "me@here.com",
		"password": "verysecret",
	}
	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

}
func Test_Authenticate_InvalidJSON(t *testing.T) {
	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader([]byte("{invalid json")))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func Test_Authenticate_InvalidCredentials(t *testing.T) {
	postBody := map[string]any{
		"email":    "invalid@user.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func Test_Authenticate_ValidCredentials(t *testing.T) {
	postBody := map[string]any{
		"email":    "me@here.com",
		"password": "verysecret",
	}
	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

	var response jsonResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if response.Error {
		t.Errorf("expected error to be false, got true")
	}

	if response.Message != "Logged in user me@here.com" {
		t.Errorf("unexpected message: got %v want %v", response.Message, "Logged in user me@here.com")
	}
}
