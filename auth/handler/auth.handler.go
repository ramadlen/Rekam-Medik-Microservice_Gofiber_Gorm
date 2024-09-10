package handler

import (
	"biodata/db"
	model "biodata/model/entity"
	"biodata/model/request"
	"biodata/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
loginRequest := new(request.LoginRequest)
if err:= ctx.BodyParser(loginRequest);err!= nil {
	return err
}
log.Println(loginRequest)

//Validasi Request
validate:= validator.New()
errValidate:=validate.Struct(loginRequest) 
if errValidate != nil{
	return ctx.Status(400).JSON(fiber.Map{
		"message": "failed",
		"error": errValidate.Error(),
	})
}

var user model.DataIdentitas
err:= db.DbInit().First(&user, "email=?", loginRequest.Email).Error
if err != nil {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "user not found",
	})
}
isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
if !isValid {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "wrong credential",
	})
}

//GENERATE JWT
claims := jwt.MapClaims{}
claims["name_panggilan"] = user.NamaPanggilan
claims["email"] = user.Email
claims["golongan_darah"] = user.GolonganDarah
claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

token,errGenerateToken  := utils.GenerateToken(&claims)
if errGenerateToken != nil {
	log.Println(errGenerateToken)
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "wrong credential",
	})
}
return ctx.JSON(fiber.Map{
	"token": token,
})
}