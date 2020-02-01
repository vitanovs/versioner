BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'add_user_to_role_relation'::text; $$ IMMUTABLE LANGUAGE sql;

ALTER TABLE core.user
ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT FALSE;

CREATE OR REPLACE FUNCTION core.users_roles_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.user_id, NEW.role_id), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.users_roles (
    id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    role_id VARCHAR(64) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT users_roles_pk PRIMARY KEY (id),
    CONSTRAINT users_roles_to_user_fk FOREIGN KEY (user_id)
        REFERENCES core.user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT users_roles_to_role_fk FOREIGN KEY (role_id)
        REFERENCES core.role (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER users_roles_create_pk_tr
     BEFORE INSERT ON core.users_roles
     FOR EACH ROW
     EXECUTE PROCEDURE core.users_roles_create_pk();

CREATE TRIGGER users_roles_set_created_tr
     BEFORE INSERT ON core.users_roles
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER users_roles_set_updated_tr
     BEFORE UPDATE ON core.users_roles
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
