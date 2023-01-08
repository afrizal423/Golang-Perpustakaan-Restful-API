package BukuController

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses/formaterror"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models/BukuModels"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/biezhi/gorm-paginator/pagination"
)

func LihatSemuaBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.Buku{}
	err, querydb := buku.LihatSemuaBuku(config.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	// Lihat Detail jenis buku
	if id != 0 {
		bukus, err := buku.DetailBuku(config.Db, uint32(id))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		// fmt.Println(bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(bukus))
		return
	}
	if query.Get("search") == "true" {
		cariBuku(w, r)
		return
	}
	var tmpbuku []BukuModels.Buku
	paginator := pagination.Paging(&pagination.Param{
		DB:    querydb,
		Page:  int(helper.StringkeInt(query.Get("page"))),
		Limit: 3,
		// OrderBy: []string{"id desc"},
	}, &tmpbuku)
	responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
}

func cariBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.Buku{}
	query := r.URL.Query()
	if query.Get("q") != "" && query.Get("orderby") != "" {
		kata := helper.NonLatinChar(query.Get("q"))[0]
		orderby := helper.NonLatinChar(query.Get("orderby"))[0]
		_, data := buku.SearchBuku(config.Db, kata, orderby)

		var bukus []BukuModels.Buku
		paginator := pagination.Paging(&pagination.Param{
			DB:    data,
			Page:  int(helper.StringkeInt(query.Get("page"))),
			Limit: 3,
			// OrderBy: []string{"id desc"},
		}, &bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
		fmt.Println(kata, orderby)
		return
	}
}

func TambahBuku(w http.ResponseWriter, r *http.Request) {
	// body, _ := ioutil.ReadAll(r.Body)
	buku := BukuModels.Buku{}
	buku.ISBN = r.FormValue("isbn")
	buku.IDKategoriJenis = int32(helper.StringkeInt(r.FormValue("id_kategori_buku")))
	buku.Judul = r.FormValue("judul_buku")
	buku.IDPenulisBuku = int32(helper.StringkeInt(r.FormValue("id_penulis_buku")))
	buku.IDPenerbitBuku = int32(helper.StringkeInt(r.FormValue("id_penerbit_buku")))
	buku.ThnTerbit = r.FormValue("tahun_terbit")
	buku.StokBuku = int32(helper.StringkeInt(r.FormValue("stok_buku")))
	buku.RakBuku = r.FormValue("rak_buku")
	buku.DeskripsiBuku = r.FormValue("deskripsi_buku")
	//gambar disini
	buku.Gambarbuku = "assets/buku/" + uploadGambar(w, r)
	//end gambar
	buku.Kondisibuku = r.FormValue("kondisi_buku")

	// lihat semua data
	r.ParseForm()
	for key, values := range r.Form {
		fmt.Println(key, values)
	}
	fmt.Println(buku)

	tmbhData, err := buku.TambahBuku(config.Db)
	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, tmbhData.IDBuku))
	responses.JSON(w, http.StatusCreated, responses.Sukses(tmbhData))
}

func UpdateBuku(w http.ResponseWriter, r *http.Request) {
	var buku BukuModels.Buku
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	if id == 0 {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("URL Invalid"))
		return
	}

	// hapus gambar terlebih dahulu
	hapusGambar(uint32(id), w)

	buku.ISBN = r.FormValue("isbn")
	buku.IDKategoriJenis = int32(helper.StringkeInt(r.FormValue("id_kategori_buku")))
	buku.Judul = r.FormValue("judul_buku")
	buku.IDPenulisBuku = int32(helper.StringkeInt(r.FormValue("id_penulis_buku")))
	buku.IDPenerbitBuku = int32(helper.StringkeInt(r.FormValue("id_penerbit_buku")))
	buku.ThnTerbit = r.FormValue("tahun_terbit")
	buku.StokBuku = int32(helper.StringkeInt(r.FormValue("stok_buku")))
	buku.RakBuku = r.FormValue("rak_buku")
	buku.DeskripsiBuku = r.FormValue("deskripsi_buku")
	//gambar disini
	buku.Gambarbuku = "assets/buku/" + uploadGambar(w, r)
	//end gambar
	buku.Kondisibuku = r.FormValue("kondisi_buku")

	updateData, err := buku.UpdateBuku(config.Db, uint32(id))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, responses.Sukses(updateData))
}

func hapusGambar(id uint32, w http.ResponseWriter) {
	buku := BukuModels.Buku{}
	bukus, err := buku.DetailBuku(config.Db, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	r, _ := regexp.Compile("^assets/buku/*")
	namafile := r.ReplaceAllString(bukus.Gambarbuku, "")
	os.Remove("assets/buku/" + namafile)
	// fmt.Println()
	// fmt.Println(bukus.Gambarbuku)

}

func uploadGambar(w http.ResponseWriter, r *http.Request) string {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("gambar_buku")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//this is path which  we want to store the file
	f, err := os.OpenFile("assets/buku/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return ""
	}
	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path
	// return that we have successfully uploaded our file!
	fmt.Println(w, "Successfully Uploaded File\n")

	return handler.Filename
}

func HapusBuku(w http.ResponseWriter, r *http.Request) {
	var buku BukuModels.Buku
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	if id == 0 {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("URL Invalid"))
		return
	}

	// hapus gambar terlebih dahulu
	hapusGambar(uint32(id), w)

	_, err := buku.HapusBuku(config.Db, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusOK, responses.Sukses(fmt.Sprintf("Data dengan id %d sukses terhapus ", id)))
}
