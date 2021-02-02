CREATE USER superuser WITH PASSWORD 'superuser';
CREATE DATABASE fuckthesuits OWNER superuser ENCODING 'UTF8';
GRANT ALL ON DATABASE fuckthesuits TO superuser;
ALTER USER superuser CREATEDB;
