package tool_test

import (
	"testing"

	"github.com/orvosi/api/internal/tool"
	"github.com/stretchr/testify/assert"
)

const (
	audience = "test-audience"
)

func TestNewIDTokenDecoder(t *testing.T) {
	t.Run("successfully create an instance of IDTokenDecoder", func(t *testing.T) {
		dec := tool.NewIDTokenDecoder(audience)
		assert.NotNil(t, dec)
	})
}
