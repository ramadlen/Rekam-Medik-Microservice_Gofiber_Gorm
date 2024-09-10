package handler

import (
	"biodata/db"
	model "biodata/model/entity"
	"biodata/model/request"
	"biodata/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BiodataHandler(bh *fiber.Ctx) error{
	var users []model.DataIdentitas
	result := db.DbInit().Find(&users)
	if result.Error != nil {
log.Println(result.Error)
	}
// 	err := db.DbInit().Find(&users).Error
// 	if err != nil {
// log.Println(err)
// 	}
return  bh.JSON(users)
}

func BiodataHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserRequestCreateRequest)
	if err:= ctx.BodyParser(user); err!= nil {
		return err
	}
	//Validasi Request
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}
	newUser := model.DataIdentitas{
		NamaPanggilan: user.NamaPanggilan,
		GolonganDarah: user.GolonganDarah,
		Email: user.Email,
	}
	hashedPassword, err  := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	newUser.Password = hashedPassword

	errCreateUser := db.DbInit().Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to save your data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newUser,
	})
}


func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user model.DataIdentitas
	err := db.DbInit().First(&user, "id = ?", userId).Error
	if err != nil {
	return ctx.Status(404).JSON(fiber.Map{
		"message": "user not found",
	})		
	}
	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	NamaPanggilan: user.NamaPanggilan,
	// 	GolonganDarah: user.GolonganDarah,
	// 	CreateAt: user.CreateAt,
	// 	UpdateAt: user.UpdateAt,
	// }
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
		"message": "bad request",
	})
	}
	var user model.DataIdentitas

	userId := ctx.Params("id")
	// CHECK KETERSEDIAAN USER
	err := db.DbInit().First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
		"message": "user not found",
	})
	}
	//UPDATE USER DATA
	if userRequest.NamaPanggilan != "" {
		user.NamaPanggilan = userRequest.NamaPanggilan
	}
	user.GolonganDarah = userRequest.GolonganDarah
	errUpdate := db.DbInit().Save(&user).Error
	if errUpdate != nil {
			return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
return ctx.JSON(fiber.Map{
	"message": "success",
	"data": user,
})	
}


func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user model.DataIdentitas
	//CHECK AVAILABLE USER
	err := db.DbInit().Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	errDelete := db.DbInit().Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}


func UserHandlerUpdateEmail(ctx *fiber.Ctx) error  {
	userRequest:= new(request.UserEmailRequest)
	if err:= ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": " bad request",
		})
	}
	var user model.DataIdentitas
	var isEmailUserExist model.DataIdentitas
	userId:= ctx.Params("id")

	//Check Ketersediaan User
	err:= db.DbInit().First(&user, "id = ?", userId).Error
	if err!= nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//Check Available Email
	errCheckEmail := db.DbInit().First(&isEmailUserExist, "email = ?", &userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "email alreay used.",
		})
	}

	//Update data User

	user.Email= userRequest.Email

	errUpdate := db.DbInit().Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}