alter table users 
 drop column if exists status;


alter table users 
 drop constraint if exists check_status_valid;