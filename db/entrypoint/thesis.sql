-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2
-- PostgreSQL version: 12.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- object: thesis | type: DATABASE --
-- DROP DATABASE IF EXISTS thesis;
CREATE DATABASE thesis;
-- ddl-end --
\connect thesis;

-- object: public.thesis_thesis_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS public.thesis_thesis_id_seq CASCADE;
CREATE SEQUENCE public.thesis_thesis_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE public.thesis_thesis_id_seq OWNER TO postgres;
-- ddl-end --

-- object: public.thesis | type: TABLE --
-- DROP TABLE IF EXISTS public.thesis CASCADE;
CREATE TABLE thesis (
	thesis_id integer NOT NULL DEFAULT nextval('public.thesis_thesis_id_seq'::regclass),
	user_id integer NOT NULL,
	section_id integer NOT NULL,
	subject character varying(255) NOT NULL,
	fields text,
	CONSTRAINT thesis_pk PRIMARY KEY (thesis_id)

);
-- ddl-end --
-- ALTER TABLE public.thesis OWNER TO postgres;
-- ddl-end --

-- object: public."Status" | type: TYPE --
-- DROP TYPE IF EXISTS public."Status" CASCADE;
CREATE TYPE public."Status" AS
 ENUM ('Accepted','Declined','To refine','Considering');
-- ddl-end --
-- ALTER TYPE public."Status" OWNER TO postgres;
-- ddl-end --

-- object: public.changes_change_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS public.changes_change_id_seq CASCADE;
CREATE SEQUENCE public.changes_change_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE public.changes_change_id_seq OWNER TO postgres;
-- ddl-end --

-- object: public.changes | type: TABLE --
-- DROP TABLE IF EXISTS public.changes CASCADE;
CREATE TABLE public.changes (
	change_id integer NOT NULL DEFAULT nextval('public.changes_change_id_seq'::regclass),
	thesis_id integer NOT NULL,
	date timestamp with time zone NOT NULL,
	user_id integer NOT NULL,
	status public."Status" NOT NULL,
	admin_notes varchar(255),
	CONSTRAINT changes_pk PRIMARY KEY (change_id)

);
-- ddl-end --
-- ALTER TABLE public.changes OWNER TO postgres;
-- ddl-end --

-- object: changes_fk0 | type: CONSTRAINT --
-- ALTER TABLE public.changes DROP CONSTRAINT IF EXISTS changes_fk0 CASCADE;
ALTER TABLE public.changes ADD CONSTRAINT changes_fk0 FOREIGN KEY (thesis_id)
REFERENCES public.thesis (thesis_id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


