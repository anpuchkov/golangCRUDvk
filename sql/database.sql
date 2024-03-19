--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6 (Homebrew)
-- Dumped by pg_dump version 15.6 (Homebrew)

-- Started on 2024-03-18 22:37:11 MSK

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
-- TOC entry 5 (class 2615 OID 2200)
-- Name: filmoteka; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA filmoteka;


--
-- TOC entry 3643 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA filmoteka; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA filmoteka IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16473)
-- Name: actors; Type: TABLE; Schema: filmoteka; Owner: -
--

CREATE TABLE filmoteka.actors (
                                  id integer NOT NULL,
                                  name character varying(255),
                                  gender character varying(10),
                                  birthdate date
);


--
-- TOC entry 214 (class 1259 OID 16472)
-- Name: actors_id_seq; Type: SEQUENCE; Schema: filmoteka; Owner: -
--

CREATE SEQUENCE filmoteka.actors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3646 (class 0 OID 0)
-- Dependencies: 214
-- Name: actors_id_seq; Type: SEQUENCE OWNED BY; Schema: filmoteka; Owner: -
--

ALTER SEQUENCE filmoteka.actors_id_seq OWNED BY filmoteka.actors.id;


--
-- TOC entry 218 (class 1259 OID 16488)
-- Name: actorsmovies; Type: TABLE; Schema: filmoteka; Owner: -
--

CREATE TABLE filmoteka.actorsmovies (
                                        actor_id integer NOT NULL,
                                        movie_id integer NOT NULL
);


--
-- TOC entry 217 (class 1259 OID 16480)
-- Name: movies; Type: TABLE; Schema: filmoteka; Owner: -
--

CREATE TABLE filmoteka.movies (
                                  id integer NOT NULL,
                                  title character varying(150) NOT NULL,
                                  description text NOT NULL,
                                  release_date date NOT NULL,
                                  rating numeric(3,1) NOT NULL
);


--
-- TOC entry 216 (class 1259 OID 16479)
-- Name: movies_id_seq; Type: SEQUENCE; Schema: filmoteka; Owner: -
--

CREATE SEQUENCE filmoteka.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3649 (class 0 OID 0)
-- Dependencies: 216
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: filmoteka; Owner: -
--

ALTER SEQUENCE filmoteka.movies_id_seq OWNED BY filmoteka.movies.id;


--
-- TOC entry 220 (class 1259 OID 16504)
-- Name: users; Type: TABLE; Schema: filmoteka; Owner: -
--

CREATE TABLE filmoteka.users (
                                 id integer NOT NULL,
                                 username character varying(255),
                                 password character varying(255),
                                 registration_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- TOC entry 219 (class 1259 OID 16503)
-- Name: users_id_seq; Type: SEQUENCE; Schema: filmoteka; Owner: -
--

CREATE SEQUENCE filmoteka.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3651 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: filmoteka; Owner: -
--

ALTER SEQUENCE filmoteka.users_id_seq OWNED BY filmoteka.users.id;


--
-- TOC entry 3475 (class 2604 OID 16476)
-- Name: actors id; Type: DEFAULT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.actors ALTER COLUMN id SET DEFAULT nextval('filmoteka.actors_id_seq'::regclass);


--
-- TOC entry 3476 (class 2604 OID 16483)
-- Name: movies id; Type: DEFAULT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.movies ALTER COLUMN id SET DEFAULT nextval('filmoteka.movies_id_seq'::regclass);


--
-- TOC entry 3477 (class 2604 OID 16507)
-- Name: users id; Type: DEFAULT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.users ALTER COLUMN id SET DEFAULT nextval('filmoteka.users_id_seq'::regclass);


--
-- TOC entry 3632 (class 0 OID 16473)
-- Dependencies: 215
-- Data for Name: actors; Type: TABLE DATA; Schema: filmoteka; Owner: -
--

COPY filmoteka.actors (id, name, gender, birthdate) FROM stdin;
1	Илья Муромец	Мужчина	2024-03-23
\.


--
-- TOC entry 3635 (class 0 OID 16488)
-- Dependencies: 218
-- Data for Name: actorsmovies; Type: TABLE DATA; Schema: filmoteka; Owner: -
--

COPY filmoteka.actorsmovies (actor_id, movie_id) FROM stdin;
1	1
\.


--
-- TOC entry 3634 (class 0 OID 16480)
-- Dependencies: 217
-- Data for Name: movies; Type: TABLE DATA; Schema: filmoteka; Owner: -
--

COPY filmoteka.movies (id, title, description, release_date, rating) FROM stdin;
1	Муму	Фильм про собаку и человека	1905-12-12	5.0
2	Человек	Фильм про человека	2001-12-11	10.0
3	The Shawshank Redemption	Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.	2024-03-18	9.3
4	Fast and Furious	There are a lot of different actions	2024-03-18	6.0
\.


--
-- TOC entry 3637 (class 0 OID 16504)
-- Dependencies: 220
-- Data for Name: users; Type: TABLE DATA; Schema: filmoteka; Owner: -
--

COPY filmoteka.users (id, username, password, email, registration_date) FROM stdin;
\.


--
-- TOC entry 3652 (class 0 OID 0)
-- Dependencies: 214
-- Name: actors_id_seq; Type: SEQUENCE SET; Schema: filmoteka; Owner: -
--

SELECT pg_catalog.setval('filmoteka.actors_id_seq', 1, true);


--
-- TOC entry 3653 (class 0 OID 0)
-- Dependencies: 216
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: filmoteka; Owner: -
--

SELECT pg_catalog.setval('filmoteka.movies_id_seq', 4, true);


--
-- TOC entry 3654 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: filmoteka; Owner: -
--

SELECT pg_catalog.setval('filmoteka.users_id_seq', 1, false);


--
-- TOC entry 3480 (class 2606 OID 16478)
-- Name: actors actors_pkey; Type: CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.actors
    ADD CONSTRAINT actors_pkey PRIMARY KEY (id);


--
-- TOC entry 3484 (class 2606 OID 16492)
-- Name: actorsmovies actorsmovies_pkey; Type: CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.actorsmovies
    ADD CONSTRAINT actorsmovies_pkey PRIMARY KEY (actor_id, movie_id);


--
-- TOC entry 3482 (class 2606 OID 16487)
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- TOC entry 3486 (class 2606 OID 16512)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3487 (class 2606 OID 16493)
-- Name: actorsmovies actorsmovies_actor_id_fkey; Type: FK CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.actorsmovies
    ADD CONSTRAINT actorsmovies_actor_id_fkey FOREIGN KEY (actor_id) REFERENCES filmoteka.actors(id);


--
-- TOC entry 3488 (class 2606 OID 16498)
-- Name: actorsmovies actorsmovies_movie_id_fkey; Type: FK CONSTRAINT; Schema: filmoteka; Owner: -
--

ALTER TABLE ONLY filmoteka.actorsmovies
    ADD CONSTRAINT actorsmovies_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES filmoteka.movies(id);


--
-- TOC entry 3644 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA filmoteka; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON SCHEMA filmoteka TO andrejpuckov;
GRANT ALL ON SCHEMA filmoteka TO admin;
GRANT USAGE ON SCHEMA filmoteka TO "user";


--
-- TOC entry 3645 (class 0 OID 0)
-- Dependencies: 215
-- Name: TABLE actors; Type: ACL; Schema: filmoteka; Owner: -
--

GRANT ALL ON TABLE filmoteka.actors TO admin;
GRANT SELECT ON TABLE filmoteka.actors TO "user";


--
-- TOC entry 3647 (class 0 OID 0)
-- Dependencies: 218
-- Name: TABLE actorsmovies; Type: ACL; Schema: filmoteka; Owner: -
--

GRANT ALL ON TABLE filmoteka.actorsmovies TO admin;
GRANT SELECT ON TABLE filmoteka.actorsmovies TO "user";


--
-- TOC entry 3648 (class 0 OID 0)
-- Dependencies: 217
-- Name: TABLE movies; Type: ACL; Schema: filmoteka; Owner: -
--

GRANT ALL ON TABLE filmoteka.movies TO admin;
GRANT SELECT ON TABLE filmoteka.movies TO "user";


--
-- TOC entry 3650 (class 0 OID 0)
-- Dependencies: 220
-- Name: TABLE users; Type: ACL; Schema: filmoteka; Owner: -
--

GRANT ALL ON TABLE filmoteka.users TO admin;


-- Completed on 2024-03-18 22:37:11 MSK

--
-- PostgreSQL database dump complete
--

