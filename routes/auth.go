package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moaabid/golang-fiber-jwt/data"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}

func SignUp(c *fiber.Ctx) error {

	req := new(SignupRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if req.Name == " " || req.Email == " " || req.Password == " " {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid signup credentials")

	}

	_, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &data.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err = data.Engine.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

func Login(c *fiber.Ctx) error {

	return nil
}

func Private(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"success": true, "path": "private"})
}

func Public(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"success": true, "path": "public"})
}

// func createJWTToken() (string, int64, error) {
// 	exp := time.Now().Add(time.Minute * 30).Unix()
// 	token := jwt.New(jwt.SigningMethodES256)

// 	claims := token.Claims.(jwt.MapClaims)

// }
