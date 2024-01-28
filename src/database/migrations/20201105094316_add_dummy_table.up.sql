BEGIN;

CREATE TABLE IF NOT EXISTS microservice.dummy
(
    ID SERIAL NOT NULL CONSTRAINT PK_dummy PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    createdBy INT NOT NULL DEFAULT 1,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedBy INT NOT NULL DEFAULT 1,
    isDeleted BOOLEAN NOT NULL DEFAULT FALSE,
    guid UUID NOT NULL DEFAULT public.uuid_generate_v4(),
    name VARCHAR(255) NULL
);

ALTER TABLE microservice.dummy
    ADD CONSTRAINT dummy_guid_unique UNIQUE (guid);

CREATE TRIGGER update_updateAt
    BEFORE
        UPDATE ON microservice.dummy
    FOR EACH ROW
EXECUTE PROCEDURE
    public.update_updatedAt();

COMMIT;