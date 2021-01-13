package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

// Principal represents a logged in user in the system.
type Principal struct {
	FirstName string
	LastName  string
	Email     string
	Picture   string
	Roles     string
	UserID    string
	jwt.StandardClaims
}

// TokenConfig interface represents config required for generating access token and refresh token.
type TokenConfig interface {
	GetAccessTokenExpInMinute() int
	GetTokenIssuer() string
	RefreshTokenExpInMinute() int
	AccessTokenSigningKey() []byte
	RefreshTokenSigningKey() []byte
}

// CreateToken function creates access token refresh token.
func CreateToken(principal *Principal, cfg TokenConfig, aud string) (accessToken, refreshToken string, err error) {
	atExpires := time.Now().Add(time.Minute * time.Duration(cfg.GetAccessTokenExpInMinute())).Unix()
	accessUUID := uuid.New().String()
	rtExpires := time.Now().Add(time.Hour * time.Duration(cfg.RefreshTokenExpInMinute())).Unix()
	refreshUUID := uuid.New().String()
	// Creating Access Token
	atClaims := mapClaims(cfg.GetTokenIssuer(), principal.UserID, aud,
		atExpires, time.Now().Unix(), time.Now().Unix(), accessUUID, "jwt")
	addProfileClaims(&atClaims, principal)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, accessTokenErr := at.SignedString(cfg.AccessTokenSigningKey())
	if accessTokenErr != nil {
		return "", "", accessTokenErr
	}
	// Creating Refresh Token
	rtClaims := mapClaims(cfg.GetTokenIssuer(), principal.UserID, aud,
		rtExpires, time.Now().Unix(), time.Now().Unix(), refreshUUID, "jwt")
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, refreshTokenErr := rt.SignedString(cfg.RefreshTokenSigningKey())
	if refreshTokenErr != nil {
		return "", "", refreshTokenErr
	}
	return accessToken, refreshToken, nil
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
func ExtractPrincipalFromToken(encodedToken, secret string) (*Principal, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &Principal{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
