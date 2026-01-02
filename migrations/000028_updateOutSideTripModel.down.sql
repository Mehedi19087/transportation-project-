-- Drop indexes first
DROP INDEX IF EXISTS idx_outside_trips_product_id;
DROP INDEX IF EXISTS idx_outside_trips_deleted_at;

-- Drop columns
ALTER TABLE outside_trips 
DROP COLUMN IF EXISTS product_id;

ALTER TABLE outside_trips 
DROP COLUMN IF EXISTS deleted_at;