package wso2jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

type keys struct {
	Keys []struct {
		E   string   `json:"e"`
		Kty string   `json:"kty"`
		Use string   `json:"use"`
		Kid string   `json:"kid"`
		N   string   `json:"n"`
		X5C []string `json:"x5c"`
	} `json:"keys"`
}

// ValidateJWT is the middleware function
func ValidateJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if len(os.Getenv("LOCAL_ENVIRONMENT")) == 0 {
				token := context.Request().Header().Get("X-jwt-assertion")
				if token == "" {
					return jsonresp.New(context, http.StatusBadRequest, "No WSO2-provided `X-jwt-assertion` header present")
				}

				err := validate(token)
				if err != nil {
					return jsonresp.New(context, http.StatusBadRequest, err.Error())
				}
			}

			return next(context)
		}
	}
}

func validate(token string) error {
	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if parsedToken.Method.Alg() != "RS256" { // Check that our keys are signed with RS256 as expected (https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/)
			return nil, fmt.Errorf("Unexpected signing method: %v", parsedToken.Header["alg"]) // This error never gets returned to the user but may be useful for debugging/logging at some point
		}

		// Look up key
		key, err := lookupSigningKey()
		if err != nil {
			return nil, err
		}

		// Unpack key from PEM encoded PKCS8
		return jwt.ParseRSAPublicKeyFromPEM(key)
	})

	log.Printf("%+v", parsedToken)

	if parsedToken.Valid {
		return nil
	} else if validationError, ok := err.(*jwt.ValidationError); ok {
		if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("Authorization token is malformed")
		} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errors.New("Authorization token is expired")
		}
	}

	return errors.New("Not authorized")
}

func lookupSigningKey() ([]byte, error) {
	response, err := http.Get("https://api.byu.edu/.well-known/byucerts")
	if err != nil {
		return nil, err
	}

	allKeys := keys{}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &allKeys)
	if err != nil {
		return nil, err
	}

	certificate := "-----BEGIN CERTIFICATE-----\n" + allKeys.Keys[0].X5C[0] + "\n-----END CERTIFICATE-----"
	log.Println(certificate)
	return []byte(certificate), nil
}
