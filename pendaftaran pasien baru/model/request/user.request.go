package request

type UserRequestCreateRequest struct {
	GelarDepan        string `json:"gelar_depan" gorm:"column:gelar_depan"`
	NamaLengkap       string `json:"nama_lengkap" gorm:"column:nama_lengkap" validate:"required" `
	GelarBelakang     string `gorm:"column:email" json:"email"`
	NamaPanggilan     string `json:"nama_panggilan" gorm:"column:nama_panggilan" validate:"required"`
	TempatLahir       string `json:"tempat_lahir" gorm:"column:tempat_lahir" validate:"required"`
	TanggalLahir      string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Umur              string `json:"umur" gorm:"column:umur" `
	JenisKelamin      string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" validate:"required"`
	Agama             string `json:"agama" gorm:"column:agama" validate:"required"`
	Status_Perkawinan string `json:"status_perkawinan" gorm:"column:status_perkawinan" validate:"required"`
	Pendidikan        string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan         string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Kebangsaan        string `json:"kebangsaan" gorm:"column:kebangsaan"`
}
type UserUpdateRequest struct {
	GelarDepan        string `json:"gelar_depan" gorm:"column:gelar_depan"`
	NamaLengkap       string `json:"nama_lengkap" gorm:"column:nama_lengkap" validate:"required" `
	GelarBelakang     string `gorm:"column:email" json:"email"`
	NamaPanggilan     string `json:"nama_panggilan" gorm:"column:nama_panggilan" validate:"required"`
	TempatLahir       string `json:"tempat_lahir" gorm:"column:tempat_lahir" validate:"required"`
	TanggalLahir      string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Umur              string `json:"umur" gorm:"column:umur" `
	JenisKelamin      string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" validate:"required"`
	Agama             string `json:"agama" gorm:"column:agama" validate:"required"`
	Status_Perkawinan string `json:"status_perkawinan" gorm:"column:status_perkawinan" validate:"required"`
	Pendidikan        string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan         string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Kebangsaan        string `json:"kebangsaan" gorm:"column:kebangsaan"`
}

//	type UserEmailRequest struct {
//		Email string `gorm:"column:email" json:"email" validate:"required"`
//	}
type AlamatSekarangCreateRequest struct {
	UserId              string `column:"user_id"`
	Alamat              string `json:"alamat" gorm:"column:alamat" validate:"required"`
	RT                  string `json:"rt" gorm:"column:rt"`
	RW                  string `gorm:"column:rw" json:"rw"`
	KodePos             string `json:"kodepos" gorm:"column:kodepos" validate:"required"`
	Provinsi            string `json:"provinsi" gorm:"column:provinsi" validate:"required"`
	Kab_atau_Kota       string `json:"kab_atau_kota" gorm:"column:kab_atau_kota" validate:"required"`
	Kecamatan           string `json:"kecamatan" gorm:"column:kecamatan" validate:"required"`
	Kelurahan_atau_desa string `json:"kelurahan_atau_desa" gorm:"column:kelurahan_atau_desa" validate:"required"`
}
type KontakCreateRequest struct {
	UserId              string `column:"user_id"`
	NamaKeluarga        string `json:"nama_keluarga" gorm:"column:nama_keluarga" `
	JenisKelamin        string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" `
	TanggalLahir        string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Pendidikan          string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan           string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Alamat              string `json:"alamat" gorm:"column:alamat" `
	JenisKartuIdentitas string `json:"jenis_kartu_identitas" gorm:"column:jenis_kartu_identitas" `
	NoKartu             string `json:"no_kartu" gorm:"column:no_kartu" `
}
type AlamatSekarangUpdateRequest struct {
	UserId              string `column:"user_id"`
	Alamat              string `json:"alamat" gorm:"column:alamat" validate:"required"`
	RT                  string `json:"rt" gorm:"column:rt"`
	RW                  string `gorm:"column:rw" json:"rw"`
	KodePos             string `json:"kodepos" gorm:"column:kodepos" validate:"required"`
	Provinsi            string `json:"provinsi" gorm:"column:provinsi" validate:"required"`
	Kab_atau_Kota       string `json:"kab_atau_kota" gorm:"column:kab_atau_kota" validate:"required"`
	Kecamatan           string `json:"kecamatan" gorm:"column:kecamatan" validate:"required"`
	Kelurahan_atau_desa string `json:"kelurahan_atau_desa" gorm:"column:kelurahan_atau_desa" validate:"required"`
}
type KontakUpdateRequest struct {
	UserId              string `column:"user_id"`
	NamaKeluarga        string `json:"nama_keluarga" gorm:"column:nama_keluarga" `
	JenisKelamin        string `json:"jenis_kelamin" gorm:"column:jenis_kelamin" `
	TanggalLahir        string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Pendidikan          string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan           string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Alamat              string `json:"alamat" gorm:"column:alamat" `
	JenisKartuIdentitas string `json:"jenis_kartu_identitas" gorm:"column:jenis_kartu_identitas" `
	NoKartu             string `json:"no_kartu" gorm:"column:no_kartu" `
}