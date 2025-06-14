-- public."location" definition


CREATE TABLE public."location" (
	id serial4 NOT NULL,
	"name" varchar(255) NULL,
	CONSTRAINT location_pkey PRIMARY KEY (id)
);