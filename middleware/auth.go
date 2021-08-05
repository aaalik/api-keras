package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aaalik/api-keras/helper"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	jwtf3t "github.com/form3tech-oss/jwt-go"
)

type TokenDetails struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	RtExpires    int64     `json:"rt_expires"`
	Scope        string    `json:"scope"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateToken(userid int64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.CreatedAt = time.Now()

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["id_user"] = userid
	atClaims["exp"] = td.ExpiresIn
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		helper.Log.Error(err)
		return nil, err
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))
	if err != nil {
		helper.Log.Error(err)
		return nil, err
	}
	return td, nil
}

func VerifyToken(bearToken string) error {
	tokenString := ""

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		helper.Log.Error(err)
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		helper.Log.Error(err)
		return err
	}

	return nil
}

func InitMiddleware() *jwtmiddleware.JWTMiddleware {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwtf3t.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware
}
