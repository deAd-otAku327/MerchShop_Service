CREATE TABLE IF NOT EXISTS "items"
(
    "id" SERIAL PRIMARY KEY,
    "type" VARCHAR(15) NOT NULL UNIQUE,
    "price" INTEGER NOT NULL
);