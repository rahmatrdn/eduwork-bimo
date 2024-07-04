package models

import (
	"database/sql"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"

	_ "github.com/go-sql-driver/mysql"
)

type ProductModel struct {
	db *sql.DB
}

func NewProductModel(db *sql.DB) *ProductModel {
	return &ProductModel{db: db}
}

func (m ProductModel) GetProductById(productId int) (*entities.Product, error) {
	row := m.db.QueryRow("SELECT product_id, product_name, description, price, stock FROM products WHERE product_id = ?", productId)
	var product entities.Product
	err := row.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m ProductModel) GetAll() ([]entities.Product, error) {
	rows, err := m.db.Query("SELECT product_id, product_name, description, price, stock FROM products")
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	products := []entities.Product{}
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (m ProductModel) GetProductStock(productId int) (int, error) {
	var stock int
	err := m.db.QueryRow("SELECT stock FROM products WHERE product_id = ?", productId).Scan(&stock)
	if err != nil {
		return -1, err
	}

	return stock, nil
}

func (m ProductModel) Create(product *entities.Product) error {
	result, err := m.db.Exec("INSERT INTO products (product_name, description, price, stock) VALUES (?, ?, ?, ?)", product.ProductName, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ProductID = int(id)
	return nil
}

func (m ProductModel) Update(product *entities.Product) error {
	_, err := m.db.Exec("UPDATE products SET product_name = ?, description = ?, price = ?, stock = ? WHERE product_id = ?", product.ProductName, product.Description, product.Price, product.Stock, product.ProductID)
	return err
}

func (m ProductModel) Delete(ProductID int) error {
	_, err := m.db.Exec("DELETE FROM products WHERE product_id = ?", ProductID)
	if err != nil {
		return err
	}
	return nil
}
