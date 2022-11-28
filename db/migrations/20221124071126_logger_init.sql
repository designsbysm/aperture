-- migrate:up
CREATE SCHEMA logger;

CREATE TYPE logger.services AS ENUM (
  'authentication'
);

CREATE TYPE logger.level AS ENUM (
  'info',
  'warn',
  'error'
);

CREATE TABLE logger.events (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  service logger.services NOT NULL,
  level logger.level NOT NULL,
  message text NOT NULL
);
CREATE INDEX idx_logs_deleted_at ON logger.events USING btree (deleted_at);
CREATE INDEX idx_logs_level ON logger.events USING btree (level);
CREATE INDEX idx_logs_message ON logger.events USING btree (message);
CREATE INDEX idx_logs_service ON logger.events USING btree (service);


-- migrate:down
DROP INDEX logger.idx_logs_deleted_at;
DROP INDEX logger.idx_logs_level;
DROP INDEX logger.idx_logs_message;
DROP INDEX logger.idx_logs_service;
DROP TABLE logger.events;

DROP TYPE logger.services;
DROP TYPE logger.level;

DROP SCHEMA logger;
