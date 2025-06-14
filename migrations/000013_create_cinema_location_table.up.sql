-- public.cinema_location definition


CREATE TABLE public.cinema_location (
	id serial4 NOT NULL,
	cinema_id int4 NULL,
	location_id int4 NULL,
	address varchar(255) NULL,
	CONSTRAINT cinema_location_pkey PRIMARY KEY (id)
);


-- public.cinema_location foreign keys

ALTER TABLE public.cinema_location ADD CONSTRAINT cinema_location_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinema(id);
ALTER TABLE public.cinema_location ADD CONSTRAINT cinema_location_location_id_fkey FOREIGN KEY (location_id) REFERENCES public."location"(id);