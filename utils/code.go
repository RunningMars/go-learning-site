package utils

import (
    "math/rand"
    "time"
)

func GenerateRandomCode(length int) string {
    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    code := make([]byte, length)
    for i := 0; i < length; i++ {
        code[i] = digits[rand.Intn(len(digits))]
    }
    return string(code)
}