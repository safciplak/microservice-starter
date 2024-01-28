BEGIN;

DROP TRIGGER IF EXISTS update_updateAt ON microservice.schema_migrations;

ALTER TABLE IF EXISTS microservice.schema_migrations
DROP COLUMN IF EXISTS updatedAt;

COMMIT;