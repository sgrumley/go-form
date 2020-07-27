package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sgrumley/go-form/app"
)

func saveToDB() {

}


func saveToMonday() {

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "pages/home.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() error: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website = %v\n", r.PostForm)
		name := r.FormValue("name")
		fmt.Fprintf(w, "Name = %s\n", name)

	default:
		fmt.Fprintf(w, "Unsupported request\n")
	}
}



func filesPageHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/files.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() error: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website = %v\n", r.PostForm)
		name := r.FormValue("name")
		fmt.Fprintf(w, "Name = %s\n", name)

	default:
		fmt.Fprintf(w, "Unsupported request\n")
	}
}

func main () {

	// Creates a database connection and closing it again before exit.
	// db, err := CreateConnection()
	// defer db.Close()
	// db.DB().SetMaxIdleConns(1)
	// db.DB().SetMaxOpenConns(2)

	// if err != nil {
	// 	log.Fatalf("Could not connect to DB: %v", err)
	// }

	// // Automatically migrates the user struct into database columns/types etc
	// db.AutoMigrate(&User{})

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/form", formPageHandler)
	http.HandleFunc("/files", filesPageHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}