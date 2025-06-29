package controllers

import (
	"encoding/json"
	"github.com/idstev/marketplace/models"
	"net/http"
	"strconv"
	"strings"
)

// ----- 1. Registro API -----
func ApiRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := models.CreateUser(user)
	if err != nil {
		http.Error(w, "Error al registrar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario creado correctamente"})
}

// ----- 2. Login API -----
func ApiLogin(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	user, err := models.GetUserByEmailAndPassword(data["email"], data["password"])
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_user_id",
		Value: strconv.Itoa(user.ID),
		Path:  "/",
	})

	json.NewEncoder(w).Encode(map[string]string{"message": "Login exitoso"})
}

// ----- 3. Obtener usuario actual -----
func ApiMe(w http.ResponseWriter, r *http.Request) {
	userID, err := GetSessionUserID(r)
	if err != nil {
		http.Error(w, "No autenticado", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"user_id": userID})
}

// ----- 4. Listar productos -----
func ApiListProducts(w http.ResponseWriter, r *http.Request) {
	userID, _ := GetSessionUserID(r)
	products, _ := models.GetProductsByUser(userID)
	json.NewEncoder(w).Encode(products)
}

// ----- 5. Obtener producto por ID -----
func ApiGetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)
	product, err := models.GetProductByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// ----- 6. Crear producto -----
func ApiCreateProduct(w http.ResponseWriter, r *http.Request) {
	userID, _ := GetSessionUserID(r)
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	p.UserID = userID

	err := models.CreateProduct(p)
	if err != nil {
		http.Error(w, "Error al crear producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto creado"})
}

// ----- 7. Actualizar producto -----
func ApiUpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)

	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	p.ID = id

	err := models.UpdateProduct(p)
	if err != nil {
		http.Error(w, "Error al actualizar", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Producto actualizado"})
}

// ----- 8. Eliminar producto -----
func ApiDeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, _ := strconv.Atoi(idStr)

	err := models.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Error al eliminar", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado"})
}
