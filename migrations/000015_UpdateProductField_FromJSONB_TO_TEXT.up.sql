-- Drop the old JSONB column
ALTER TABLE products DROP COLUMN IF EXISTS trip_fields;

-- Add new TEXT[] column
ALTER TABLE products ADD COLUMN trip_fields TEXT[] DEFAULT ARRAY[]::TEXT[];