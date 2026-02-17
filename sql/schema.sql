-- TODO icon/picture for each category(additional field)
CREATE TABLE IF NOT EXISTS catalogs(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    ru_name VARCHAR(50) UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS catalog_parameters(
    KEY VARCHAR(50) PRIMARY KEY,
    VALUE VARCHAR(50) NOT NULL,
    ru_key VARCHAR(50) PRIMARY KEY,
    ru_value VARCHAR(50) NOT NULL,
    catalog_id INT REFERENCES catalogs(id)
);
