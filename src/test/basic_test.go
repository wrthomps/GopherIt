package test

import (
	"lib"
	"testing"
)

func TestGetMessage(t *testing.T) {
	if lib.GetMessage() == "" {
		t.Error("Expected non-empty string")
	}
}
