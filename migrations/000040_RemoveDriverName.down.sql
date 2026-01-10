ALTER TABLE own_vehicle_trips ADD COLUMN IF NOT EXISTS driver_name TEXT;
CREATE INDEX IF NOT EXISTS idx_own_vehicle_trips_driver_name ON own_vehicle_trips(driver_name);