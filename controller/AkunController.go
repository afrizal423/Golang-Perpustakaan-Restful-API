package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses/formaterror"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.Anggota
	// decode data json request ke struct user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user.Persiapan("tambah")
	err = user.ValidasiLogin("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	regisAnggota, err := user.Register(config.Db)
	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	m := make(map[string]interface{})
	m["status"] = "sukses"
	m["data"] = regisAnggota
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, regisAnggota.IDAnggota))
	responses.JSON(w, http.StatusCreated, m)
	// log.Print(user)
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var user models.Pegawai
	// decode data json request ke struct user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user.Persiapan("")
	err = user.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, _ := user.ProsesLogin()
	// log.Print(token)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	if token == nil {
		m := make(map[string]string)
		m["status"] = "failed"
		m["pesan"] = "username/password salah"
		responses.JSON(w, http.StatusOK, m)
	} else {
		responses.JSON(w, http.StatusOK, token)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["status"] = "sukses"
	m["pesan"] = "Ini adalah halaman sementara dashboard, untuk update selanjutnya bakal diisi seperti atribut fitur ini"
	m["fitur"] = "hitung jumlah peminjaman dalam tahun ini, total denda, jumlah data anggota"
	responses.JSON(w, http.StatusOK, m)
}
