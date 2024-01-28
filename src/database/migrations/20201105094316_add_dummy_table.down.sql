BEGIN;

DROP TRIGGER IF EXISTS update_updateAt ON microservice.dummy;

ALTER TABLE microservice.dummy
    DROP CONSTRAINT dummy_guid_unique;

DROP TABLE IF EXISTS microservice.dummy CASCADE;

COMMIT;