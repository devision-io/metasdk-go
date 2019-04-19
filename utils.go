// Файл для написания вспомогательных функций
package metasdk

import (
	"github.com/dgrijalva/jwt-go"
	"log"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// функция декодирования токена
func JwtDecode(input, key string) string {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(input, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if token.Valid {
		return claims["sub"].(string)
	}
	return ""
}
