CREATE TABLE kidneysmart.nutrient (
    id SERIAL PRIMARY KEY,
    id_type INTEGER NOT NULL,
    nutrient TEXT NOT NULL,
    ru_name TEXT NOT NULL,
    en_name TEXT NOT NULL,
    ru_unit TEXT NOT NULL,
    en_unit TEXT NOT NULL
);