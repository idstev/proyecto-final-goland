package models

import (
	"github.com/idstev/marketplace/config"
)

func CreateProduct(p Product) error {
	query := `INSERT INTO products (user_id, name, description, price, stock) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := config.DB.Exec(query, p.UserID, p.Name, p.Description, p.Price, p.Stock)
	return err
}

func GetAllProducts() ([]Product, error) {
	query := `SELECT id, user_id, name, description, price, stock FROM products`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(id int) (Product, error) {
	var p Product
	query := `SELECT id, user_id, name, description, price, stock FROM products WHERE id=$1`
	err := config.DB.QueryRow(query, id).Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.Price, &p.Stock)
	return p, err
}

func UpdateProduct(p Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5`
	_, err := config.DB.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.ID)
	return err
}

func DeleteProduct(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id=$1`, id)
	return err
}

func GetProductsByUser(userID int) ([]Product, error) {
    query := `SELECT id, user_id, name, description, price, stock FROM products WHERE user_id=$1`
    rows, err := config.DB.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var p Product
        err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.Price, &p.Stock)
        if err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

