CREATE TABLE public.users (
	id serial4 NOT NULL,
	name varchar(255) NULL,
	email varchar(255) NULL,
	address text NULL,
	"password" varchar(255) NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);