package cmd

import (
	"database/sql"
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateTable(db *sql.DB) {
	// tabel products
	createProducts := `
	DROP TABLE IF EXISTS variants; -- products dirujuk oleh variants jadi harus didrop dulu
	DROP TABLE IF EXISTS products;
	CREATE TABLE products (
		id serial PRIMARY KEY,
		name varchar(20) NOT NULL UNIQUE,
		created_at timestamp DEFAULT current_timestamp,
		updated_at timestamp
	);
	-- Membuat trigger yang mengisi updated_at
	CREATE OR REPLACE FUNCTION update_timestamp()
	RETURNS TRIGGER AS $$
	BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	-- Menambahkan trigger ke tabel "products"
	CREATE TRIGGER products_updated_at
	BEFORE UPDATE ON products
	FOR EACH ROW
	EXECUTE FUNCTION update_timestamp();
	`
	_, err := db.Query(createProducts)
	CheckError(err)

	// tabel variants
	createVariants := `
	DROP TABLE IF EXISTS variants;
	CREATE TABLE variants (
		id serial PRIMARY KEY,
		variant_name varchar(20) NOT NULL UNIQUE,
		quantity int CHECK(quantity >= 0) NOT NULL,
		product_id int REFERENCES products(id) NOT NULL,
		created_at timestamp DEFAULT current_timestamp,
		updated_at timestamp
	);
	-- Membuat trigger yang mengisi updated_at
	CREATE OR REPLACE FUNCTION update_timestamp()
	RETURNS TRIGGER AS $$
	BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	-- Menambahkan trigger ke tabel "variants"
	CREATE TRIGGER products_updated_at
	BEFORE UPDATE ON variants
	FOR EACH ROW
	EXECUTE FUNCTION update_timestamp();
	`
	_, err = db.Query(createVariants)
	CheckError(err)
	log.Println("Successfully created products and variants table!")
}

func CreateProduct(db *sql.DB, name string) {
	createProduct := `INSERT INTO products (name) VALUES ($1)`
	_, err := db.Query(createProduct, name)
	CheckError(err)
	fmt.Println(fmt.Sprintf("Inserted product: %s", name))
}

func UpdateProduct(db *sql.DB, id int, newName string) {
	updateProduct := `
	UPDATE products
	SET name = $1
	WHERE id = $2;
	`
	_, err := db.Query(updateProduct, id, newName)
	CheckError(err)
	fmt.Println(fmt.Sprintf("Updated product with id %d: %s", id, newName))
}

func GetProductById(db *sql.DB, id int) {
	getProduct := `
	SELECT *
	FROM products
	WHERE id = $1;
	`
	row, err := db.Query(getProduct, id)
	CheckError(err)
	fmt.Println(row)
}

func UpdateVariantById(db *sql.DB, id int, variantName string, productId int) {
	updateVariant := `
	UPDATE variants
	SET variant_name = $1, product_id = $2
	WHERE id = $3;
	`
	_, err := db.Query(updateVariant, variantName, productId, id)
	CheckError(err)
	fmt.Println(fmt.Sprintf("Updated variant with id %d: %s; product id: %d", id, variantName, productId))
}

func CreateVariant(db *sql.DB, variantName string, productId int) {
	createVariant := `
	INSERT INTO variants (variant_name, product_id)
	VALUES ($1, $2) RETURNING id;
	`
	_, err := db.Query(createVariant, variantName, productId)
	CheckError(err)
	fmt.Println(fmt.Sprintf("Created variant with"))
}
