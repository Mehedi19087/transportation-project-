alter table daily_side_cash 
      add column if not exists without_remaining decimal(10,2);