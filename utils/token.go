package utils

import (
	"math/rand"
	"time"
)

func GenerateNewOtpValue() string {
	rand.Seed(time.Now().UnixNano())
	const digits = "0123456789"
	otpValue := make([]byte, 6)
	for i := range otpValue {
		otpValue[i] = digits[rand.Intn(len(digits))]
	}
	return string(otpValue)
}

func GenerateNewRefCode() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	refCode := make([]byte, 10)
	for i := range refCode {
		refCode[i] = letters[rand.Intn(len(letters))]
	}
	return string(refCode)
}
