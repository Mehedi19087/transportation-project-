CREATE TABLE calans (
    id BIGSERIAL PRIMARY KEY,
    date TEXT,
    vehicle_no VARCHAR(100),
    destination VARCHAR(255),
    amount DECIMAL(10,2),
    status VARCHAR(50) DEFAULT 'unpaid',
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_calans_deleted_at ON calans(deleted_at);