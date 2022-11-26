
CREATE TABLE cloud_symbol_server.stores (
  store_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name varchar NOT NULL UNIQUE
);

CREATE TYPE cloud_symbol_server.upload_status AS ENUM (
  'unknown',
  'in_progress',
  'completed',
  'aborted',
  'expired'
);

CREATE TABLE cloud_symbol_server.uploads (
  upload_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_id integer REFERENCES cloud_symbol_server.stores,
  description varchar NOT NULL,
  build varchar NOT NULL,
  timestamp timestamp NOT NULL,
  status cloud_symbol_server.upload_status NOT NULL
);

CREATE TYPE cloud_symbol_server.file_status AS ENUM (
  'unknown',
  'already_present',
  'pending',
  'uploaded',
  'aborted',
  'expired'
);

CREATE TABLE cloud_symbol_server.files (
  file_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  upload_id integer REFERENCES cloud_symbol_server.uploads,
  file_name varchar NOT NULL,
  hash varchar NOT NULL,
  upload_file_index integer NOT NULL,
  status cloud_symbol_server.file_status NOT NULL
);
