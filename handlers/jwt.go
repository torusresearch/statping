package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/torusresearch/statping/types/users"
	"net/http"
	"time"
)

type JwtClaim struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

func removeJwtToken(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &c)
}

func setJwtToken(user *users.User, w http.ResponseWriter) (JwtClaim, string) {
	expirationTime := time.Now().Add(72 * time.Hour)
	jwtClaim := JwtClaim{
		Username: user.Username,
		Admin:    user.Admin.Bool,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Errorln("error setting token: ", err)
	}
	user.Token = tokenString
	// set cookies
	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		Value:   tokenString,
		Expires: expirationTime,
		MaxAge:  int(time.Duration(72 * time.Hour).Seconds()),
		Path:    "/",
	})
	return jwtClaim, tokenString
}

func getJwtToken(r *http.Request) (JwtClaim, error) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return JwtClaim{}, err
		}
		return JwtClaim{}, err
	}

	var claims JwtClaim
	tkn, err := jwt.ParseWithClaims(c.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return JwtClaim{}, err
		}
		return JwtClaim{}, err
	}
	if !tkn.Valid {
		return claims, errors.New("token is not valid")
	}
	return claims, err
}
