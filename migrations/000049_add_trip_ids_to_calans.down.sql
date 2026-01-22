DROP INDEX IF EXISTS idx_calans_outside_trip_id;
DROP INDEX IF EXISTS idx_calans_trip_id;

ALTER TABLE calans DROP COLUMN IF EXISTS outside_trip_id;
ALTER TABLE calans DROP COLUMN IF EXISTS trip_id;