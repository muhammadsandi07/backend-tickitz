-- public.movie definition



CREATE TABLE public.movie (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	duration int4 NULL,
	synopsis text NULL,
	img_movie varchar(255) NULL,
	backdrop varchar(255) NULL,
	release_date date NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL,
	CONSTRAINT movie_pkey PRIMARY KEY (id)
);