-- public.profile definition


CREATE TABLE public.profile (
	user_id int4 NULL,
	firstname varchar(255) NULL,
	lastname varchar(255) NULL,
	phone_number varchar(50) NULL,
	point int4 NULL,
	id_member int4 NULL,
	profile_image varchar(255) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp NULL
);


-- public.profile foreign keys

ALTER TABLE public.profile ADD CONSTRAINT profile_id_member_fkey FOREIGN KEY (id_member) REFERENCES public.tier(id);
ALTER TABLE public.profile ADD CONSTRAINT profile_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);