-- purchases.up.sql
CREATE TABLE IF NOT EXISTS purchases (
    id            BIGSERIAL PRIMARY KEY,
    supplier_name TEXT,
    order_date    TIMESTAMPTZ,
    delivery_date TIMESTAMPTZ,
    items         TEXT,
    quantity      TEXT,
    category      TEXT,
    total_amount  DOUBLE PRECISION,
    vehicle_no    TEXT,
    notes         TEXT,
    payment_status TEXT,
    bill_image     TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMPTZ
);

-- Soft-delete index
CREATE INDEX IF NOT EXISTS idx_purchases_deleted_at ON purchases (deleted_at);