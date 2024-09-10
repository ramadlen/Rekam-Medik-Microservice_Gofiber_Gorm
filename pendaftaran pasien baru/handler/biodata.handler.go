package handler

import (
	"pendaftaran_pasien_baru/db"
	model "pendaftaran_pasien_baru/model/entity"
	"pendaftaran_pasien_baru/model/request"

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
		GelarDepan: user.GelarDepan,
		NamaLengkap: user.NamaLengkap,
		GelarBelakang: user.GelarBelakang,
		NamaPanggilan: user.NamaPanggilan,
		TempatLahir: user.TempatLahir,
		Umur: user.Umur,
		JenisKelamin: user.JenisKelamin,
		Agama: user.Agama,
		Status_Perkawinan: user.Status_Perkawinan,
		Pendidikan: user.Pendidikan,
		Pekerjaan: user.Pekerjaan,
		Kebangsaan: user.Kebangsaan,
	}
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
	//URL menjadi acak
	userId := ctx.Params("id")
	// hashedParam, errHashedParam  := bcrypt.GenerateFromPassword([]byte("id"), bcrypt.DefaultCost)
	// if errHashedParam != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).SendString("Error encrypting parameter")
	// }

	// return ctx.SendString(fmt.Sprintf("Encrypted parameter: %s", string(hashedParam)))
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
	user.GelarDepan = userRequest.GelarDepan
	user.NamaLengkap = userRequest.NamaLengkap
	user.GelarBelakang = userRequest.GelarBelakang
	user.NamaPanggilan = userRequest.NamaPanggilan
	user.TempatLahir = userRequest.TempatLahir
	user.Umur = userRequest.Umur
	user.JenisKelamin = userRequest.JenisKelamin
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




func AlamatSekarangHandler(bh *fiber.Ctx) error{
	var alamats []model.AlamatSekarang
	result := db.DbInit().Find(&alamats)
	if result.Error != nil {
log.Println(result.Error)
	}
// 	err := db.DbInit().Find(&users).Error
// 	if err != nil {
// log.Println(err)
// 	}
return  bh.JSON(alamats)
}

func AlamatSekarangCreate(ctx *fiber.Ctx) error {
	alamatRequest := new(request.AlamatSekarangCreateRequest)
	if err:= ctx.BodyParser(alamatRequest); err!= nil {
		return err
	}
	//Validasi Request
	validate := validator.New()
	errValidate := validate.Struct(alamatRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}
	newAlamat := model.AlamatSekarang{
	Alamat: alamatRequest.Alamat,
	RT: alamatRequest.RT,
	RW: alamatRequest.RW,
	KodePos: alamatRequest.KodePos,
	Provinsi: alamatRequest.Provinsi,
	Kab_atau_Kota: alamatRequest.Kab_atau_Kota,
	Kecamatan: alamatRequest.Kecamatan,
	Kelurahan_atau_desa: alamatRequest.Kelurahan_atau_desa,
	}
	errCreateUser := db.DbInit().Create(&newAlamat).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to save your data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newAlamat,
	})
}


func AlamatSekarangGetById(ctx *fiber.Ctx) error {
	alamatId := ctx.Params("id")
	var alamat model.AlamatSekarang
	err := db.DbInit().First(&alamat, "id = ?", alamatId).Error
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
		"data": alamat,
	})
}

func AlamatSekarangUpdate(ctx *fiber.Ctx) error {
	alamatRequest := new(request.AlamatSekarangUpdateRequest)
	if err := ctx.BodyParser(alamatRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
		"message": "bad request",
	})
	}
	var alamat model.AlamatSekarang
	//URL menjadi acak
	alamatId := ctx.Params("id")
	// hashedParam, errHashedParam  := bcrypt.GenerateFromPassword([]byte("id"), bcrypt.DefaultCost)
	// if errHashedParam != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).SendString("Error encrypting parameter")
	// }

	// return ctx.SendString(fmt.Sprintf("Encrypted parameter: %s", string(hashedParam)))
	// CHECK KETERSEDIAAN USER
	err := db.DbInit().First(&alamat, "id = ?", alamatId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
		"message": "alamat tidak ditemukan",
	})
	}
	//UPDATE USER DATA
	if alamatRequest.Alamat != "" {
		alamat.Alamat = alamatRequest.Alamat
	}
	alamat.Alamat = alamatRequest.Alamat
	alamat.RT = alamatRequest.RT
	alamat.RW = alamatRequest.RW
	alamat.KodePos = alamatRequest.KodePos
	alamat.Provinsi = alamatRequest.Provinsi
	alamat.Kab_atau_Kota = alamatRequest.Kab_atau_Kota
	alamat.Kecamatan = alamatRequest.Kecamatan
	alamat.Kelurahan_atau_desa = alamatRequest.Kelurahan_atau_desa
	errUpdate := db.DbInit().Save(&alamat).Error
	if errUpdate != nil {
			return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
return ctx.JSON(fiber.Map{
	"message": "success",
	"data": alamat,
})	
}


func AlamatSekarangDelete(ctx *fiber.Ctx) error {
	alamatId:= ctx.Params("id")
	var alamat model.AlamatSekarang
	//CHECK AVAILABLE USER
	err := db.DbInit().Debug().First(&alamat, "id=?", alamatId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "alamat tidak ditemukan",
		})
	}
	errDelete := db.DbInit().Debug().Delete(&alamat).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "alamat was deleted",
	})
}

func KontakHandler(bh *fiber.Ctx) error{
	var kontaks []model.Kontak
	result := db.DbInit().Find(&kontaks)
	if result.Error != nil {
log.Println(result.Error)
	}
// 	err := db.DbInit().Find(&kontaks).Error
// 	if err != nil {
// log.Println(err)
// 	}
return  bh.JSON(kontaks)
}

func KontakHandlerCreate(ctx *fiber.Ctx) error {
	kontak := new(request.KontakCreateRequest)
	if err:= ctx.BodyParser(kontak); err!= nil {
		return err
	}
	//Validasi Request
	validate := validator.New()
	errValidate := validate.Struct(kontak)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}
	newKontak := model.Kontak{
		NamaKeluarga: kontak.NamaKeluarga,
		JenisKelamin: kontak.JenisKelamin,
		TanggalLahir: kontak.TanggalLahir,
		Pendidikan: kontak.Pendidikan,
		Pekerjaan: kontak.Pekerjaan,
		Alamat: kontak.Alamat,
		JenisKartuIdentitas: kontak.JenisKartuIdentitas,
		NoKartu: kontak.NoKartu,
	}
	errCreateKontak := db.DbInit().Create(&newKontak).Error
	if errCreateKontak != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to save your data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newKontak,
	})
}


func KontakHandlerGetById(ctx *fiber.Ctx) error {
	kontakId := ctx.Params("id")
	var kontak model.Kontak
	err := db.DbInit().First(&kontak, "id = ?", kontakId).Error
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
		"data": kontak,
	})
}

func KontakHandlerUpdate(ctx *fiber.Ctx) error {
	kontakRequest := new(request.KontakUpdateRequest)
	if err := ctx.BodyParser(kontakRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
		"message": "bad request",
	})
	}
	var kontak model.Kontak
	//URL menjadi acak
	kontakId := ctx.Params("id")
	// hashedParam, errHashedParam  := bcrypt.GenerateFromPassword([]byte("id"), bcrypt.DefaultCost)
	// if errHashedParam != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).SendString("Error encrypting parameter")
	// }

	// return ctx.SendString(fmt.Sprintf("Encrypted parameter: %s", string(hashedParam)))
	// CHECK KETERSEDIAAN USER
	err := db.DbInit().First(&kontak, "id = ?", kontakId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
		"message": "user not found",
	})
	}
	//UPDATE USER DATA
	if kontakRequest.NamaKeluarga != "" {
		kontak.NamaKeluarga = kontakRequest.NamaKeluarga
	}
	kontak.NamaKeluarga = kontakRequest.NamaKeluarga
	kontak.JenisKelamin = kontakRequest.JenisKelamin
	kontak.TanggalLahir = kontakRequest.TanggalLahir
	kontak.Pendidikan = kontakRequest.Pendidikan
	kontak.Pekerjaan = kontakRequest.Pekerjaan
	kontak.Alamat = kontakRequest.Alamat
	kontak.JenisKartuIdentitas = kontakRequest.JenisKartuIdentitas
	kontak.NoKartu = kontakRequest.NoKartu
	errUpdate := db.DbInit().Save(&kontak).Error
	if errUpdate != nil {
			return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
return ctx.JSON(fiber.Map{
	"message": "success",
	"data": kontak,
})	
}


func KontakHandlerDelete(ctx *fiber.Ctx) error {
	kontakId := ctx.Params("id")
	var kontak model.Kontak
	//CHECK AVAILABLE USER
	err := db.DbInit().Debug().First(&kontak, "id=?", kontakId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	errDelete := db.DbInit().Debug().Delete(&kontak).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "kontak was deleted",
	})
}