-- public.orders_seats definition



CREATE TABLE public.orders_seats (
	orders_id int4 NULL,
	seats_id varchar NULL
);


-- public.orders_seats foreign keys

ALTER TABLE public.orders_seats ADD CONSTRAINT orders_seats_orders_id_fkey FOREIGN KEY (orders_id) REFERENCES public.orders(id);
