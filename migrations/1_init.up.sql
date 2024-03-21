CREATE TABLE IF NOT EXISTS shelves
(
    id        serial PRIMARY KEY,
    name     VARCHAR(100) NOT NULL
);

CREATE INDEX IF NOT EXISTS index_name_shelf ON shelves(name);

CREATE TABLE IF NOT EXISTS products
(
    id        serial PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    id_main_shelf INTEGER NOT NULL REFERENCES shelves(id)
);

CREATE TABLE IF NOT EXISTS products_shelves
(
    id_product INTEGER REFERENCES products(id),
    id_shelf INTEGER REFERENCES shelves(id),
    PRIMARY KEY (id_product, id_shelf)
);

CREATE TABLE IF NOT EXISTS orders
(
    id        serial PRIMARY KEY,
    num     INTEGER NOT NULL,
    count_product INTEGER NOT NULL,
    id_product INTEGER NOT NULL REFERENCES products(id)
);

CREATE INDEX IF NOT EXISTS index_num_order ON orders(num);
