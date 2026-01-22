ALTER TABLE daily_side_cash 
    ALTER COLUMN cash TYPE NUMERIC,
    ALTER COLUMN remaining_balance TYPE NUMERIC,
    ALTER COLUMN trip_cost TYPE NUMERIC,
    ALTER COLUMN other_cost TYPE NUMERIC;