package msgoutils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type JWTService struct {
	secret string
}

type JWTClaim struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: secret,
	}
}

func (srv *JWTService) ValidateToken(token string) (JWTClaim, error) {
	t, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(srv.secret), nil
	})

	if err != nil {
		return JWTClaim{}, err
	}

	if !t.Valid {
		return JWTClaim{}, errors.New("invalid JWT token")
	}

	claims := t.Claims.(*JWTClaim)

	return *claims, nil
}

func JWTMiddleware(secret string) fiber.Handler {
	jwtService := NewJWTService(secret)
	logger := logrus.WithField("context", "jwt_middleware")

	return func(c *fiber.Ctx) error {
		logger.Info("Validating JWT Token")

		headers := c.GetReqHeaders()["Authentication"]
		if len(headers) == 0 {
			logger.Debug("No header found")
			return c.SendStatus(401)
		}

		token := headers[0]
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			logger.WithError(err).Error("Cannot validate JWT token")
			return c.SendStatus(401)
		}

		c.Locals("jwt", claims)
		return c.Next()
	}
}
