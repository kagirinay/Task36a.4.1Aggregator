package conf

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	t.Log(config)
}
