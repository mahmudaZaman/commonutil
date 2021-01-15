package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

// Principal represents a logged in user in the system.
type Principal struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	Roles     string `json:"roles"`
	UserID    string `json:"sub"`
	jwt.StandardClaims
}

// CreateJwtToken Creates jwt token.
func CreateJwtToken(
	issuer, aud string, signingKey []byte, atExpiresUnix int64, principal *Principal, includeProfile bool) (jwtToken string, err error) {
	jti := uuid.New().String()
	atClaims := mapClaims(issuer, principal.UserID, aud,
		atExpiresUnix, time.Now().Unix(), time.Now().Unix(), jti, "jwt")
	if includeProfile {
		addProfileClaims(&atClaims, principal)
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString(signingKey)
}

func addProfileClaims(claims *jwt.MapClaims, principal *Principal) {
	(*claims)["firstName"] = principal.FirstName
	(*claims)["lastName"] = principal.LastName
	(*claims)["email"] = principal.Email
	(*claims)["picture"] = principal.Picture
	(*claims)["roles"] = principal.Roles
}

// https://en.wikipedia.org/wiki/JSON_Web_Token
func mapClaims(iss, sub, aud string, exp, nbf, iat int64, jti, typ string) jwt.MapClaims {
	atClaims := jwt.MapClaims{}
	atClaims["iss"] = iss
	atClaims["sub"] = sub
	atClaims["aud"] = aud
	atClaims["exp"] = exp
	atClaims["nbf"] = nbf
	atClaims["iat"] = iat
	atClaims["jti"] = jti
	atClaims["typ"] = typ
	return atClaims
}

// ExtractPrincipalFromToken extracts Principal from token.
func ExtractPrincipalFromToken(encodedToken string, secret []byte) (*Principal, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &Principal{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token not valid")
	}
	principal, ok := token.Claims.(*Principal)
	if !ok {
		return nil, errors.New("claims could not cast to principal")
	}
	return principal, nil
}
