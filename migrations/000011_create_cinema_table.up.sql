-- public.cinema definition



CREATE TABLE public.cinema (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	CONSTRAINT cinema_pkey PRIMARY KEY (id)
);