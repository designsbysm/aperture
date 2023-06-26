SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: authentication; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA authentication;


--
-- Name: logger; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA logger;


--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: user_role; Type: TYPE; Schema: authentication; Owner: -
--

CREATE TYPE authentication.user_role AS ENUM (
    'admin',
    'operations',
    'provider',
    'user',
    'disabled'
);


--
-- Name: level; Type: TYPE; Schema: logger; Owner: -
--

CREATE TYPE logger.level AS ENUM (
    'info',
    'warn',
    'error'
);


--
-- Name: services; Type: TYPE; Schema: logger; Owner: -
--

CREATE TYPE logger.services AS ENUM (
    'authentication',
    'rest-api'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: access_token_refreshes; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.access_token_refreshes (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    refresh_token uuid NOT NULL
);


--
-- Name: access_token_refreshes_id_seq; Type: SEQUENCE; Schema: authentication; Owner: -
--

ALTER TABLE authentication.access_token_refreshes ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME authentication.access_token_refreshes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: refresh_tokens; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.refresh_tokens (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    user_id uuid NOT NULL
);


--
-- Name: roles; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.roles (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    user_id uuid NOT NULL,
    role authentication.user_role DEFAULT 'user'::authentication.user_role NOT NULL
);


--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: authentication; Owner: -
--

ALTER TABLE authentication.roles ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME authentication.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    mobile text NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);


--
-- Name: events; Type: TABLE; Schema: logger; Owner: -
--

CREATE TABLE logger.events (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    service logger.services NOT NULL,
    level logger.level NOT NULL,
    message text NOT NULL
);


--
-- Name: events_id_seq; Type: SEQUENCE; Schema: logger; Owner: -
--

ALTER TABLE logger.events ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME logger.events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: access_token_refreshes access_token_refreshes_pkey; Type: CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.access_token_refreshes
    ADD CONSTRAINT access_token_refreshes_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: logger; Owner: -
--

ALTER TABLE ONLY logger.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: idx_access_token_refreshes_deleted_at; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_access_token_refreshes_deleted_at ON authentication.access_token_refreshes USING btree (deleted_at);


--
-- Name: idx_access_token_refreshes_refresh_token_fkey; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_access_token_refreshes_refresh_token_fkey ON authentication.access_token_refreshes USING btree (refresh_token);


--
-- Name: idx_refresh_tokens_deleted_at; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_refresh_tokens_deleted_at ON authentication.refresh_tokens USING btree (deleted_at);


--
-- Name: idx_refresh_tokens_user_id_fkey; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_refresh_tokens_user_id_fkey ON authentication.refresh_tokens USING btree (user_id);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_roles_deleted_at ON authentication.roles USING btree (deleted_at);


--
-- Name: idx_roles_user_id_fkey; Type: INDEX; Schema: authentication; Owner: -
--

CREATE UNIQUE INDEX idx_roles_user_id_fkey ON authentication.roles USING btree (user_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: authentication; Owner: -
--

CREATE INDEX idx_users_deleted_at ON authentication.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: authentication; Owner: -
--

CREATE UNIQUE INDEX idx_users_email ON authentication.users USING btree (email);


--
-- Name: idx_users_mobile; Type: INDEX; Schema: authentication; Owner: -
--

CREATE UNIQUE INDEX idx_users_mobile ON authentication.users USING btree (mobile);


--
-- Name: idx_logs_deleted_at; Type: INDEX; Schema: logger; Owner: -
--

CREATE INDEX idx_logs_deleted_at ON logger.events USING btree (deleted_at);


--
-- Name: idx_logs_level; Type: INDEX; Schema: logger; Owner: -
--

CREATE INDEX idx_logs_level ON logger.events USING btree (level);


--
-- Name: idx_logs_message; Type: INDEX; Schema: logger; Owner: -
--

CREATE INDEX idx_logs_message ON logger.events USING btree (message);


--
-- Name: idx_logs_service; Type: INDEX; Schema: logger; Owner: -
--

CREATE INDEX idx_logs_service ON logger.events USING btree (service);


--
-- Name: access_token_refreshes access_token_refreshes_refresh_token_fkey; Type: FK CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.access_token_refreshes
    ADD CONSTRAINT access_token_refreshes_refresh_token_fkey FOREIGN KEY (refresh_token) REFERENCES authentication.refresh_tokens(id);


--
-- Name: refresh_tokens refresh_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.refresh_tokens
    ADD CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES authentication.users(id);


--
-- Name: roles roles_user_id_fkey; Type: FK CONSTRAINT; Schema: authentication; Owner: -
--

ALTER TABLE ONLY authentication.roles
    ADD CONSTRAINT roles_user_id_fkey FOREIGN KEY (user_id) REFERENCES authentication.users(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20221124071125'),
    ('20221124071126'),
    ('20221127000046');
