CREATE TABLE IF NOT EXISTS bill_statuses (
    id SERIAL PRIMARY KEY,
    date varchar(30),
    company_name TEXT,
    bill_amount decimal(10,2),
    vat_status TEXT,
    status TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);