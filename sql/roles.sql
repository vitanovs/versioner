BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'roles'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.role_create_pk() RETURNS trigger AS $make_pk$
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

CREATE TABLE core.role (
    id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TRIGGER role_create_pk_tr
     BEFORE INSERT ON core.role
     FOR EACH ROW
     EXECUTE PROCEDURE core.role_create_pk();

CREATE TRIGGER role_set_created_tr
     BEFORE INSERT ON core.role
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER role_set_updated_tr
     BEFORE UPDATE ON core.role
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

INSERT INTO core.role (name) VALUES 
    ('super_admin'),
    ('admin'),
    ('user');

COMMIT;
