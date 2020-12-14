package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CheckPasswordHash(t *testing.T) {
	originPassword := "0123456789"
	hashPassword, err := HashPassword(originPassword)
	require.NoError(t, err)

	c := CheckPasswordHash(originPassword, hashPassword)

	assert.True(t, c)
}
