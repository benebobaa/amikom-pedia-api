package utils

import (
	"amikom-pedia-api/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := utils.RandomString(6)

	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = utils.CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := utils.RandomString(6)
	err = utils.CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
