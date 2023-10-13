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
		quantity int CHECK(quantity >= 0),
		product_id int REFERENCES products(id),
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
