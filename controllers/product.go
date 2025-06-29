package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/idstev/marketplace/models"
)


// ProductList muestra la lista de productos del usuario autenticado
// Esta función obtiene el ID del usuario de la sesión y muestra sus productos
func ProductList(w http.ResponseWriter, r *http.Request) {
	userID, _ := GetSessionUserID(r)
	products, _ := models.GetProductsByUser(userID)

	tmpl := template.Must(template.ParseFiles("templates/products.html"))
	tmpl.Execute(w, products)
}
//funcion para crear nuevo producto
// Esta función muestra el formulario para crear un nuevo producto
func NewProductForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/new_product.html"))
	tmpl.Execute(w, nil)
}


// CreateProduct procesa la creación de un nuevo producto
// Esta función obtiene los datos del formulario, crea un nuevo producto y lo guarda en la base de datos
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	// traer el ID del usuario desde la sesión
	//  GetSessionUserID es una función que obtiene el ID del usuario de la sesión
	idStr := r.URL.Query().Get("id")
	userID, _ := strconv.Atoi(idStr)

	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	p := models.Product{
		UserID:      userID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
	}

	models.CreateProduct(p)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}


// EditProductForm muestra el formulario para editar un producto
// Esta función obtiene el ID del producto desde la URL y muestra sus datos en el formulario
func EditProductForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	product, err := models.GetProductByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/edit_product.html"))
	tmpl.Execute(w, product)
}


// UpdateProduct procesa la actualización de un producto
// Esta función obtiene los datos del formulario, actualiza el producto en la base de datos y redirige a la lista de productos
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	product := models.Product{
		ID:          id,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
	}

	models.UpdateProduct(product)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

// DeleteProduct procesa la eliminación de un producto
// Esta función obtiene el ID del producto desde la URL y lo elimina de la base de datos
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_ = models.DeleteProduct(id)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
