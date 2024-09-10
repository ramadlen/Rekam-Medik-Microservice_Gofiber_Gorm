package response

type UserResponse struct {
	ID                uint64 `gorm:"primaryKey" json:"id"`
	GelarDepan        string `json:"gelar_depan" gorm:"column:gelar_depan"`
	NamaLengkap       string `json:"nama_lengkap" gorm:"column:nama_lengkap" `
	GelarBelakang     string `gorm:"column:email" json:"email"`
	NamaPanggilan     string `json:"nama_panggilan" gorm:"column:nama_panggilan"`
	TempatLahir       string `json:"tempat_lahir" gorm:"column:tempat_lahir"`
	TanggalLahir      string `json:"tanggal_lahir" gorm:"column:tanggal_lahir" `
	Umur              string `json:"umur" gorm:"column:umur" `
	JenisKelamin      string `json:"jenis_kelamin" gorm:"column:jenis_kelamin"`
	Agama             string `json:"agama" gorm:"column:agama"`
	Status_Perkawinan string `json:"status_perkawinan" gorm:"column:status_perkawinan"`
	Pendidikan        string `json:"pendidikan" gorm:"column:pendidikan" `
	Pekerjaan         string `json:"pekerjaan" gorm:"column:pekerjaan" `
	Kebangsaan        string `json:"kebangsaan" gorm:"column:kebangsaan"`
}
type AlamatSekarangResponse struct {
	ID                  uint64 `json:"id" gorm:"primaryKey;column:id"`
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
type KontakResponse struct {
	ID                  uint64 `json:"id" gorm:"primaryKey;column:id"`
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