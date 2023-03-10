package repository

import (
	"database/sql"

	"github.com/Julio-Norberto/api-message/internal/entity"
)

// minha instancia de banco de dados
type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("Insert into products (id, name, price) values(?, ?, ?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}
