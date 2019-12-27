BEGIN TRANSACTION;

CREATE SCHEMA core;
CREATE SCHEMA utility;

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;
CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;

CREATE OR REPLACE FUNCTION utility.set_created() RETURNS trigger AS $set_created$
     BEGIN
         NEW.created := now();
         NEW.updated := now();
         RETURN NEW;
     END;
 $set_created$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION utility.set_updated() RETURNS trigger AS $set_updated$
     BEGIN
         NEW.updated := now();
         RETURN NEW;
     END;
 $set_updated$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION core.schema_version() RETURNS text AS $$ SELECT 'base'::text; $$ IMMUTABLE LANGUAGE sql;

COMMIT;
