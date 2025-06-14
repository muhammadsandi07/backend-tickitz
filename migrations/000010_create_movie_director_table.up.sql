-- public.movie_director definition



CREATE TABLE public.movie_director (
	movie_id int4 NULL,
	director_id int4 NULL
);


-- public.movie_director foreign keys

ALTER TABLE public.movie_director ADD CONSTRAINT movie_director_director_id_fkey FOREIGN KEY (director_id) REFERENCES public.director(id);
ALTER TABLE public.movie_director ADD CONSTRAINT movie_director_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movie(id);