   DROP INDEX IF EXISTS ux_users_google_id;

   alter table users 
    drop column if exists google_id,
    ADD COLUMN IF NOT EXISTS password TEXT NOT NULL,
    ADD COLUMN IF NOT EXISTS role TEXT NOT NULL DEFAULT 'manager';

alter table users 
add constraint check_role_valid check (role in ('admin','manager'));