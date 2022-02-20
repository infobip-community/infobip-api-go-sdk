package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetupValidation()
	os.Exit(m.Run())
}
