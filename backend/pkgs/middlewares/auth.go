package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func APIKeyAuthMiddleware() func(*fiber.Ctx) error {
	return keyauth.New(keyauth.Config{
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.SendStatus(401)
		},

		Validator: func(c *fiber.Ctx, key string) (bool, error) {

			hashedAPIKey := sha256.Sum256([]byte(viper.GetString("secret.token")))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, nil
		},
	})
}

func APICookieAuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("_token")

	secretKey := viper.GetString("secret.token")

	_, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.SendStatus(401)
	}

	return c.Next()
}
