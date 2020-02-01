BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'meeting_room'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.meeting_room_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.name, NEW.location), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.meeting_room (
    id VARCHAR(64) NOT NULL,
    type_id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    description VARCHAR(256) NOT NULL,
    location VARCHAR(64) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT meeting_room_pk PRIMARY KEY (id),
    CONSTRAINT meeting_room_type_fk FOREIGN KEY (type_id)
        REFERENCES core.meeting_room_type (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER meeting_room_create_pk_tr
     BEFORE INSERT ON core.meeting_room
     FOR EACH ROW
     EXECUTE PROCEDURE core.meeting_room_create_pk();

CREATE TRIGGER meeting_room_set_created_tr
     BEFORE INSERT ON core.meeting_room
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER meeting_room_set_updated_tr
     BEFORE UPDATE ON core.meeting_room
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
