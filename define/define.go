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
	EmqxAddr   = "http://192.168.1.8:18083/api/v5"
	EmqxKey    = "1f9c5b734fe27865"
	EmqxSecret = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
)
