package controller

import (
	"log"
	"net/http"

	"github.com/WuShaoQiang/mega/vm"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println("middle:", username)
		if username != "" {
			log.Println("Last seen:", username)
			vm.UpdateLastSeen(username)
		}
		if err != nil {
			log.Println("middle get session error and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
