CREATE TABLE IF NOT EXISTS kidneysmart.type_nutrient (
    id SERIAL PRIMARY KEY,
    id_type_nutrient INT NOT NULL,
    ru_name TEXT NOT NULL,
    en_name TEXT NOT NULL
);