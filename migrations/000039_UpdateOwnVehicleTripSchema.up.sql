-- 1. Add new columns
ALTER TABLE own_vehicle_trips ADD COLUMN IF NOT EXISTS date TEXT;
ALTER TABLE own_vehicle_trips ADD COLUMN IF NOT EXISTS outside_trip_id INTEGER;
ALTER TABLE own_vehicle_trips ADD COLUMN IF NOT EXISTS trip_id INTEGER;

-- 2. Ensure trip_id is nullable (in case it was NOT NULL before)
ALTER TABLE own_vehicle_trips ALTER COLUMN trip_id DROP NOT NULL;

-- 3. Add Indexes
CREATE INDEX IF NOT EXISTS idx_own_vehicle_trips_date ON own_vehicle_trips(date);
CREATE INDEX IF NOT EXISTS idx_own_vehicle_trips_trip_id ON own_vehicle_trips(trip_id);
CREATE INDEX IF NOT EXISTS idx_own_vehicle_trips_outside_trip_id ON own_vehicle_trips(outside_trip_id);
