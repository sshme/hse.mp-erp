package di

import (
	"testing"
)

func TestInitializeApp(t *testing.T) {
	app, err := InitializeApp()
	if err != nil {
		t.Errorf("InitializeApp() failed with error: %v", err)
	}

	if app == nil {
		t.Error("InitializeApp() returned nil app")
		return
	}

	if app.Logger == nil {
		t.Error("Logger dependency was not initialized")
	}
	if app.VetClinic == nil {
		t.Error("VetClinic dependency was not initialized")
	}
	if app.Zoo == nil {
		t.Error("Zoo dependency was not initialized")
	}
}
