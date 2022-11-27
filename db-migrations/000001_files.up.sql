
CREATE TABLE cloud_symbol_server.stores (
  store_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name varchar NOT NULL UNIQUE,
  next_store_upload_index integer NOT NULL
);

CREATE TYPE cloud_symbol_server.store_upload_status AS ENUM (
  'unknown',
  'in_progress',
  'completed',
  'aborted',
  'expired'
);

CREATE TABLE cloud_symbol_server.store_uploads (
  upload_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  store_id integer REFERENCES cloud_symbol_server.stores,
  store_upload_index integer NOT NULL,
  description varchar NOT NULL,
  build varchar NOT NULL,
  timestamp timestamp NOT NULL,
  status cloud_symbol_server.store_upload_status NOT NULL
);

CREATE TYPE cloud_symbol_server.store_upload_file_status AS ENUM (
  'unknown',
  'already_present',
  'pending',
  'uploaded',
  'aborted',
  'expired'
);

CREATE TABLE cloud_symbol_server.store_upload_files (
  file_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  upload_id integer REFERENCES cloud_symbol_server.store_uploads,
  upload_file_index integer NOT NULL,
  file_name varchar NOT NULL,
  hash varchar NOT NULL,
  status cloud_symbol_server.store_upload_file_status NOT NULL
);
