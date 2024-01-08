package utils

import (
	"amikom-pedia-api/utils"
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	var randomString []string
	for i := 1; i <= 10; i++ {
		randomString = append(randomString, utils.RandomString(16))
		fmt.Println(randomString)
	}

	require.Len(t, randomString, 10)
	require.Len(t, randomString[0], 16)

}

func TestGenerateRandomInt(t *testing.T) {
	var randomInt []int64
	for i := 1; i <= 10; i++ {
		randomInt = append(randomInt, utils.RandomInt(1, 10))
		fmt.Println(randomInt)
	}

	require.Len(t, randomInt, 10)
	require.GreaterOrEqual(t, randomInt[0], int64(1))
	require.LessOrEqual(t, randomInt[0], int64(10))
}

func TestGenerateRandomCombine(t *testing.T) {
	var randomCombine []string
	for i := 1; i <= 10; i++ {
		randomNumber := utils.RandomInt(1, 100000)
		stringNumber := strconv.FormatInt(randomNumber, 10)
		randomCombine = append(randomCombine, utils.RandomString(10)+stringNumber)
		fmt.Println(randomCombine)
	}

	require.Len(t, randomCombine, 10)

	fmt.Println("===========================================")

	for i := 1; i <= 10; i++ {
		randomCombine := utils.RandomCombineIntAndString()
		fmt.Println(randomCombine)
	}
}
