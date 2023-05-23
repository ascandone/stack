#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE auth;
	CREATE DATABASE ledger;
	GRANT ALL PRIVILEGES ON DATABASE auth TO $POSTGRES_USER;
	GRANT ALL PRIVILEGES ON DATABASE ledger TO $POSTGRES_USER;
EOSQL