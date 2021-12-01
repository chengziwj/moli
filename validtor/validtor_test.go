package validtor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMail(t *testing.T) {
	assert.Equal(t, Mail("abc@xyz.com"), true, "this is mail")
}

func TestUserName(t *testing.T) {
	assert.Equal(t, UserName("ABCd12"), true, "this is user name")
	assert.NotEqual(t, UserName("1ABCd"), true, "this is user name")
	assert.Equal(t, UserName("ABC_d"), true, "this is user name")
	assert.Equal(t, UserName("ABC_d"), true, "this is user name")
}

func TestIpv4(t *testing.T) {
	assert.Equal(t, Ipv4("127.0.0.255"), true, "this is not ipv4")
	assert.Equal(t, Ipv4("0.0.0.255"), true, "this is not ipv4")
}

func TestDomain(t *testing.T) {
	assert.Equal(t, Domain("abc.co"), true, "this is not domain")
	assert.Equal(t, Domain("adl-xyz.abc.co"), true, "this is not domain")
	assert.Equal(t, Domain("jd.com"), true, "this is not domain")
}

func BenchmarkIsMail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mail("abc@xyz.com")
	}
}
