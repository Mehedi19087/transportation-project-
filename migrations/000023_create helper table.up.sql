CREATE TABLE helpers (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  mobile TEXT NOT NULL,
  emergency TEXT,
  address TEXT,
  salary TEXT,
  assigned_vehicle TEXT,
  joining_date TIMESTAMPTZ,
  experience TEXT,
  nid TEXT,
  status TEXT NOT NULL DEFAULT 'Active',
  image TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  deleted_at TIMESTAMPTZ
);
CREATE INDEX ix_helpers_deleted_at ON helpers(deleted_at);
