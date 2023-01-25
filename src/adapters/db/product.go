package db

import (
	"database/sql"

	"github.com/brunohubner/fc2-hexagonal-architecture/src/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
}

func (p *ProductDb) Get(id string) (application.IProduct, error) {
	var product application.Product

	stmt, err := p.db.Prepare(`
		select id, name, price, status from products where id = ?
	`)
	if err != nil {
		return nil, err
	}

	if err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	return &application.Product{}, nil
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	return &application.Product{}, nil
}
