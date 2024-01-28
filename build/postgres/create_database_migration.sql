-- Create the database if it doesn't exists
-- CREATE DATABASE microservice_database;

-- CREATE ROLES / USERS
CREATE ROLE microservice_read
    WITH
    NOSUPERUSER
    NOCREATEDB
    NOCREATEROLE
    INHERIT
    LOGIN
    ENCRYPTED PASSWORD ''; -- use env variable

CREATE ROLE microservice_write
    WITH
    NOSUPERUSER
    NOCREATEDB
    NOCREATEROLE
    INHERIT
    LOGIN
    ENCRYPTED PASSWORD ''; -- use env variable

CREATE ROLE microservice_sa
    WITH
    NOSUPERUSER
    CREATEDB
    CREATEROLE
    INHERIT
    LOGIN
    ENCRYPTED PASSWORD ''; -- use env variable



-- GRANT NEW ROLES TO SA user
GRANT microservice_sa TO sa;
GRANT microservice_write TO microservice_sa;
GRANT microservice_read TO microservice_sa;

GRANT ALL PRIVILEGES ON DATABASE microservice_database TO sa;
GRANT CONNECT ON DATABASE microservice_database TO microservice_sa;

-- CREATE SCHEMA
CREATE SCHEMA IF NOT EXISTS microservice
    AUTHORIZATION microservice_sa;

-- GRANT ADMIN ACCESS TO SA USER
GRANT CONNECT ON DATABASE microservice_database TO microservice_sa;
GRANT ALL ON SCHEMA microservice TO microservice_sa;
GRANT ALL ON ALL TABLES IN SCHEMA microservice TO microservice_sa;
GRANT ALL ON ALL SEQUENCES IN SCHEMA microservice TO microservice_sa;
GRANT ALL ON ALL FUNCTIONS IN SCHEMA microservice TO microservice_sa;
GRANT ALL ON ALL FUNCTIONS IN SCHEMA public TO microservice_sa;

ALTER DEFAULT PRIVILEGES IN SCHEMA microservice
    GRANT ALL ON TABLES TO microservice_sa;

-- GRANT READ ACCESS TO READ USER
GRANT CONNECT ON DATABASE microservice_database TO microservice_read;
GRANT USAGE ON SCHEMA microservice TO microservice_read;
GRANT SELECT ON ALL TABLES IN SCHEMA microservice TO microservice_read;
GRANT SELECT ON ALL SEQUENCES IN SCHEMA microservice TO microservice_read;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA microservice TO microservice_read;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO microservice_read;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT SELECT ON TABLES TO microservice_read;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT SELECT ON SEQUENCES TO microservice_read;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT EXECUTE ON FUNCTIONS TO microservice_read;

-- GRANT WRITE ACCESS TO WRITE USER
GRANT CONNECT ON DATABASE microservice_database TO microservice_write;
GRANT USAGE ON SCHEMA microservice TO microservice_write;
GRANT SELECT,INSERT,UPDATE ON ALL TABLES IN SCHEMA microservice TO microservice_write;
GRANT SELECT,USAGE ON ALL SEQUENCES IN SCHEMA microservice TO microservice_write;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA microservice TO microservice_write;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO microservice_write;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT SELECT,INSERT,UPDATE ON TABLES TO microservice_write;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT SELECT,USAGE ON SEQUENCES TO microservice_write;

ALTER DEFAULT PRIVILEGES
    FOR ROLE    microservice_sa
    IN SCHEMA   microservice
    GRANT EXECUTE ON FUNCTIONS TO microservice_write;

-- The UUID plugin won't work without SUPERUSER rights and they cannot be given during creation so alter it here --
ALTER ROLE microservice_sa
    WITH
    SUPERUSER
    CREATEDB
    CREATEROLE
    INHERIT
    LOGIN;
