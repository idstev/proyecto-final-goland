package config

import (
    "database/sql"
    _"github.com/lib/pq"
    "log"
)

var DB *sql.DB

func Connect() {
    var err error
    connStr := "host=localhost port=5432 user=postgres password=ivChFtg^9sxlix4ZvG dbname=marketplace sslmode=disable"

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("No se pudo establecer la conexión:", err)
    }

    log.Println("✅ Conexión exitosa a PostgreSQL")
}
