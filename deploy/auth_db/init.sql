CREATE USER auth_user_owner WITH LOGIN PASSWORD 'owner_pass';
CREATE USER auth_user_app WITH LOGIN PASSWORD 'app_pass';

ALTER database auth OWNER TO auth_user_owner;

CREATE SCHEMA app;
ALTER schema app OWNER TO auth_user_owner;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA app TO auth_user_owner;
GRANT USAGE ON SCHEMA app TO auth_user_app;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";