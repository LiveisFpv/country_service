CREATE TABLE IF NOT EXISTS country
(
    country_id SERIAL PRIMARY KEY,
    country_title TEXT NOT NULL UNIQUE,
    country_capital TEXT NOT NULL,
    country_area TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_country_title ON country(country_title);

-- ONLY FOR DEBUG
INSERT INTO country (country_title, country_capital, country_area) VALUES
    ('RUSSIA', 'MOSCOW', 'TOO BIG'),
    ('USA', 'NEW YORK', 'TOO BIG'),
    ('CHINA', 'BEIJING', 'TOO BIG'),
    ('CANADA', 'OTTAWA', 'TOO BIG'),
    ('AUSTRALIA', 'CANBERRA', 'TOO BIG');