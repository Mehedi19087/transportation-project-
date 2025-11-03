CREATE TABLE vehicles (
    id BIGSERIAL PRIMARY KEY,
    driver_id BIGINT,
    vehicle_name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    size VARCHAR(50),
    vehicle_no VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    joining_date DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_driver
        FOREIGN KEY (driver_id)
        REFERENCES drivers(id)
        ON DELETE SET NULL
);