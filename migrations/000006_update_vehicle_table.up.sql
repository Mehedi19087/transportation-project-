ALTER TABLE vehicles
    ADD COLUMN driver_name VARCHAR(255),
    ADD COLUMN driver_contact VARCHAR(20),
    ADD COLUMN brand VARCHAR(100),
    ADD COLUMN model VARCHAR(100),
    ADD COLUMN year INT,
    ADD COLUMN mileage INT,
    ADD COLUMN fuel_capacity INT,
    ADD COLUMN registration_date DATE,
    ADD COLUMN insurance_expiry DATE,
    ADD COLUMN fitness_expiry DATE,
    ADD COLUMN tax_token_expiry DATE,
    ADD COLUMN permit_expiry DATE,
    ADD COLUMN image_url TEXT;