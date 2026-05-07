package constant

import "time"

const (
	RequestIDKey    = "request_id"
	RequestIDHeader = "X-Request-ID"
)

type JwtType time.Time

var (
	JwtAccessTokenExpires  = JwtType(time.Now().Add(24 * time.Hour))      // 24 hours
	JwtRefreshTokenExpires = JwtType(time.Now().Add(30 * 24 * time.Hour)) // 1 month
)
