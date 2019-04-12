package main

import (
	"net/http"

	"github.com/WuShaoQiang/mega/controller"
	"github.com/WuShaoQiang/mega/model"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	// http.ListenAndServe(":8000", nil)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", nil)
}
