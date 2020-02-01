BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'events_and_meeting_rooms'::text; $$ IMMUTABLE LANGUAGE sql;

CREATE OR REPLACE FUNCTION core.events_and_meeting_rooms_create_pk() RETURNS trigger AS $make_pk$
 DECLARE
     pk char(64);
 BEGIN
     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.event_id, NEW.meeting_room_id), 'sha256'), 'hex'));
     IF NEW.id IS DISTINCT FROM pk THEN
         NEW.id := pk;
     END IF;
     RETURN NEW;
 END;
 $make_pk$ LANGUAGE plpgsql;

CREATE TABLE core.events_and_meeting_rooms (
    id VARCHAR(64) NOT NULL,
    event_id VARCHAR(64) NOT NULL,
    meeting_room_id VARCHAR(64) NOT NULL,
    CONSTRAINT events_and_meeting_rooms_pk PRIMARY KEY (id),
    CONSTRAINT events_and_meeting_rooms_to_event_fk FOREIGN KEY (event_id)
        REFERENCES core.event (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT events_and_meeting_rooms_to_meeting_room_fk FOREIGN KEY (meeting_room_id)
        REFERENCES core.meeting_room (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TRIGGER events_and_meeting_rooms_create_pk_tr
     BEFORE INSERT ON core.events_and_meeting_rooms
     FOR EACH ROW
     EXECUTE PROCEDURE core.events_and_meeting_rooms_create_pk();

CREATE TRIGGER events_and_meeting_rooms_set_created_tr
     BEFORE INSERT ON core.events_and_meeting_rooms
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_created();

CREATE TRIGGER events_and_meeting_rooms_set_updated_tr
     BEFORE UPDATE ON core.events_and_meeting_rooms
     FOR EACH ROW
     EXECUTE PROCEDURE utility.set_updated();

COMMIT;
