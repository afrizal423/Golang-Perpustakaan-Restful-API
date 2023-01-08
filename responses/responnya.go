package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Status string `json:"status"`
			Pesan  string `json:"pesan"`
		}{
			Status: "error",
			Pesan:  err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

func Sukses(data interface{}) interface{} {
	m := make(map[string]interface{})
	m["status"] = "sukses"
	m["data"] = data
	return m
}

func Gagal(data interface{}, pesan string) interface{} {
	m := make(map[string]interface{})
	m["status"] = "gagal"
	m["status"] = pesan
	m["data"] = data
	return m
}
