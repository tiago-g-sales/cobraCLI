package database

import (
	"database/sql"
  _ "github.com/mattn/go-sqlite3"
  	"github.com/tiago-g-sales/cobraCLI/internal/database"
)


func getDB() *sql.DB{
	db, err := sql.Open("sqlite3","./database.db" )
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}
