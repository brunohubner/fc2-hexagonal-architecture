package db

import (
	"database/sql"
	"fmt"

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
	var rows int
	p.db.QueryRow(`
		select count(id) from products where id = ?
	`, product.GetID()).Scan(&rows)
	fmt.Print("ROWS", rows)
	if rows == 0 {
		if _, err := p.create(product); err != nil {
			return nil, err
		}
	} else if _, err := p.update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare(`
		insert into products (id, name, price, status) values(?, ?, ?, ?);
	`)
	if err != nil {
		return nil, err
	}

	if _, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	); err != nil {
		return nil, err
	}

	if err = stmt.Close(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	if _, err := p.db.Exec(`
		update products set name = ?, price = ?, status = ? where id = ?;
	`,
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	); err != nil {
		return nil, err
	}

	return product, nil
}
