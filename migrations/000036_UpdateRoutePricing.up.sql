ALTER TABLE route_pricings RENAME COLUMN customer_name TO dealer_name;
   ALTER TABLE route_pricings RENAME COLUMN unload_point TO destination;

   -- Change rate type from INT to FLOAT (DOUBLE PRECISION in PostgreSQL)
   ALTER TABLE route_pricings ALTER COLUMN rate TYPE DECIMAL(10, 2);

   -- Remove unwanted columns
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS size;
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS load_point;
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS weight;
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS notes;
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS status;
   ALTER TABLE route_pricings DROP COLUMN IF EXISTS vehicle_category;