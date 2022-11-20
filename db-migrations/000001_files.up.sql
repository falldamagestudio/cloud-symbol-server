CREATE TABLE cloud_symbol_server.files (
  id serial PRIMARY KEY,
  file_name varchar NOT NULL,
  hash varchar NOT NULL,
  status varchar NOT NULL
);
