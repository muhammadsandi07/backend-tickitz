-- public.casts definition

CREATE TABLE public.casts (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	CONSTRAINT casts_pkey PRIMARY KEY (id)
);