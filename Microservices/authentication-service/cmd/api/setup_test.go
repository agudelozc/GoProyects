package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	// setup
	code := m.Run()
	repo := data.NewPostgresRepository(nil)
	testApp.Repo = repo
	// teardown
	os.Exit(code)
}
