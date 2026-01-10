ALTER TABLE outside_trips ADD COLUMN date VARCHAR(25);



create index idx_outside_trips_date on outside_trips(date);