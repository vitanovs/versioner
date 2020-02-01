BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'occurrence_type'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.occurrence_type_create_pk() RETURNS trigger AS $make_pk$
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

CREATE TABLE core.occurrence_type (
    id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT occurrence_type_pk PRIMARY KEY (id)
);

CREATE TRIGGER occurrence_type_create_pk_tr
     BEFORE INSERT ON core.occurrence_type
     FOR EACH ROW
     EXECUTE PROCEDURE core.occurrence_type_create_pk();

CREATE TRIGGER occurrence_type_set_created_tr
     BEFORE INSERT ON core.occurrence_type
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER occurrence_type_set_updated_tr
     BEFORE UPDATE ON core.occurrence_type
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

INSERT INTO core.occurrence_type (name) VALUES 
    ('weekly'),
    ('monthly');

COMMIT;
