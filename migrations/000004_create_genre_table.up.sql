-- public.genre definition


CREATE TABLE public.genre (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	CONSTRAINT genre_pkey PRIMARY KEY (id)
);