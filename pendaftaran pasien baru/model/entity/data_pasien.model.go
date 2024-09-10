package model

import (
	"time"

	"gorm.io/gorm"
)

type DataIdentitas struct {
	ID                 uint64 `json:"id" gorm:"primaryKey;column:id"`
	GelarDepan      	string `json:"gelar_depan" gorm:"column:gelar_depan"`
	NamaLengkap      	string `json:"nama_lengkap" gorm:"column:nama_lengkap" validate:"required" `
	GelarBelakang      	string `gorm:"column:email" json:"email"`
	NamaPanggilan      	string `json:"nama_panggilan" gorm:"column:nama_panggilan" validate:"required"`
	TempatLahir      	string `json:"tempat_lahir" gorm:"column:tempat_lahir" validate:"required"`
	TanggalLahir      	string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Umur      	   		string `json:"umur" gorm:"column:umur" `
	JenisKelamin      	string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" validate:"required"`
	Agama      	   		string `json:"agama" gorm:"column:agama" validate:"required"`
	Status_Perkawinan   string `json:"status_perkawinan" gorm:"column:status_perkawinan" validate:"required"`
	Pendidikan      	string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan      	   	string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Kebangsaan      	string `json:"kebangsaan" gorm:"column:kebangsaan"`
	CreateAt           time.Time `json:"create_at" gorm:"column:create_at;autoCreateTime:true;"`
	UpdateAt           time.Time `json:"update_at" gorm:"column:update_at;autoCreateTime;autoUpdateTime:true;" `
	DeleteAt      	   gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at"`
	AlamatSekarang     []AlamatSekarang `gorm:"foreignKey:user_id;references:id"`
	Kontak			   []Kontak 		`gorm:"foreignKey:user_id;references:id"`
}

func (d *DataIdentitas) DataIdentitas() string {
	return "data_identitas_pasien_baru"
}
type AlamatSekarang struct {
	ID                 			uint64 `json:"id" gorm:"primaryKey;column:id"`
	UserId						string `column:"user_id"`
	Alamat      				string `json:"alamat" gorm:"column:alamat" validate:"required"`
	RT      					string `json:"rt" gorm:"column:rt"`
	RW      					string `gorm:"column:rw" json:"rw"`
	KodePos      				string `json:"kodepos" gorm:"column:kodepos" validate:"required"`
	Provinsi      				string `json:"provinsi" gorm:"column:provinsi" validate:"required"`
	Kab_atau_Kota      			string `json:"kab_atau_kota" gorm:"column:kab_atau_kota" validate:"required"`
	Kecamatan      	   			string `json:"kecamatan" gorm:"column:kecamatan" validate:"required"`
	Kelurahan_atau_desa      	string `json:"kelurahan_atau_desa" gorm:"column:kelurahan_atau_desa" validate:"required"`
	CreateAt          			time.Time `json:"create_at" gorm:"column:create_at;autoCreateTime:true;"`
	UpdateAt           			time.Time `json:"update_at" gorm:"column:update_at;autoCreateTime;autoUpdateTime:true;" `
	DeleteAt      				gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at" `
}

func (d *AlamatSekarang) AlamatSekarang() string {
	return "alamat_pasien_sekarang"
}
type Kontak struct {
	ID                 		uint64 `json:"id" gorm:"primaryKey;column:id"`
	UserId					string `column:"user_id"`
	NamaKeluarga      		string `json:"nama_keluarga" gorm:"column:nama_keluarga" `
	JenisKelamin      		string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" `
	TanggalLahir      		string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Pendidikan      		string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan      	   		string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Alamat      	   		string `json:"alamat" gorm:"column:alamat" `
	JenisKartuIdentitas     string `json:"jenis_kartu_identitas" gorm:"column:jenis_kartu_identitas" `
	NoKartu      	   		string `json:"no_kartu" gorm:"column:no_kartu" `
	CreateAt         		time.Time `json:"create_at" gorm:"column:create_at;autoCreateTime:true;"`
	UpdateAt           		time.Time `json:"update_at" gorm:"column:update_at;autoCreateTime;autoUpdateTime:true;" `
	DeleteAt      			gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at" `
}

func (d *Kontak) Kontak() string {
	return "kontak_pasien"
}