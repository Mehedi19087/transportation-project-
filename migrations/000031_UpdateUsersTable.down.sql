alter table users 
    drop column if exists product_id;

drop index if exists idx_users_product_id;