package main

import (
	"github.com/idstev/marketplace/config"
	"github.com/idstev/marketplace/controllers"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.RegisterForm(w, r)
		} else {
			controllers.RegisterUser(w, r)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.LoginForm(w, r)
		} else {
			controllers.LoginUser(w, r)
		}
	})

	http.HandleFunc("/products", controllers.AuthMiddleware(controllers.ProductList))

	http.HandleFunc("/products/new", controllers.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.NewProductForm(w, r)
		} else {
			controllers.CreateProduct(w, r)
		}
	}))

	http.HandleFunc("/products/edit", controllers.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.EditProductForm(w, r)
		} else {
			controllers.UpdateProduct(w, r)
		}
	}))

	http.HandleFunc("/products/delete", controllers.AuthMiddleware(controllers.DeleteProduct))

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:   "session_user_id",
			Value:  "",
			Path:   "/",
			MaxAge: -1, // eliminar cookie
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.HandleFunc("/api/register", controllers.ApiRegister)
	http.HandleFunc("/api/login", controllers.ApiLogin)
	http.HandleFunc("/api/me", controllers.AuthMiddleware(controllers.ApiMe))

	http.HandleFunc("/api/products", controllers.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ApiListProducts(w, r)
		case http.MethodPost:
			controllers.ApiCreateProduct(w, r)
		}
	}))

	http.HandleFunc("/api/products/", controllers.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ApiGetProduct(w, r)
		case http.MethodPut:
			controllers.ApiUpdateProduct(w, r)
		case http.MethodDelete:
			controllers.ApiDeleteProduct(w, r)
		}
	}))

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
