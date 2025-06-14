-- public.movie_casts definition



CREATE TABLE public.movie_casts (
	movie_id int4 NULL,
	casts_id int4 NULL
);


-- public.movie_casts foreign keys

ALTER TABLE public.movie_casts ADD CONSTRAINT movie_casts_casts_id_fkey FOREIGN KEY (casts_id) REFERENCES public.casts(id);
ALTER TABLE public.movie_casts ADD CONSTRAINT movie_casts_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movie(id);