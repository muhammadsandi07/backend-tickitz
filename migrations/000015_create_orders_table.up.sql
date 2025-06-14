-- public.orders definition


CREATE TABLE public.orders (
	id serial4 NOT NULL,
	id_user int4 NULL,
	id_schedule int4 NULL,
	id_payment int4 NULL,
	ispaid bool DEFAULT false NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL,
	fullname varchar NULL,
	phone_number varchar NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
);