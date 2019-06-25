package main

import (
	"fmt"
	"net/http"

	"github.com/houseofbosons/houseofbosons/services/backend/db"
	"github.com/houseofbosons/houseofbosons/services/backend/middleware"

	"github.com/houseofbosons/houseofbosons/services/backend/router"
)

func main() {
	admin := db.Admin{
		Email: "ritwik3@ritwiksaha.com",
		// GoogleID: "1234567890112",
	}

	admin.Insert()
	// admin.Query()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api", http.StatusSeeOther)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to houseofbosons!")
	})

	http.HandleFunc("/api/auth/google", router.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", router.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", middleware.CheckAuth(router.CurrentUser))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
