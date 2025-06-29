package models

type Product struct {
    ID          int
    UserID      int
    Name        string
    Description string
    Price       float64
    Stock       int
}
