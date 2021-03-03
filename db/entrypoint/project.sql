-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2
-- PostgreSQL version: 12.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: project | type: DATABASE --
-- -- DROP DATABASE IF EXISTS project;
CREATE DATABASE project;
-- -- ddl-end --
-- 
\connect project;

-- object: public.projects | type: TABLE --
-- DROP TABLE IF EXISTS public.projects CASCADE;
CREATE TABLE public.projects (
	project_id integer NOT NULL,
	title varchar(255) NOT NULL,
	date_create timestamp with time zone NOT NULL,
	description text,
	thesis_opening_date timestamp with time zone NOT NULL,
	thesis_ending_apply_date timestamp with time zone NOT NULL,
	thesis_ending_edit_date timestamp with time zone NOT NULL,
	thesis_finalize_date timestamp with time zone NOT NULL,
	is_external boolean NOT NULL,
	external_domain varchar(255),
	internal_domain_id integer,
	multisectional boolean NOT NULL,
	CONSTRAINT projects_pk PRIMARY KEY (project_id)

);
-- ddl-end --
-- ALTER TABLE public.projects OWNER TO postgres;
-- ddl-end --


