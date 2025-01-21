package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtSecret = []byte("your-secret-key") // 在实际应用中应该从配置文件读取

func GenerateJWTToken(userID int64) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}