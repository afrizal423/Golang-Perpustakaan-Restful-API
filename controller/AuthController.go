package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/auth"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses/formaterror"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	type token_refresh struct {
		Refresh_token string `json:"refresh_token"`
	}
	var token_old token_refresh
	var new_token map[string]string
	err := json.NewDecoder(r.Body).Decode(&token_old)
	if err != nil {
		panic(err.Error())
	}
	if auth.TokenCek(token_old.Refresh_token) != nil {
		responses.JSON(w, http.StatusInternalServerError, "Expired")
		return
	} else {
		id, err := auth.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		new_token, err = auth.BuatToken(id)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
			return
		}
		responses.JSON(w, http.StatusOK, new_token)
	}

}
