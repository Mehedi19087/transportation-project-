ALTER TABLE vehicles
    DROP COLUMN IF EXISTS image_url,
    DROP COLUMN IF EXISTS permit_expiry,
    DROP COLUMN IF EXISTS tax_token_expiry,
    DROP COLUMN IF EXISTS fitness_expiry,
    DROP COLUMN IF EXISTS insurance_expiry,
    DROP COLUMN IF EXISTS registration_date,
    DROP COLUMN IF EXISTS fuel_capacity,
    DROP COLUMN IF EXISTS mileage,
    DROP COLUMN IF EXISTS year,
    DROP COLUMN IF EXISTS model,
    DROP COLUMN IF EXISTS brand,
    DROP COLUMN IF EXISTS driver_contact,
    DROP COLUMN IF EXISTS driver_name;