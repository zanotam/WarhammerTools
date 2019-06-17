package main

import (
	"WarhammerTools/KTRoster/ktDatabasePrep"
	"WarhammerTools/WebServer"
	"fmt"
	"log"
)

//import (
//	"ktDatabasePrep"
//)

func main() {

	err := ktDatabasePrep.LoadKTDB()
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
