alter table users 
     add column if not exists product_id bigint not null;

create index idx_users_product_id on users(product_id)