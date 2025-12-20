create table own_vehicle_trips (
       id  bigserial PRIMARY KEY,
       vehicle_no varchar(255) not null,
       driver_name varchar(255) not null,
       load_point varchar(255) not null,
       unload_point varchar(255) not null,
       pump   varchar(255) not null,
       rent  decimal(10,4),
       advance decimal(10,4),
       trip_cost decimal(10,4),
       diesel     decimal(10,4),
       diesel_price decimal(10,4),
       diesel_taka decimal(10,4),
       extra_cost decimal(10,4),
       commission decimal(10,4),
       profit     decimal(10,4),
       created_at timestamptz not null default timezone('UTC', now()),
       updated_at timestamptz not null default timezone('UTC', now()),
       deleted_at timestamptz

);

create index idx_own_vehicle_trips_vehicle_no on own_vehicle_trips(vehicle_no);
create index idx_own_vehicle_trips_driver_name on own_vehicle_trips(driver_name);
create index idx_own_vehicle_trips_deleted_at on own_vehicle_trips(deleted_at);