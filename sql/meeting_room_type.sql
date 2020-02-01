BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'meeting_room_type'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.meeting_room_type_create_pk() RETURNS trigger AS $make_pk$
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

CREATE TABLE core.meeting_room_type (
    id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT meeting_room_type_pk PRIMARY KEY (id)
);

CREATE TRIGGER meeting_room_type_create_pk_tr
     BEFORE INSERT ON core.meeting_room_type
     FOR EACH ROW
     EXECUTE PROCEDURE core.meeting_room_type_create_pk();

CREATE TRIGGER meeting_room_type_set_created_tr
     BEFORE INSERT ON core.meeting_room_type
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER meeting_room_type_set_updated_tr
     BEFORE UPDATE ON core.meeting_room_type
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();


INSERT INTO core.meeting_room_type (name) VALUES 
    ('seminar_hall'),
    ('group_work_hall');

COMMIT;
