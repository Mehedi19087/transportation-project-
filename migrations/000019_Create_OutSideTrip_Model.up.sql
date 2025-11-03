create table outside_trips (
     id BIGSERIAL PRIMARY KEY,
     load_point text ,
     unload_point text,
     rent DECIMAL(10,2),
     advance DECIMAL(10,2),
     trip_cost DECIMAL(10,2),
     diesel DECIMAL(10,2),
     extra_cost DECIMAL(10,2),
     diesel_taka DECIMAL(10,2),
     pamp TEXT,
     commission DECIMAL(10,2),
     month TEXT,
     vehicle_name TEXT,
     vehicle_number TEXT,
     driver_name text,
     driver_phone text,
     created_at TIMESTAMPTZ DEFAULT NOW(),
     updated_at TIMESTAMPTZ DEFAULT NOW()
);

create index idx_vehicle_month on outside_trips(vehicle_number, month);