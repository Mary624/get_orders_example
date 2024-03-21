package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"test_orders/internal/config"
	"test_orders/internal/service"
	"test_orders/internal/storage"
)

type Storage struct {
	db *sql.DB
}

func New(cfg config.Config) (*Storage, error) {
	const op = "storage.New"

	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.HostDB, cfg.PortDB, cfg.UserDB, cfg.PassDB, cfg.DBName)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetOrders(orders []int64) ([]storage.Product, error) {
	rows, err := s.db.Query(getOrderQuery(orders))

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, storage.ErrNoResult
		}
		return nil, err
	}

	res := make([]storage.Product, 0, 200)
	for rows.Next() {
		var product storage.Product
		rows.Scan(&product.MainShelf, &product.NameProduct, &product.IdProduct, &product.Order,
			&product.CountProduct, &product.AdditionalSheves)
		res = append(res, product)
	}

	return res, nil
}

func getOrderQuery(orders []int64) string {
	return fmt.Sprintf(`WITH main AS (SELECT s."name" AS shelf_name, p."name" AS product_name, p.id AS product_id, o.num AS order_num,
	o.count_product AS count_product, s.id AS shelf_id, p.id_main_shelf AS product_main_shelf_id
	FROM orders o
	INNER JOIN products p ON o.id_product  = p.id
	INNER JOIN products_shelves pos ON p.id = pos.id_product
	INNER JOIN shelves s ON pos.id_shelf = s.id
	WHERE o.num IN (%s)),

	a AS (SELECT m.product_id AS add_id, array_agg(m.shelf_name) AS add_shelves
	FROM main m
	WHERE m.shelf_id <> m.product_main_shelf_id
	GROUP BY m.product_id)

	SELECT m.shelf_name, m.product_name, m.product_id, m.order_num, m.count_product, a.add_shelves 
	FROM main m
	LEFT JOIN a ON a.add_id = m.product_id
	WHERE m.shelf_id = m.product_main_shelf_id
	ORDER BY m.shelf_name, m.order_num, m.product_id;`, service.GetStringOrders(orders))
}
