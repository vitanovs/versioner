BEGIN TRANSACTION;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'add_role_primary_key'::text; $$ IMMUTABLE LANGUAGE sql;

ALTER TABLE core.role
ADD CONSTRAINT role_pk PRIMARY KEY (id);

COMMIT;
