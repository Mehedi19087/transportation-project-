create table dealers (
     id bigserial PRIMARY key,
     name varchar(255) not null ,
     destination varchar(255) not null,
     status varchar(255),
     created_at TIMESTAMPTZ,
     updated_at TIMESTAMPTZ
);