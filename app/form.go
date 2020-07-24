// handler function for the POST /form

// create a struct that has all the values of the form and create a new instance with these values

package main

import (
	"fmt"
)

func formPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/form.html")

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