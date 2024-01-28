BEGIN;

ALTER TABLE microservice.schema_migrations
ADD COLUMN updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

CREATE TRIGGER update_updateAt
BEFORE
    UPDATE ON microservice.schema_migrations
    FOR EACH ROW
    EXECUTE PROCEDURE
    public.update_updatedAt();

COMMIT;


COMMIT;