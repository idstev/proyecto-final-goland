package controllers

import (
    "net/http"
    "strconv"
)



func GetSessionUserID(r *http.Request) (int, error) {
    cookie, err := r.Cookie("session_user_id")
    if err != nil {
        return 0, err
    }
    return strconv.Atoi(cookie.Value)
}
