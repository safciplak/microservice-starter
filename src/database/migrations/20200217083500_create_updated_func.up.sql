BEGIN;

DO
$$
    BEGIN
        IF NOT EXISTS(
            SELECT 	*
            FROM    pg_proc
            WHERE   proname = 'update_updatedat'
            AND     pronamespace::regnamespace::text = 'public'
            )
        THEN
            CREATE OR REPLACE FUNCTION public.update_updatedAt()
                RETURNS TRIGGER AS $updated_trigger$
            BEGIN
                NEW.updatedAt = now();
                RETURN NEW;
            END;
            $updated_trigger$ LANGUAGE 'plpgsql';
        END IF ;
    END
$$;

COMMIT;