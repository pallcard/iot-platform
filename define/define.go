package define

import "github.com/golang-jwt/jwt/v4"

type M map[string]interface{}

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"Identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var (
	JwtKey     = "iot_platform"
	EmqxKey    = "iot_platform"
	EmqxSecret = "iot_platform"
)
