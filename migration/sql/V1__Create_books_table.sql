CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY,
    title     TEXT NOT NULL,
	author    text not null,
	rating    float(2),
	date_read  TIMESTAMP,
	date_added TIMESTAMP,
	isbn      text not null,
	isbn13    text not null,
    status text not null
);
