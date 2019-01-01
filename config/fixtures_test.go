package config

import (
	"testing"
)

func TestParseCertificate(t *testing.T) {
	_, err := ParseCertificate()
	if err != nil {
		t.Errorf("ParseCertificate() error = %v", err)
	}
}
