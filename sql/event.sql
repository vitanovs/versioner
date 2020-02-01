BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'event'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.event_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.name), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.event (
    id VARCHAR(64) NOT NULL,
    creator_id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    description VARCHAR(256) NOT NULL,
    is_all_day BOOLEAN DEFAULT FALSE NOT NULL,
    from_date TIMESTAMP WITH TIME ZONE NOT NULL,
    to_date TIMESTAMP WITH TIME ZONE NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT event_pk PRIMARY KEY (id),
    CONSTRAINT event_creator_fk FOREIGN KEY (creator_id)
        REFERENCES core.user (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER event_create_pk_tr
     BEFORE INSERT ON core.event
     FOR EACH ROW
     EXECUTE PROCEDURE core.event_create_pk();

CREATE TRIGGER event_set_created_tr
     BEFORE INSERT ON core.event
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER event_set_updated_tr
     BEFORE UPDATE ON core.event
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
