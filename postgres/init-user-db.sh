#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER hahnair;
    CREATE DATABASE cards;
    GRANT ALL PRIVILEGES ON DATABASE cards TO hahnair;
EOSQL