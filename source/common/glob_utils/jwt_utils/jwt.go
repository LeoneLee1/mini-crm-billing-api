package jwtutils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetSecretKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "default-secret-key"
	}
	return []byte(key)
}

type JWTClaim struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Nik     string `json:"nik"`
	Jabatan string `json:"jabatan"`
	Divisi  string `json:"divisi"`
	Role    string `json:"role"`
	Type    string `json:"type"`
	jwt.RegisteredClaims
}

func CreateAccessToken(ID string, Name string, Nik string, Jabatan string, Divisi string, Role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaim{
		ID:      ID,
		Name:    Name,
		Nik:     Nik,
		Jabatan: Jabatan,
		Divisi:  Divisi,
		Role:    Role,
		Type:    "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetSecretKey())
}

func CreateRefreshToken(ID string) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := &JWTClaim{
		ID:   ID,
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetSecretKey())
}

func VerifyToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return GetSecretKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			if token != nil {
				claims, ok := token.Claims.(*JWTClaim)
				if ok {
					return claims, errors.New("token expired")
				}
			}
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}

func GetCurrentUser(c *gin.Context) (*JWTClaim, bool) {
	userValue, exists := c.Get("jwt_claim")
	if !exists {
		return nil, false
	}

	jwtClaims, ok := userValue.(*JWTClaim)
	return jwtClaims, ok
}

func GetCurrentUserID(c *gin.Context) (string, bool) {
	userIDValue, exists := c.Get("id")
	if !exists {
		return "", false
	}

	userID, ok := userIDValue.(string)
	return userID, ok
}
