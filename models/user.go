package models

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
    Role     string // buyer o seller
}
