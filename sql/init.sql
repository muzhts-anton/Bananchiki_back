DROP DATABASE IF EXISTS tpfinal;
DROP ROLE IF EXISTS bananchiki;

CREATE ROLE bananchiki WITH PASSWORD '1234';
ALTER ROLE bananchiki WITH LOGIN superuser;

CREATE DATABASE tpfinal
WITH OWNER = bananchiki
ENCODING = 'UTF8'
TABLESPACE = pg_default
CONNECTION LIMIT = -1;
