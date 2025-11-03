CREATE TABLE drivers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mobile VARCHAR(15) UNIQUE NOT NULL,
    address TEXT,
    emergency VARCHAR(15),
    license VARCHAR(50) UNIQUE,
    expired DATE,
    nid VARCHAR(20) UNIQUE,
    joining_date DATE,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);