-- migrate:up
CREATE SCHEMA authentication;

CREATE TABLE authentication.users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);
CREATE INDEX idx_users_deleted_at ON authentication.users USING btree (deleted_at);
CREATE UNIQUE INDEX idx_users_email ON authentication.users USING btree (email);

CREATE TYPE authentication.user_role AS ENUM (
    'user',
    'provider',
    'operations',
    'admin'
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

CREATE TABLE authentication.sessions (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid NOT NULL,
    token text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES authentication.users(id)
);
CREATE INDEX idx_sessions_deleted_at ON authentication.sessions USING btree (deleted_at);
CREATE UNIQUE INDEX idx_sessions_user_id_fkey ON authentication.sessions USING btree (user_id);
CREATE UNIQUE INDEX idx_sessions_token ON authentication.sessions USING btree (token);

-- migrate:down
DROP INDEX authentication.idx_sessions_deleted_at;
DROP INDEX authentication.idx_sessions_user_id_fkey;
DROP INDEX authentication.idx_sessions_token;
DROP TABLE authentication.sessions;

DROP INDEX authentication.idx_roles_deleted_at;
DROP INDEX authentication.idx_roles_user_id_fkey;
DROP TABLE authentication.roles;

DROP TYPE authentication.user_role;

DROP INDEX authentication.idx_users_deleted_at;
DROP INDEX authentication.idx_users_email;
DROP TABLE authentication.users;
DROP SCHEMA authentication;
