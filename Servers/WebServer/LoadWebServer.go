package WebServer

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

//LoadWebServer() Starts up the main web server for WarhammerTools.com
func LoadWebServer() (err error) {
	//	http.HandleFunc("/test", Responder)
	//	r := mux.NewRouter()
	static := http.FileServer(http.Dir("Static/")) //currently just the index
	//	kt := http.FileServer(http.Dir("Static/KT"))
	http.Handle("/", http.StripPrefix("/", static)) //seems to work only for homepage atm for some reason
	fmt.Printf("launching web server")
	err = http.ListenAndServe(":80", nil)
	return
}
