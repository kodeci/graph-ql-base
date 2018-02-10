--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: ichabod; Type: DATABASE; Schema: -; Owner: ichabod
--
DROP DATABASE IF EXISTS ichabod;

CREATE DATABASE ichabod WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE ichabod OWNER TO ichabod;

\connect ichabod

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


CREATE FUNCTION created_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
	NEW.updated_at = NOW();
	NEW.created_at = NOW();
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.created_at_column() OWNER TO ichabod;

--
-- TOC entry 190 (class 1255 OID 36646)
-- Name: update_at_column(); Type: FUNCTION; Schema: public; Owner: ichabod
--

CREATE FUNCTION update_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.update_at_column() OWNER TO ichabod;


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: applications; Type: TABLE; Schema: public; Owner: ichabod; Tablespace:
--

CREATE TABLE applications (
    id integer NOT NULL,
    title character varying,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE applications OWNER TO ichabod;

--
-- Name: applications_id_seq; Type: SEQUENCE; Schema: public; Owner: ichabod
--

CREATE SEQUENCE applications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE applications_id_seq OWNER TO ichabod;

--
-- Name: applications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ichabod
--

ALTER SEQUENCE applications_id_seq OWNED BY applications.id;


--
-- Name: environments; Type: TABLE; Schema: public; Owner: ichabod; Tablespace:
--

CREATE TABLE environments (
    id integer NOT NULL,
    application_id integer,
    title character varying,
    slug character varying,
    values json,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE environments OWNER TO ichabod;

--
-- Name: environments_id_seq; Type: SEQUENCE; Schema: public; Owner: ichabod
--

CREATE SEQUENCE environments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE environments_id_seq OWNER TO ichabod;

--
-- Name: environments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ichabod
--

ALTER SEQUENCE environments_id_seq OWNED BY environments.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: ichabod
--

ALTER TABLE ONLY applications ALTER COLUMN id SET DEFAULT nextval('applications_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: ichabod
--

ALTER TABLE ONLY environments ALTER COLUMN id SET DEFAULT nextval('environments_id_seq'::regclass);


--
-- Data for Name: applications; Type: TABLE DATA; Schema: public; Owner: ichabod
--

COPY applications (id, title, updated_at, created_at) FROM stdin;
\.


--
-- Name: applications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ichabod
--

SELECT pg_catalog.setval('applications_id_seq', 1, false);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: ichabod
--

COPY environments (id, application_id, title, slug, values, updated_at, created_at) FROM stdin;
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ichabod
--

SELECT pg_catalog.setval('environments_id_seq', 1, false);


--
-- Name: applications_id; Type: CONSTRAINT; Schema: public; Owner: ichabod; Tablespace:
--

ALTER TABLE ONLY applications
    ADD CONSTRAINT applications_id PRIMARY KEY (id);


--
-- Name: user_id; Type: CONSTRAINT; Schema: public; Owner: ichabod; Tablespace:
--

ALTER TABLE ONLY environments
    ADD CONSTRAINT environments_id PRIMARY KEY (id);


--
-- Name: applications_user_id; Type: FK CONSTRAINT; Schema: public; Owner: ichabod
--

ALTER TABLE ONLY environments
    ADD CONSTRAINT applications_environment_id FOREIGN KEY (application_id) REFERENCES applications(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2284 (class 2620 OID 36647)
-- Name: applications create_applications_created_at; Type: TRIGGER; Schema: public; Owner: ichabod
--

CREATE TRIGGER create_applications_created_at BEFORE INSERT ON applications FOR EACH ROW EXECUTE PROCEDURE created_at_column();


--
-- TOC entry 2286 (class 2620 OID 36653)
-- Name: user create_user_created_at; Type: TRIGGER; Schema: public; Owner: ichabod
--

CREATE TRIGGER create_environments_created_at BEFORE INSERT ON environments FOR EACH ROW EXECUTE PROCEDURE created_at_column();


--
-- TOC entry 2285 (class 2620 OID 36648)
-- Name: applications update_applications_updated_at; Type: TRIGGER; Schema: public; Owner: ichabod
--

CREATE TRIGGER update_applications_updated_at BEFORE UPDATE ON applications FOR EACH ROW EXECUTE PROCEDURE update_at_column();


--
-- TOC entry 2287 (class 2620 OID 36654)
-- Name: user update_user_updated_at; Type: TRIGGER; Schema: public; Owner: ichabod
--

CREATE TRIGGER update_environments_updated_at BEFORE UPDATE ON environments FOR EACH ROW EXECUTE PROCEDURE update_at_column();



--
-- Name: public; Type: ACL; Schema: -; Owner: ichabod
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM ichabod;
GRANT ALL ON SCHEMA public TO ichabod;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--
