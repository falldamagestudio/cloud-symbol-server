-- Run all SQL operations in the context of the group role, not the specific user
-- This ensures that all DB objects that we create will have the group role as its owner
SET ROLE cloud_symbol_server_admin;

BEGIN;

CREATE TABLE cloud_symbol_server.personal_access_tokens (
  pat_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Identity of owner
  -- This is typically an email address
  owner varchar NOT NULL,

  -- Personal Access Token
  -- This token can be used for authentication when accessing non-token related APIs
  --   and is the authentication method used outside in non-web scenarios
  --   (CLI / other tool integrations)
  token varchar NOT NULL UNIQUE,

  -- Textual description of token
  -- Users fill this in to remind themselves the purpose of a token and/or where it is used
  description varchar NOT NULL,

  -- Creation timestamp, in RFC3339 format
  -- Example: 1985-04-12T23:20:50.52Z
  creation_timestamp timestamp NOT NULL
);

-- Ensure read-write role can access table
GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.personal_access_tokens
TO cloud_symbol_server_readwrite;

CREATE TABLE cloud_symbol_server.stores (
  store_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Textual name of store
  -- This is the primary index used in the API for identifying a store
  name varchar NOT NULL UNIQUE,

  -- Ordinal index for the next upload to this store
  next_store_upload_index integer NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.stores
TO cloud_symbol_server_readwrite;

CREATE TYPE cloud_symbol_server.store_upload_status AS ENUM (
  'in_progress',
  'completed',
  'aborted',
  'expired'
);

GRANT USAGE
ON TYPE cloud_symbol_server.store_upload_status
TO cloud_symbol_server_readwrite;

CREATE TABLE cloud_symbol_server.store_uploads (
  upload_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Reference to the store, which this file belongs to
  store_id integer REFERENCES cloud_symbol_server.stores,

  -- Ordinal index of upload within store
  -- This is the primary index used in the API for identifying an upload within a store
  store_upload_index integer NOT NULL,

  -- Textual description of upload
  description varchar NOT NULL,

  -- Textual description of build
  build varchar NOT NULL,

  -- Upload timestamp, in RFC3339 format
  -- Example: 1985-04-12T23:20:50.52Z
  timestamp timestamp NOT NULL,

  -- The upload status will change over time, based on user actions
  status cloud_symbol_server.store_upload_status NOT NULL,

  -- upload-indices are unique for a store
  UNIQUE (store_id, store_upload_index)
);

-- Ensure read-write role can access table
GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.store_uploads
TO cloud_symbol_server_readwrite;

CREATE TABLE cloud_symbol_server.store_files (
  file_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Reference to the store, which this file belongs to
  store_id integer REFERENCES cloud_symbol_server.stores,

  -- Textual name of file
  -- Must be all lowercase
  -- This is the primary index used in the API for identifying a file within a store
  file_name varchar NOT NULL,

  -- File names are unique within a store
  UNIQUE (store_id, file_name)
);

-- Ensure read-write role can access table
GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.store_files
TO cloud_symbol_server_readwrite;

CREATE TYPE cloud_symbol_server.store_file_hash_status AS ENUM (
  'pending',
  'present'
);

GRANT USAGE
ON TYPE cloud_symbol_server.store_file_hash_status
TO cloud_symbol_server_readwrite;

CREATE TABLE cloud_symbol_server.store_file_hashes (
  hash_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Reference to the file, which this hash belongs to
  file_id integer REFERENCES cloud_symbol_server.store_files,

  -- Hash string for this blob
  -- Uppercase vs lowercase varies, depending on type of file
  --   (exe hash vs pdb has has different rules)
  -- This is the primary index used in the API for identifying a hash of a file
  hash varchar NOT NULL,

  -- The hash status will change over time, based on user actions
  status cloud_symbol_server.store_file_hash_status NOT NULL,

  -- Hashes are unique for a file
  UNIQUE (file_id, hash)
);

-- Ensure read-write role can access table
GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.store_file_hashes
TO cloud_symbol_server_readwrite;

CREATE TYPE cloud_symbol_server.store_upload_file_status AS ENUM (
  'already_present',
  'pending',
  'completed',
  'aborted',
  'expired'
);

GRANT USAGE
ON TYPE cloud_symbol_server.store_upload_file_status
TO cloud_symbol_server_readwrite;

CREATE TABLE cloud_symbol_server.store_upload_files (
  file_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  -- Reference to the upload, which this upload-file belongs to
  upload_id integer REFERENCES cloud_symbol_server.store_uploads,

  -- Reference to the hash, which this upload-file resulted in an upload of
  --   or null, if the upload has been expired
  hash_id integer REFERENCES cloud_symbol_server.store_file_hashes,

  -- Ordinal index of upload-file within upload
  -- This is the primary index used in the API for identifying an upload-hash within an upload
  upload_file_index integer NOT NULL,

  -- The upload-file status will change over time, based on user actions
  status cloud_symbol_server.store_upload_file_status NOT NULL,

  -- Textual name of file
  -- Duplicated from store_files
  --   since this will persist after the upload has been expired
  --   and the corresponding store_file might have been removed
  file_name varchar NOT NULL,

  -- Hash string for this blob
  -- Duplicated from store_file_hashes
  --   since this will persist after the upload has been expired
  --   and the corresponding store_file_hash might have been removed
  -- This is not named 'hash' as in the original table,
  --   because SQLBoiler will generate a Hash() method on the corresponding Golang type,
  --   and that method will collide with the type's Hash property
  file_hash varchar NOT NULL,

  -- upload-file-indices are unique for an upload
  UNIQUE (upload_id, upload_file_index)
);

-- Ensure read-write role can access table
GRANT SELECT, INSERT, UPDATE, DELETE
ON cloud_symbol_server.store_upload_files
TO cloud_symbol_server_readwrite;

COMMIT;