CREATE TABLE IF NOT EXISTS country
(
    country_id SERIAL PRIMARY KEY,
    country_title TEXT NOT NULL UNIQUE,
    country_capital TEXT NOT NULL,
    country_area TEXT NOT NULL,
)

-- ONLY FOR DEBUG
INSERT INTO country
(
    "RUSSIA",
    "MOSCOW",
    "TOO BIG"
)
INSERT INTO country
(
    "USA",
    "NEW-YORK",
    "TOO BIG"
)
INSERT INTO country
(
    "CHINA",
    "PEKIN",
    "TOO BIG"
)
INSERT INTO country
(
    "CANADA",
    "OTTAWA",
    "TOO BIG"
)
INSERT INTO country
(
    "AUSTRALIA",
    "CANBERRA",
    "TOO BIG"
)