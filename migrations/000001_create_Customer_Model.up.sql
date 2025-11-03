create table customers (
     id serial primary key, 
     customer_name varchar(255) not null ,
     mobile varchar(255),
     email varchar(255),
     address text,
     opening_balance varchar(255),
     status varchar(255),
     created_at timestamptz,
     updated_at timestamptz,
     deleted_at timestamptz
);

create index idx_customers_deleted_at on customers(deleted_at);