CREATE TABLE IF NOT EXISTS catalogs(
    id BIGSERIAL PRIMARY KEY,
    alias TEXT UNIQUE NOT NULL,
    img TEXT UNIQUE NOT NULL,
    ru_name TEXT UNIQUE NOT NULL,
    addition_date TIMESTAMPTZ DEFAULT now()
);
CREATE TABLE IF NOT EXISTS banners(
    id BIGSERIAL PRIMARY KEY,
    alias TEXT UNIQUE NOT NULL,
    img TEXT UNIQUE NOT NULL,
    redirect_url TEXT,
    addition_date TIMESTAMPTZ DEFAULT now()
);
CREATE TABLE IF NOT EXISTS parameters(
    id BIGSERIAL PRIMARY KEY,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    ru_key TEXT NOT NULL,
    ru_value TEXT NOT NULL,
    catalogs_id BIGINT NOT NULL REFERENCES catalogs(id),
    addition_date TIMESTAMPTZ DEFAULT now(),
    UNIQUE (key, value, ru_key, ru_value),
    UNIQUE (key, value),
    UNIQUE (ru_key, ru_value)
);
