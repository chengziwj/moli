package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert.Equal(t, CheckPassword("abc1Axyz.com"), nil, "password is valid")
	assert.NotEqual(t, CheckPassword("aB1_"), nil, "password is valid")
	assert.NotEqual(t, CheckPassword("aaaaaa"), nil, "password is valid")
	assert.NotEqual(t, CheckPassword("aaBB"), nil, "password is valid")
	assert.NotEqual(t, CheckPassword("1234"), nil, "password is valid")
}

func TestHash(t *testing.T) {
	pwd := "abcD123_jkle"
	h, _ := Hash(pwd)
	assert.Equal(t, Verify(h, pwd), true, "password is wrong")
}
