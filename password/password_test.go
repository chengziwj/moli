package password

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestPassword(t *testing.T) {
	assert.Equal(t, CheckPassword("abcAxyz.com"), nil, "password is valid")
}
