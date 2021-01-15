package routes

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/controller"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/controller/BukuController"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/seed", migrate.MigrateKeDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", middlewares.RenderKeJSON(controller.Register)).Methods("POST", "OPTIONS")
	router.HandleFunc("/loginadmin", middlewares.RenderKeJSON(controller.LoginAdmin)).Methods("POST")
	router.HandleFunc("/api/home", middlewares.RenderKeJSON(middlewares.HarusAuth(controller.Home))).Methods("GET")
	router.HandleFunc("/api/refresh", middlewares.RenderKeJSON(middlewares.HarusAuth(controller.RefreshToken))).Methods("POST")

	// route BUKU
	router.HandleFunc("/api/jenisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.LihatJenisBuku))).Methods("GET")
	router.HandleFunc("/api/jenisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.UpdateJenisBuku))).Methods("PUT")
	router.HandleFunc("/api/jenisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.TambahJenisBuku))).Methods("POST")
	router.HandleFunc("/api/jenisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.HapusJenisBuku))).Methods("DELETE")

	router.HandleFunc("/api/penulisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.LihatPenulisBuku))).Methods("GET")
	router.HandleFunc("/api/penulisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.TambahPenulisBuku))).Methods("POST")
	router.HandleFunc("/api/penulisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.UpdatePenulisBuku))).Methods("PUT")
	router.HandleFunc("/api/penulisbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.HapusPenulisBuku))).Methods("DELETE")

	router.HandleFunc("/api/penerbitbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.LihatPenerbitBuku))).Methods("GET")
	router.HandleFunc("/api/penerbitbuku", middlewares.RenderKeJSON(middlewares.HarusAuth(BukuController.TambahPenerbitBuku))).Methods("POST")
	return router
}
