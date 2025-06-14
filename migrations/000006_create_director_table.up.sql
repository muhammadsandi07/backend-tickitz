-- public.director definition


CREATE TABLE public.director (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	CONSTRAINT director_pkey PRIMARY KEY (id)
);