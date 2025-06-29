package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/idstev/marketplace/models"
)

// Formulario de registro
func RegisterForm(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/register.html"))
    tmpl.Execute(w, nil)
}

// Procesar registro
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/register", http.StatusSeeOther)
        return
    }

    user := models.User{
        Name:     r.FormValue("name"),
        Email:    r.FormValue("email"),
        Password: r.FormValue("password"), // ⚠️ En producción, proccesar la contraseña por cifrado
        Role:     r.FormValue("role"),
    }

    err := models.CreateUser(user)
    if err != nil {
        http.Error(w, "Error al registrar usuario", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Formulario de login
func LoginForm(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/login.html"))
    tmpl.Execute(w, nil)
}

// Procesar login
func LoginUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    user, err := models.GetUserByEmailAndPassword(email, password)
    if err != nil {
        http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
        return
    }

    // Crear cookie de sesión
    http.SetCookie(w, &http.Cookie{
        Name:  "session_user_id",
        Value: strconv.Itoa(user.ID),
        Path:  "/",
    })

    http.Redirect(w, r, "/products", http.StatusSeeOther)
}

