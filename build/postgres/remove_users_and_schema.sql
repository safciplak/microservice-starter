-- REVOKE ALL connections
REVOKE CONNECT ON DATABASE microservice_database FROM microservice_sa;
REVOKE USAGE ON SCHEMA public FROM microservice_sa;
REVOKE USAGE ON SCHEMA microservice FROM microservice_sa;

REVOKE ALL ON SCHEMA microservice FROM microservice_sa;
REVOKE ALL ON ALL TABLES IN SCHEMA microservice FROM microservice_sa;
REVOKE ALL ON ALL SEQUENCES IN SCHEMA microservice FROM microservice_sa;
REVOKE ALL ON ALL FUNCTIONS IN SCHEMA microservice FROM microservice_sa;

REVOKE USAGE ON SCHEMA public FROM microservice_write;
REVOKE USAGE ON SCHEMA microservice FROM microservice_write;
REVOKE ALL ON SCHEMA microservice FROM microservice_write;
REVOKE ALL ON ALL TABLES IN SCHEMA microservice FROM microservice_write;
REVOKE ALL ON ALL SEQUENCES IN SCHEMA microservice FROM microservice_write;
REVOKE ALL ON ALL FUNCTIONS IN SCHEMA microservice FROM microservice_write;

REVOKE USAGE ON SCHEMA public FROM microservice_read;
REVOKE USAGE ON SCHEMA microservice FROM microservice_read;
REVOKE ALL ON SCHEMA microservice FROM microservice_read;
REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA microservice FROM microservice_read;
REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA microservice FROM microservice_read;
REVOKE ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA microservice FROM microservice_read;

DROP OWNED BY microservice_sa;
DROP OWNED BY microservice_write;
DROP OWNED BY microservice_read;

-- Do this once above is run on each database (uat_1 / uat_5)
DROP ROLE microservice_sa;
DROP ROLE microservice_read;
DROP ROLE microservice_write;
