-- migrate:up
CREATE SCHEMA logger;

CREATE TABLE logger.logs (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    service text NOT NULL,
    level text NOT NULL,
    message text NOT NULL
);
CREATE INDEX idx_logs_deleted_at ON logger.logs USING btree (deleted_at);
CREATE INDEX idx_logs_level ON logger.logs USING btree (level);
CREATE INDEX idx_logs_message ON logger.logs USING btree (message);
CREATE INDEX idx_logs_service ON logger.logs USING btree (service);

-- CREATE TYPE authentication.user_role AS ENUM (
--     'admin',
--     'operations',
--     'provider',
--     'user',
--     'disabled'
-- );

-- migrate:down
DROP INDEX logger.idx_logs_deleted_at;
DROP INDEX logger.idx_logs_level;
DROP INDEX logger.idx_logs_message;
DROP INDEX logger.idx_logs_service;
DROP TABLE logger.logs;
DROP SCHEMA logger;
