package controllers

import (
    "net/http"
)

func AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        _, err := GetSessionUserID(r)
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        handler(w, r)
    }
}
