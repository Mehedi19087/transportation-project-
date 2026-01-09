ALTER TABLE route_pricings ADD COLUMN size INT NOT NULL DEFAULT 0;
   ALTER TABLE route_pricings ADD COLUMN load_point VARCHAR(255) NOT NULL DEFAULT '';
   ALTER TABLE route_pricings ADD COLUMN weight DECIMAL(10,2);
   ALTER TABLE route_pricings ADD COLUMN notes TEXT;
   ALTER TABLE route_pricings ADD COLUMN status VARCHAR(50);
   ALTER TABLE route_pricings ADD COLUMN vehicle_category VARCHAR(255);

   -- Revert rate type (Note: precision may be lost when casting back to INT)
   ALTER TABLE route_pricings ALTER COLUMN rate TYPE INT;

   -- Rename columns back to their original names
   ALTER TABLE route_pricings RENAME COLUMN dealer_name TO customer_name;
   ALTER TABLE route_pricings RENAME COLUMN destination TO unload_point;