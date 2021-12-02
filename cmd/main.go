package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/velero-poc/pkg"
	"net/http"
)

func main() {
	fmt.Println("Initialize velero poc")
	r := mux.NewRouter()
	r.HandleFunc("/backup", pkg.CreateBackup).Methods("GET")

	http.ListenAndServe(":8080", r)
}
