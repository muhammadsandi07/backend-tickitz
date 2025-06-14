-- public.users definition



CREATE TABLE public.users (
	id serial4 NOT NULL,
	email varchar(255) NULL,
	"password" varchar(255) NULL,
	"role" varchar(50) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);