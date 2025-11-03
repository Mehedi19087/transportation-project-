-- Drop the TEXT[] column
ALTER TABLE products DROP COLUMN IF EXISTS trip_fields;

-- Add back JSONB column
ALTER TABLE products ADD COLUMN trip_fields JSONB DEFAULT '[]'::jsonb;