package test

import (
	"amikom-pedia-api/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword"
	wrongPassword := "wrongPassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	t.Run("correct password", func(t *testing.T) {
		isCorrect := utils.CheckPasswordHash(password, hash)
		assert.True(t, isCorrect)
	})

	t.Run("incorrect password", func(t *testing.T) {
		isCorrect := utils.CheckPasswordHash(wrongPassword, hash)
		assert.False(t, isCorrect)
	})
}
