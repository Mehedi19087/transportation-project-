

drop index if exists idx_outside_trips_date;

ALTER TABLE outside_trips DROP COLUMN date;