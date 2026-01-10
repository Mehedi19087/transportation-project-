DROP INDEX IF EXISTS idx_own_vehicle_trips_outside_trip_id;
DROP INDEX IF EXISTS idx_own_vehicle_trips_trip_id;
DROP INDEX IF EXISTS idx_own_vehicle_trips_date;

ALTER TABLE own_vehicle_trips DROP COLUMN IF EXISTS outside_trip_id;
ALTER TABLE own_vehicle_trips DROP COLUMN IF EXISTS date;
-- We do not drop trip_id because it might have existed before, but we leave it nullable.
