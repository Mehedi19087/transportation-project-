ALTER TABLE calans ADD COLUMN IF NOT EXISTS trip_id BIGINT;
ALTER TABLE calans ADD COLUMN IF NOT EXISTS outside_trip_id BIGINT;

CREATE INDEX IF NOT EXISTS idx_calans_trip_id ON calans(trip_id);
CREATE INDEX IF NOT EXISTS idx_calans_outside_trip_id ON calans(outside_trip_id);