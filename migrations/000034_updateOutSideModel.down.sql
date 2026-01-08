alter table outside_trips 
     drop column if exists due,
     drop column if exists due_status;