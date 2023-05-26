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
-- Name: northwind; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA northwind;


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
    'authentication'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: access_token_refreshes; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.access_token_refreshes (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
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
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid NOT NULL
);


--
-- Name: roles; Type: TABLE; Schema: authentication; Owner: -
--

CREATE TABLE authentication.roles (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
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
    updated_at timestamp with time zone,
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
-- Name: categories; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.categories (
    category_id smallint NOT NULL,
    category_name character varying(15) NOT NULL,
    description text
);


--
-- Name: customers; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.customers (
    customer_id bpchar NOT NULL,
    company_name character varying(40) NOT NULL,
    contact_name character varying(30),
    contact_title character varying(30),
    address character varying(60),
    city character varying(15),
    region character varying(15),
    postal_code character varying(10),
    country character varying(15),
    phone character varying(24),
    fax character varying(24)
);


--
-- Name: employee_territories; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.employee_territories (
    employee_id smallint NOT NULL,
    territory_id character varying(20) NOT NULL
);


--
-- Name: employees; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.employees (
    employee_id smallint NOT NULL,
    last_name character varying(20) NOT NULL,
    first_name character varying(10) NOT NULL,
    title character varying(30),
    title_of_courtesy character varying(25),
    birth_date date,
    hire_date date,
    address character varying(60),
    city character varying(15),
    region character varying(15),
    postal_code character varying(10),
    country character varying(15),
    home_phone character varying(24),
    extension character varying(4),
    notes text,
    reports_to smallint,
    photo_path character varying(255)
);


--
-- Name: order_details; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.order_details (
    order_id smallint NOT NULL,
    product_id smallint NOT NULL,
    unit_price real NOT NULL,
    quantity smallint NOT NULL,
    discount real NOT NULL
);


--
-- Name: orders; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.orders (
    order_id smallint NOT NULL,
    customer_id bpchar,
    employee_id smallint,
    order_date date,
    required_date date,
    shipped_date date,
    ship_via smallint,
    freight real,
    ship_name character varying(40),
    ship_address character varying(60),
    ship_city character varying(15),
    ship_region character varying(15),
    ship_postal_code character varying(10),
    ship_country character varying(15)
);


--
-- Name: products; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.products (
    product_id smallint NOT NULL,
    product_name character varying(40) NOT NULL,
    supplier_id smallint,
    category_id smallint,
    quantity_per_unit character varying(20),
    unit_price real,
    units_in_stock smallint,
    units_on_order smallint,
    reorder_level smallint,
    discontinued integer NOT NULL
);


--
-- Name: region; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.region (
    region_id smallint NOT NULL,
    region_description bpchar NOT NULL
);


--
-- Name: shippers; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.shippers (
    shipper_id smallint NOT NULL,
    company_name character varying(40) NOT NULL,
    phone character varying(24)
);


--
-- Name: suppliers; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.suppliers (
    supplier_id smallint NOT NULL,
    company_name character varying(40) NOT NULL,
    contact_name character varying(30),
    contact_title character varying(30),
    address character varying(60),
    city character varying(15),
    region character varying(15),
    postal_code character varying(10),
    country character varying(15),
    phone character varying(24),
    fax character varying(24),
    homepage text
);


--
-- Name: territories; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.territories (
    territory_id character varying(20) NOT NULL,
    territory_description bpchar NOT NULL,
    region_id smallint NOT NULL
);


--
-- Name: us_states; Type: TABLE; Schema: northwind; Owner: -
--

CREATE TABLE northwind.us_states (
    state_id smallint NOT NULL,
    state_name character varying(100),
    state_abbr character varying(2),
    state_region character varying(50)
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
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
-- Name: categories pk_categories; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.categories
    ADD CONSTRAINT pk_categories PRIMARY KEY (category_id);


--
-- Name: customers pk_customers; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.customers
    ADD CONSTRAINT pk_customers PRIMARY KEY (customer_id);


--
-- Name: employee_territories pk_employee_territories; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.employee_territories
    ADD CONSTRAINT pk_employee_territories PRIMARY KEY (employee_id, territory_id);


--
-- Name: employees pk_employees; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.employees
    ADD CONSTRAINT pk_employees PRIMARY KEY (employee_id);


--
-- Name: order_details pk_order_details; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.order_details
    ADD CONSTRAINT pk_order_details PRIMARY KEY (order_id, product_id);


--
-- Name: orders pk_orders; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.orders
    ADD CONSTRAINT pk_orders PRIMARY KEY (order_id);


--
-- Name: products pk_products; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.products
    ADD CONSTRAINT pk_products PRIMARY KEY (product_id);


--
-- Name: region pk_region; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.region
    ADD CONSTRAINT pk_region PRIMARY KEY (region_id);


--
-- Name: shippers pk_shippers; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.shippers
    ADD CONSTRAINT pk_shippers PRIMARY KEY (shipper_id);


--
-- Name: suppliers pk_suppliers; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.suppliers
    ADD CONSTRAINT pk_suppliers PRIMARY KEY (supplier_id);


--
-- Name: territories pk_territories; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.territories
    ADD CONSTRAINT pk_territories PRIMARY KEY (territory_id);


--
-- Name: us_states pk_usstates; Type: CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.us_states
    ADD CONSTRAINT pk_usstates PRIMARY KEY (state_id);


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
-- Name: employee_territories fk_employee_territories_employees; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.employee_territories
    ADD CONSTRAINT fk_employee_territories_employees FOREIGN KEY (employee_id) REFERENCES northwind.employees(employee_id);


--
-- Name: employee_territories fk_employee_territories_territories; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.employee_territories
    ADD CONSTRAINT fk_employee_territories_territories FOREIGN KEY (territory_id) REFERENCES northwind.territories(territory_id);


--
-- Name: employees fk_employees_employees; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.employees
    ADD CONSTRAINT fk_employees_employees FOREIGN KEY (reports_to) REFERENCES northwind.employees(employee_id);


--
-- Name: order_details fk_order_details_orders; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.order_details
    ADD CONSTRAINT fk_order_details_orders FOREIGN KEY (order_id) REFERENCES northwind.orders(order_id);


--
-- Name: order_details fk_order_details_products; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.order_details
    ADD CONSTRAINT fk_order_details_products FOREIGN KEY (product_id) REFERENCES northwind.products(product_id);


--
-- Name: orders fk_orders_customers; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.orders
    ADD CONSTRAINT fk_orders_customers FOREIGN KEY (customer_id) REFERENCES northwind.customers(customer_id);


--
-- Name: orders fk_orders_employees; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.orders
    ADD CONSTRAINT fk_orders_employees FOREIGN KEY (employee_id) REFERENCES northwind.employees(employee_id);


--
-- Name: orders fk_orders_shippers; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.orders
    ADD CONSTRAINT fk_orders_shippers FOREIGN KEY (ship_via) REFERENCES northwind.shippers(shipper_id);


--
-- Name: products fk_products_categories; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.products
    ADD CONSTRAINT fk_products_categories FOREIGN KEY (category_id) REFERENCES northwind.categories(category_id);


--
-- Name: products fk_products_suppliers; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.products
    ADD CONSTRAINT fk_products_suppliers FOREIGN KEY (supplier_id) REFERENCES northwind.suppliers(supplier_id);


--
-- Name: territories fk_territories_region; Type: FK CONSTRAINT; Schema: northwind; Owner: -
--

ALTER TABLE ONLY northwind.territories
    ADD CONSTRAINT fk_territories_region FOREIGN KEY (region_id) REFERENCES northwind.region(region_id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20221124071125'),
    ('20221124071126'),
    ('20221127000046'),
    ('20221203171301');
