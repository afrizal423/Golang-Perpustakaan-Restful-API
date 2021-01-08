package controller

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Print(user)
}
