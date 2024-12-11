package productmodel

import (
	"go-web/config"
	"go-web/entities"
	"time"
)

func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			users.name as user_name,
			products.quantity, 
			products.description, 
			products.created_at, 
			products.updated_at 
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN users ON products.user_id = users.id  -- Corrected here
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.User.Name, // Ensure User is populated
			&product.Quantity,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := config.DB.Exec(`
		INSERT INTO products(
			name, category_id, user_id, quantity, description, created_at, updated_at  -- Corrected here
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.User.Id, // Corrected here
		product.Quantity,
		product.Description,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			users.name as user_name,
			products.quantity, 
			products.description, 
			products.created_at, 
			products.updated_at 
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN users ON products.user_id = users.id
		WHERE products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.User.Name,
		&product.Quantity,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) (bool, error) {
	query, err := config.DB.Exec(`
		UPDATE products SET 
			name = ?, 
			category_id = ?,
			user_id = ?,
			quantity = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?`,
		product.Name,
		product.Category.Id,
		product.User.Id,
		product.Quantity,
		product.Description,
		time.Now(),
		id,
	)

	if err != nil {
		return false, err
	}

	result, err := query.RowsAffected()
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
