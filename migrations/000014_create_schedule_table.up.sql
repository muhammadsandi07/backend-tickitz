-- public.schedule definition


CREATE TABLE public.schedule (
	id serial4 NOT NULL,
	id_movie int4 NULL,
	id_cinema int4 NULL,
	price int4 NULL,
	"date" timestamp NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL,
	CONSTRAINT schedule_pkey PRIMARY KEY (id)
);