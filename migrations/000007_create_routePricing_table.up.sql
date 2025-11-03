create table route_pricings(
     id bigserial PRIMARY key, 
     customer_name varchar(255) not null,
     size int not null,
     load_point varchar(255) not null,
     unload_point varchar(255) not null,
     rate int not null,
     weight DECIMAL(10,2),
     notes text,
     status varchar(50),
     created_at TIMESTAMPTZ not null DEFAULT NOW(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()

)