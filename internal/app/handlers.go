package app

import (
	"check-mate/internal/model"
	"check-mate/pkg/bcrypt"
	"check-mate/pkg/config"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func (s *Server) SignupHandler(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	u, err := s.userService.GetUser(user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if u != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "username is already busy",
		})
	}

	hash, err := bcrypt.Generate(user.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "username is already busy",
		})
	}

	user.Password = hash

	id, err := s.userService.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	t, err := getToken(id, user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}

func (s *Server) SigninHandler(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	u, err := s.userService.GetUser(user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "user not found",
		})
	}

	err = bcrypt.Compare(u.Password, user.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	t, err := getToken(u.Id, u.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}

func getToken(id uint, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString(config.SECRET)
	return t, err
}
