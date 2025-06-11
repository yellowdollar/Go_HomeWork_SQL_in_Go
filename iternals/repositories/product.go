package repositories

import (
	"fmt"

	"home_work_sql_gin/iternals/models"

	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) {

	stmt := `
		CREATE TABLE IF NOT EXISTS tb_products(
		id SERIAL PRIMARY KEY,
		name VARCHAR,
		price INT
		)
	`

	db.MustExec(stmt)
	fmt.Println("Table Created!")
}

func AddNewProduct(db *sqlx.DB, p models.Product) (models.Product, error) {
	stmt := `
		INSERT INTO tb_products(name, price)
		VALUES($1, $2)
	`

	_, err := db.Exec(stmt, p.Name, p.Price)

	dataItem := models.Product{
		Name:  p.Name,
		Price: p.Price,
	}

	return dataItem, err
}

func GetAllProducts(db *sqlx.DB) ([]models.Product, error) {
	stmt := `
		SELECT *  FROM tb_products
	`

	var products []models.Product

	err := db.Select(&products, stmt)

	return products, err
}

func GetProductById(productId int, db *sqlx.DB) ([]models.Product, error) {
	stmt := `
		SELECT * FROM tb_products 
		WHERE id = $1
	`

	var product []models.Product

	err := db.Select(&product, stmt, productId)

	return product, err
}

func UpdateProductPriceById(productId int, newPrice int, db *sqlx.DB) error {
	stmt := `
		UPDATE tb_products SET price = $1 WHERE id = $2
	`

	_, err := db.Exec(stmt, newPrice, productId)

	// id, err := strconv.Atoi(idStr)

	// product, err1 := GetProductById(id, db)

	return err
}

func DeleteProductById(productId int, db *sqlx.DB) error {
	stmt := `
		DELETE FROM tb_products WHERE id=$1
	`

	_, err := db.Exec(stmt, productId)

	return err
}
