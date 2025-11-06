CREATE TABLE IF NOT EXISTS own_vehicles (
    id BIGSERIAL PRIMARY KEY,
    load_point VARCHAR(255) NOT NULL DEFAULT '',
    unload_point VARCHAR(255) NOT NULL DEFAULT '',
    rent DOUBLE PRECISION NOT NULL DEFAULT 0,
    advance DOUBLE PRECISION NOT NULL DEFAULT 0,
    trip_cost DOUBLE PRECISION NOT NULL DEFAULT 0,
    diesel DOUBLE PRECISION NOT NULL DEFAULT 0,
    extra_cost DOUBLE PRECISION NOT NULL DEFAULT 0,
    diesel_taka DOUBLE PRECISION NOT NULL DEFAULT 0,
    pamp VARCHAR(255) NOT NULL DEFAULT '',
    commission DOUBLE PRECISION NOT NULL DEFAULT 0,
    trip_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create unique index on trip_id to prevent duplicates
CREATE UNIQUE INDEX IF NOT EXISTS idx_own_vehicles_trip_id ON own_vehicles(trip_id);

-- Create index on deleted_at for soft delete queries
CREATE INDEX IF NOT EXISTS idx_own_vehicles_deleted_at ON own_vehicles(deleted_at);