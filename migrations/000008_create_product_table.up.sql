create table products(
     id bigserial PRIMARY key,
     name varchar(100) not null,
     created_at timestamptz not null DEFAULT NOW()
);