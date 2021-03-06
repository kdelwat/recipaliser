-- +goose Up
CREATE TABLE IF NOT EXISTS ingredients (
    name TEXT NOT NULL PRIMARY KEY,
    ausnut_id TEXT NOT NULL,
    energy_with_dietary_fibre REAL NOT NULL,
    energy_without_dietary_fibre REAL NOT NULL,
    moisture REAL NOT NULL,
    protein REAL NOT NULL,
    total_fat REAL NOT NULL,
    available_carbohydrates_with_sugar_alcohols REAL NOT NULL,
    available_carbohydrates_without_sugar_alcohol REAL NOT NULL,
    starch REAL NOT NULL,
    total_sugars REAL NOT NULL,
    added_sugars REAL NOT NULL,
    free_sugars REAL NOT NULL,
    dietary_fibre REAL NOT NULL,
    alcohol REAL NOT NULL,
    ash REAL NOT NULL,
    preformed_vitamin_a_retinol REAL NOT NULL,
    beta_carotene REAL NOT NULL,
    provitamin_a_beta_carotene_equivalents REAL NOT NULL,
    vitamin_a_retinol_equivalents REAL NOT NULL,
    thiamin_b1 REAL NOT NULL,
    riboflavin_b2 REAL NOT NULL,
    niacin_b3 REAL NOT NULL,
    niacin_derived_equivalents REAL NOT NULL,
    folate_natural REAL NOT NULL,
    folic_acid REAL NOT NULL,
    total_folates REAL NOT NULL,
    dietary_folate_equivalents REAL NOT NULL,
    vitamin_b6 REAL NOT NULL,
    vitamin_b12 REAL NOT NULL,
    vitamin_c REAL NOT NULL,
    alpha_tocopherol REAL NOT NULL,
    vitamin_e REAL NOT NULL,
    calcium_ca REAL NOT NULL,
    iodine_i REAL NOT NULL,
    iron_fe REAL NOT NULL,
    magnesium_mg REAL NOT NULL,
    phosphorus_p REAL NOT NULL,
    potassium_k REAL NOT NULL,
    selenium_se REAL NOT NULL,
    sodium_na REAL NOT NULL,
    zinc_zn REAL NOT NULL,
    caffeine REAL NOT NULL,
    cholesterol REAL NOT NULL,
    tryptophan REAL NOT NULL,
    total_saturated_fat REAL NOT NULL,
    total_monounsaturated_fat REAL NOT NULL,
    total_polyunsaturated_fat REAL NOT NULL,
    linoleic_acid REAL NOT NULL,
    alphalinolenic_acid REAL NOT NULL,
    c205w3_eicosapentaenoic REAL NOT NULL,
    c225w3_docosapentaenoic REAL NOT NULL,
    c226w3_docosahexaenoic REAL NOT NULL,
    total_long_chain_omega3_fatty_acids REAL NOT NULL,
    total_trans_fatty_acids REAL NOT NULL
);

CREATE TABLE recipes (
    name TEXT NOT NULL PRIMARY KEY
);

CREATE TABLE recipe_ingredients (
    recipe_name TEXT NOT NULL,
    ingredient_name TEXT NOT NULL,
    amount REAL NOT NULL,
    FOREIGN KEY(recipe_name) REFERENCES recipes(name),
    FOREIGN KEY(ingredient_name) REFERENCES ingredients(name)
);

-- +goose Down
DROP TABLE ingredients;
DROP TABLE recipes;
DROP TABLE recipe_ingredients;