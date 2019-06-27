// Файл для написания вспомогательных функций
package metasdk

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// функция декодирования токена
func DecodeJwt(input, key string) string {
	claims := jwt.MapClaims{}
	input = strings.Replace(input, "v2:", "", 1)
	token, _ := jwt.ParseWithClaims(input, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if token.Valid {
		return claims["sub"].(string)
	}
	return ""
}

func DecodeJwtJSON(input, key string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	input = strings.Replace(input, "v2:", "", 1)
	token, _ := jwt.ParseWithClaims(input, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if token.Valid {
		return claims
	}
	return nil
}
