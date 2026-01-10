DROP INDEX IF EXISTS idx_own_vehicle_trips_driver_name;
ALTER TABLE own_vehicle_trips DROP COLUMN IF EXISTS driver_name;