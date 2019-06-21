/*
package WebServer Runs the main website back-end for WarhammerTools.com
*/
package WebServer

import (
	"fmt"
	"net/http"
)

//LoadWebServer() Starts up the main web server for WarhammerTools.com
func LoadWebServer() (err error) {
	//	http.HandleFunc("/test", Responder)
	fs := http.FileServer(http.Dir("Website/"))
	http.Handle("/", http.StripPrefix("/", fs))
	fmt.Printf("launching web server")
	err = http.ListenAndServe(":80", nil)
	return
}
