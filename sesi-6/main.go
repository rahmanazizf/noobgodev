package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Fisikaasyik123"
	dbname   = "godev"
)

func main() {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	log.Println("Connected to the database!")

	// creating table
	CreateTable(db)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateTable(db *sql.DB) {
	// tabel products
	createProducts := `
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
