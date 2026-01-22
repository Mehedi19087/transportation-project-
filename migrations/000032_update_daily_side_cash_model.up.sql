ALTER TABLE daily_side_cash 
DROP COLUMN IF EXISTS suzuki,
DROP COLUMN IF EXISTS yamaha,
DROP COLUMN IF EXISTS honda,
DROP COLUMN IF EXISTS hatim_rupgonj,
DROP COLUMN IF EXISTS rakib_benapole,
DROP COLUMN IF EXISTS mofiz_benapole,
DROP COLUMN IF EXISTS aziz,
DROP COLUMN IF EXISTS shongram;

ALTER TABLE daily_side_cash
ADD COLUMN product_id bigint,
ADD COLUMN cash decimal(10,2),
ADD COLUMN remaining_balance decimal(10,2),
ADD COLUMN trip_cost decimal(10,2),
ADD COLUMN other_cost decimal(10,2),
ADD COLUMN other_cost_details text;

CREATE INDEX idx_daily_side_cash_product_id ON daily_side_cash(product_id);
