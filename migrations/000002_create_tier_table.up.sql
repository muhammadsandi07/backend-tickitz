-- public.tier definition


CREATE TABLE public.tier (
	id serial4 NOT NULL,
	"name" varchar(100) NULL,
	max_point int4 NULL,
	CONSTRAINT tier_pkey PRIMARY KEY (id)
);