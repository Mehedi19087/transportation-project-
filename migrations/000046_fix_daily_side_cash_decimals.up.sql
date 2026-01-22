ALTER TABLE daily_side_cash 
    ALTER COLUMN cash TYPE DECIMAL(10,2),
    ALTER COLUMN remaining_balance TYPE DECIMAL(10,2),
    ALTER COLUMN trip_cost TYPE DECIMAL(10,2),
    ALTER COLUMN other_cost TYPE DECIMAL(10,2);