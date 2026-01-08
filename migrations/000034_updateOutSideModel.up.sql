alter table outside_trips 
      add column if not exists due decimal(10,2),
      add column if not exists due_status varchar(10) default 'unpaid';