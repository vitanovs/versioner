BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'user'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.user_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.email), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.user (
    id VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    email VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (id),
    CONSTRAINT user_unique UNIQUE (email)
);

CREATE TRIGGER user_create_pk_tr
     BEFORE INSERT ON core.user
     FOR EACH ROW
     EXECUTE PROCEDURE core.user_create_pk();

CREATE TRIGGER user_set_created_tr
     BEFORE INSERT ON core.user
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER user_set_updated_tr
     BEFORE UPDATE ON core.user
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
