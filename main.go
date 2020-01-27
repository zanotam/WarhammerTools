package WarhammerTools

import (
	"WarhammerTools/src/Servers/Database"
	"WarhammerTools/src/Servers/WebServer"
	"fmt"
	"log"
)

//import (
//	"Database"
//)

func main() {

	err := Database.LoadKTDB()
	if err != nil {
		_, _ = fmt.Printf("There was an error: %v", err)
	}
	err = WebServer.LoadWebServer()
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

/*func Responder (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}*/

/*
https://golang.org/doc/effective_go.html#interface_methods
*/
