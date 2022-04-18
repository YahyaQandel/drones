--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

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
-- Name: dronetask; Type: SCHEMA; Schema: -; Owner: taskuser
--

CREATE SCHEMA dronetask;


ALTER SCHEMA dronetask OWNER TO taskuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: drone_logs; Type: TABLE; Schema: dronetask; Owner: taskuser
--

CREATE TABLE dronetask.drone_logs (
    id bigint NOT NULL,
    drone_serial_number text,
    drone_battery_level numeric,
    drone_state text,
    created_time timestamp with time zone
);


ALTER TABLE dronetask.drone_logs OWNER TO taskuser;

--
-- Name: drone_logs_id_seq; Type: SEQUENCE; Schema: dronetask; Owner: taskuser
--

CREATE SEQUENCE dronetask.drone_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dronetask.drone_logs_id_seq OWNER TO taskuser;

--
-- Name: drone_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: dronetask; Owner: taskuser
--

ALTER SEQUENCE dronetask.drone_logs_id_seq OWNED BY dronetask.drone_logs.id;


--
-- Name: drone_medications; Type: TABLE; Schema: dronetask; Owner: taskuser
--

CREATE TABLE dronetask.drone_medications (
    id bigint NOT NULL,
    drone_serial_number text,
    medication_code character varying(15)
);


ALTER TABLE dronetask.drone_medications OWNER TO taskuser;

--
-- Name: drone_medications_id_seq; Type: SEQUENCE; Schema: dronetask; Owner: taskuser
--

CREATE SEQUENCE dronetask.drone_medications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dronetask.drone_medications_id_seq OWNER TO taskuser;

--
-- Name: drone_medications_id_seq; Type: SEQUENCE OWNED BY; Schema: dronetask; Owner: taskuser
--

ALTER SEQUENCE dronetask.drone_medications_id_seq OWNED BY dronetask.drone_medications.id;


--
-- Name: drones; Type: TABLE; Schema: dronetask; Owner: taskuser
--

CREATE TABLE dronetask.drones (
    id bigint NOT NULL,
    serial_number text,
    model character varying(15),
    weight numeric,
    battery_capacity numeric,
    state text DEFAULT 'IDLE'::text
);


ALTER TABLE dronetask.drones OWNER TO taskuser;

--
-- Name: drones_id_seq; Type: SEQUENCE; Schema: dronetask; Owner: taskuser
--

CREATE SEQUENCE dronetask.drones_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dronetask.drones_id_seq OWNER TO taskuser;

--
-- Name: drones_id_seq; Type: SEQUENCE OWNED BY; Schema: dronetask; Owner: taskuser
--

ALTER SEQUENCE dronetask.drones_id_seq OWNED BY dronetask.drones.id;


--
-- Name: medications; Type: TABLE; Schema: dronetask; Owner: taskuser
--

CREATE TABLE dronetask.medications (
    id bigint NOT NULL,
    name text,
    weight numeric,
    code text,
    image text
);


ALTER TABLE dronetask.medications OWNER TO taskuser;

--
-- Name: medications_id_seq; Type: SEQUENCE; Schema: dronetask; Owner: taskuser
--

CREATE SEQUENCE dronetask.medications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dronetask.medications_id_seq OWNER TO taskuser;

--
-- Name: medications_id_seq; Type: SEQUENCE OWNED BY; Schema: dronetask; Owner: taskuser
--

ALTER SEQUENCE dronetask.medications_id_seq OWNED BY dronetask.medications.id;


--
-- Name: migration_records; Type: TABLE; Schema: dronetask; Owner: taskuser
--

CREATE TABLE dronetask.migration_records (
    id bigint NOT NULL,
    version_id bigint,
    t_stamp timestamp with time zone DEFAULT now(),
    is_applied boolean
);


ALTER TABLE dronetask.migration_records OWNER TO taskuser;

--
-- Name: migration_records_id_seq; Type: SEQUENCE; Schema: dronetask; Owner: taskuser
--

CREATE SEQUENCE dronetask.migration_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dronetask.migration_records_id_seq OWNER TO taskuser;

--
-- Name: migration_records_id_seq; Type: SEQUENCE OWNED BY; Schema: dronetask; Owner: taskuser
--

ALTER SEQUENCE dronetask.migration_records_id_seq OWNED BY dronetask.migration_records.id;


--
-- Name: drone_logs id; Type: DEFAULT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drone_logs ALTER COLUMN id SET DEFAULT nextval('dronetask.drone_logs_id_seq'::regclass);


--
-- Name: drone_medications id; Type: DEFAULT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drone_medications ALTER COLUMN id SET DEFAULT nextval('dronetask.drone_medications_id_seq'::regclass);


--
-- Name: drones id; Type: DEFAULT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drones ALTER COLUMN id SET DEFAULT nextval('dronetask.drones_id_seq'::regclass);


--
-- Name: medications id; Type: DEFAULT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.medications ALTER COLUMN id SET DEFAULT nextval('dronetask.medications_id_seq'::regclass);


--
-- Name: migration_records id; Type: DEFAULT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.migration_records ALTER COLUMN id SET DEFAULT nextval('dronetask.migration_records_id_seq'::regclass);


--
-- Data for Name: drone_logs; Type: TABLE DATA; Schema: dronetask; Owner: taskuser
--

COPY dronetask.drone_logs (id, drone_serial_number, drone_battery_level, drone_state, created_time) FROM stdin;
\.


--
-- Data for Name: drone_medications; Type: TABLE DATA; Schema: dronetask; Owner: taskuser
--

COPY dronetask.drone_medications (id, drone_serial_number, medication_code) FROM stdin;
1	firstDrone	XSF1
8	fourthDrone	PDL1
\.


--
-- Data for Name: drones; Type: TABLE DATA; Schema: dronetask; Owner: taskuser
--

COPY dronetask.drones (id, serial_number, model, weight, battery_capacity, state) FROM stdin;
2	firstDrone	Lightweight	60	60	IDLE
5	thirdDrone	Cruiserweight	90	90	IDLE
7	fourthDrone	Middleweight	40	2	LOADED
4	secondDrone	Heavyweight	30	2	LOADED
\.


--
-- Data for Name: medications; Type: TABLE DATA; Schema: dronetask; Owner: taskuser
--

COPY dronetask.medications (id, name, weight, code, image) FROM stdin;
8	panadol	20.6	PDL1	PDL1.png
9	rivo	4.6	RV75	RV75.jpg
\.


--
-- Data for Name: migration_records; Type: TABLE DATA; Schema: dronetask; Owner: taskuser
--

COPY dronetask.migration_records (id, version_id, t_stamp, is_applied) FROM stdin;
1	0	2022-04-08 07:03:08.211206+02	t
2	20220406023356	2022-04-08 07:03:08.822257+02	t
3	20220408014937	2022-04-08 07:03:09.463079+02	t
4	20220408065241	2022-04-08 07:03:10.076307+02	t
5	20220416063731	2022-04-16 06:43:09.480822+02	t
6	20220416070337	2022-04-16 07:07:04.828931+02	t
\.


--
-- Name: drone_logs_id_seq; Type: SEQUENCE SET; Schema: dronetask; Owner: taskuser
--

SELECT pg_catalog.setval('dronetask.drone_logs_id_seq', 109, true);


--
-- Name: drone_medications_id_seq; Type: SEQUENCE SET; Schema: dronetask; Owner: taskuser
--

SELECT pg_catalog.setval('dronetask.drone_medications_id_seq', 8, true);


--
-- Name: drones_id_seq; Type: SEQUENCE SET; Schema: dronetask; Owner: taskuser
--

SELECT pg_catalog.setval('dronetask.drones_id_seq', 7, true);


--
-- Name: medications_id_seq; Type: SEQUENCE SET; Schema: dronetask; Owner: taskuser
--

SELECT pg_catalog.setval('dronetask.medications_id_seq', 9, true);


--
-- Name: migration_records_id_seq; Type: SEQUENCE SET; Schema: dronetask; Owner: taskuser
--

SELECT pg_catalog.setval('dronetask.migration_records_id_seq', 6, true);


--
-- Name: drone_logs drone_logs_pkey; Type: CONSTRAINT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drone_logs
    ADD CONSTRAINT drone_logs_pkey PRIMARY KEY (id);


--
-- Name: drone_medications drone_medications_pkey; Type: CONSTRAINT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drone_medications
    ADD CONSTRAINT drone_medications_pkey PRIMARY KEY (id);


--
-- Name: drones drones_pkey; Type: CONSTRAINT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.drones
    ADD CONSTRAINT drones_pkey PRIMARY KEY (id);


--
-- Name: medications medications_pkey; Type: CONSTRAINT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.medications
    ADD CONSTRAINT medications_pkey PRIMARY KEY (id);


--
-- Name: migration_records migration_records_pkey; Type: CONSTRAINT; Schema: dronetask; Owner: taskuser
--

ALTER TABLE ONLY dronetask.migration_records
    ADD CONSTRAINT migration_records_pkey PRIMARY KEY (id);


--
-- Name: idx_drones_serial_number; Type: INDEX; Schema: dronetask; Owner: taskuser
--

CREATE UNIQUE INDEX idx_drones_serial_number ON dronetask.drones USING btree (serial_number, serial_number);


--
-- Name: idx_medications_code; Type: INDEX; Schema: dronetask; Owner: taskuser
--

CREATE UNIQUE INDEX idx_medications_code ON dronetask.medications USING btree (code, code);


--
-- PostgreSQL database dump complete
--

