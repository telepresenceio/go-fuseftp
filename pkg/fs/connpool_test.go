package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnPoolCredentialsDefaultToAnonymous(t *testing.T) {
	p := &connPool{}
	user, password := p.credentials()
	assert.Equal(t, defaultUser, user)
	assert.Equal(t, defaultPassword, password)
}

func TestConnPoolCredentialsUseConfiguredValues(t *testing.T) {
	p := &connPool{user: "bob", password: "s3cr3t"}
	user, password := p.credentials()
	assert.Equal(t, "bob", user)
	assert.Equal(t, "s3cr3t", password)
}

func TestConnPoolCredentialsPartialConfiguration(t *testing.T) {
	p := &connPool{user: "bob"}
	user, password := p.credentials()
	assert.Equal(t, "bob", user)
	assert.Equal(t, defaultPassword, password)

	p = &connPool{password: "s3cr3t"}
	user, password = p.credentials()
	assert.Equal(t, defaultUser, user)
	assert.Equal(t, "s3cr3t", password)
}
