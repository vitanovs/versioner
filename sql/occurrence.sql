BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'occurrence'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.occurrence_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.event_id), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.occurrence (
    id VARCHAR(64) NOT NULL,
    event_id VARCHAR(64) NOT NULL,
    occurrence_type_id VARCHAR(256) NOT NULL,
    occurr_monday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_tuesday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_wednesday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_thursday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_friday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_saturday BOOLEAN NOT NULL DEFAULT FALSE,
    occurr_sunday BOOLEAN NOT NULL DEFAULT FALSE,
    does_not_end BOOLEAN NOT NULL DEFAULT FALSE,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT occurrence_pk PRIMARY KEY (id),
    CONSTRAINT occurrence_event_fk FOREIGN KEY (event_id)
        REFERENCES core.event (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT occurrence_type_fk FOREIGN KEY (occurrence_type_id)
        REFERENCES core.occurrence_type (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER occurrence_create_pk_tr
     BEFORE INSERT ON core.occurrence
     FOR EACH ROW
     EXECUTE PROCEDURE core.occurrence_create_pk();

CREATE TRIGGER occurrence_set_created_tr
     BEFORE INSERT ON core.occurrence
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER occurrence_set_updated_tr
     BEFORE UPDATE ON core.occurrence
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
