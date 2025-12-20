CREATE TABLE daily_side_cash (
         id BIGSERIAL PRIMARY KEY,
         date DATE NOT NULL,
         suzuki DECIMAL(15,4) DEFAULT 0.0000,
         yamaha DECIMAL(15,4) DEFAULT 0.0000,
        honda DECIMAL(15,4) DEFAULT 0.0000,
         hatim_rupgonj DECIMAL(15,4) DEFAULT 0.0000,
        rakib_benapole DECIMAL(15,4) DEFAULT 0.0000,
        mofiz_benapole DECIMAL(15,4) DEFAULT 0.0000,
        aziz DECIMAL(15,4) DEFAULT 0.0000,
        shongram DECIMAL(15,4) DEFAULT 0.0000,
        created_at TIMESTAMPTZ NOT NULL DEFAULT TIMEZONE('UTC', NOW()),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT TIMEZONE('UTC', NOW()),
        deleted_at TIMESTAMPTZ
    );


   create index idx_daily_side_cash_date on daily_side_cash(date);
   CREATE index idx_daily_side_cash_deleted_at ON daily_side_cash(deleted_at);