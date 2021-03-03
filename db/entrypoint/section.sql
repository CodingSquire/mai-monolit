-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2
-- PostgreSQL version: 12.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: section | type: DATABASE --
-- -- DROP DATABASE IF EXISTS section;
CREATE DATABASE section;
-- -- ddl-end --
-- 
\connect section;
-- object: public.sections | type: TABLE --
-- DROP TABLE IF EXISTS public.sections CASCADE;
CREATE TABLE public.sections (
	section_id integer NOT NULL,
	project_id integer NOT NULL,
	title varchar(255) NOT NULL,
	description text,
	date_create date NOT NULL,
	date_time_place text,
	thesis_opening_date date,
	thesis_ending_apply_date date,
	thesis_ending_edit_date date,
	thesis_finalize_date date,
	CONSTRAINT sections_pk PRIMARY KEY (section_id)

);
-- ddl-end --
-- ALTER TABLE public.sections OWNER TO postgres;
-- ddl-end --


