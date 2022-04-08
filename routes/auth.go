package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/moaabid/golang-fiber-jwt/data"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

//Signup

func SignUp(c *fiber.Ctx) error {

	req := new(SignupRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if req.Name == " " || req.Email == " " || req.Password == " " {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid signup credentials")

	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &data.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	_, err = data.Engine.Insert(user)
	if err != nil {
		return err
	}

	token, exp, err := createJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"auth_token": token, "expiration_time": exp, "user": user})
}

//login

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if req.Email == " " || req.Password == " " {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid signup credentials")

	}

	user := new(data.User)

	has, err := data.Engine.Where("email= ?", req.Email).Desc("id").Get(user)
	if err != nil {
		return err
	}

	if !has {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid login crenditals")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid login crenditals. Check password is correct")
	}

	token, exp, err := createJWTToken(*user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"auth_token": token, "expiration_time": exp, "user": user})

}

func Private(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"success": true, "path": "private"})
}

func Public(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"success": true, "path": "public"})
}

func createJWTToken(user data.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = exp

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}
