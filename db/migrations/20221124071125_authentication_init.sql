-- migrate:up
CREATE SCHEMA authentication;

CREATE TABLE authentication.users (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  first_name text NOT NULL,
  last_name text NOT NULL,
  mobile text NOT NULL,
  email text NOT NULL,
  password text NOT NULL
);
CREATE INDEX idx_users_deleted_at ON authentication.users USING btree (deleted_at);
CREATE UNIQUE INDEX idx_users_mobile ON authentication.users USING btree (mobile);
CREATE UNIQUE INDEX idx_users_email ON authentication.users USING btree (email);

CREATE TYPE authentication.user_role AS ENUM (
  'admin',
  'operations',
  'provider',
  'user',
  'disabled'
);

CREATE TABLE authentication.roles (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  user_id uuid NOT NULL,
  role authentication.user_role DEFAULT 'user'::authentication.user_role NOT NULL,
  FOREIGN KEY (user_id) REFERENCES authentication.users(id)
);
CREATE INDEX idx_roles_deleted_at ON authentication.roles USING btree (deleted_at);
CREATE UNIQUE INDEX idx_roles_user_id_fkey ON authentication.roles USING btree (user_id);

CREATE TABLE authentication.refresh_tokens (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  user_id uuid NOT NULL,
  FOREIGN KEY (user_id) REFERENCES authentication.users(id)
);
CREATE INDEX idx_refresh_tokens_deleted_at ON authentication.refresh_tokens USING btree (deleted_at);
CREATE INDEX idx_refresh_tokens_user_id_fkey ON authentication.refresh_tokens USING btree (user_id);

CREATE TABLE authentication.access_token_refreshes (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  refresh_token uuid NOT NULL,
  FOREIGN KEY (refresh_token) REFERENCES authentication.refresh_tokens(id)
);
CREATE INDEX idx_access_token_refreshes_deleted_at ON authentication.access_token_refreshes USING btree (deleted_at);
CREATE INDEX idx_access_token_refreshes_refresh_token_fkey ON authentication.access_token_refreshes USING btree (refresh_token);

-- migrate:down
DROP INDEX authentication.idx_access_token_refreshes_deleted_at;
DROP INDEX authentication.idx_access_token_refreshes_refresh_token_fkey;
DROP TABLE authentication.access_token_refreshes;

DROP INDEX authentication.idx_refresh_tokens_deleted_at;
DROP INDEX authentication.idx_refresh_tokens_user_id_fkey;
DROP TABLE authentication.refresh_tokens;

DROP INDEX authentication.idx_roles_deleted_at;
DROP INDEX authentication.idx_roles_user_id_fkey;
DROP TABLE authentication.roles;

DROP TYPE authentication.user_role;

DROP INDEX authentication.idx_users_deleted_at;
DROP INDEX authentication.idx_users_email;
DROP TABLE authentication.users;
DROP SCHEMA authentication;
