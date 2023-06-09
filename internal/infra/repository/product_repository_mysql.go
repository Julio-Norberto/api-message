package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

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

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product // atribuindo tipo
	for rows.Next() {
		var product entity.Product // atribuindo tipo
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	return products, nil
}
