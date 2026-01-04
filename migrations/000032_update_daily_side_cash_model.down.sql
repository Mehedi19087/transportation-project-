ALTER TABLE daily_side_cash
DROP INDEX IF EXISTS idx_daily_side_cash_product_id;

ALTER TABLE daily_side_cash
DROP COLUMN IF EXISTS product_id,
DROP COLUMN IF EXISTS cash,
DROP COLUMN IF EXISTS remaining_balance,
DROP COLUMN IF EXISTS trip_cost,
DROP COLUMN IF EXISTS other_cost,
DROP COLUMN IF EXISTS other_cost_details;

ALTER TABLE daily_side_cash 
ADD COLUMN suzuki numeric,
ADD COLUMN yamaha numeric,
ADD COLUMN honda numeric,
ADD COLUMN hatim_rupgonj numeric,
ADD COLUMN rakib_benapole numeric,
ADD COLUMN mofiz_benapole numeric,
ADD COLUMN aziz numeric,
ADD COLUMN shongram numeric;
