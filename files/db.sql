CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	first_name varchar(255) NULL,
	last_name varchar(255) NULL,
	username varchar(255) NULL,
	"password" varchar(255) NULL,
	role_id varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.user_roles(id)
);

CREATE TABLE public.user_roles (
	id varchar NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar(255) NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	deleted_at timestamp NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);

INSERT INTO public.user_roles (id,"name",created_at,updated_at,deleted_at) VALUES
	 ('df7a6762-1d8d-4433-9735-365725abefec','admin','2022-12-10 18:34:14.845','2022-12-10 18:34:14.845',NULL),
	 ('2cbc1d37-974e-4afd-9ba1-5d510e96227d','user','2022-12-10 18:34:14.847','2022-12-10 18:34:14.847',NULL);

INSERT INTO public.users (id,first_name,last_name,username,"password",role_id,created_at,updated_at,deleted_at) VALUES
	 ('20bf7f46-afb3-49e2-a3fc-1955bc489559','adinata','pratama','adiadmin','$2a$10$hlQgetc/jXqwZC1YeLtSB.dNmitObulfTiIDaguCXP2feHc0gPLk.','df7a6762-1d8d-4433-9735-365725abefec','2022-12-10 11:43:26.140','2022-12-10 11:43:26.140',NULL),
	 ('cf63b5c8-6620-4db4-8a57-6f82e7463aa7','adinata','pratama','adi1','$2a$10$QsOStay1FAOXMohRUROU1eW62/yTY4PSSwgTIWojnxS3Dz5DugZh6','df7a6762-1d8d-4433-9735-365725abefec','2022-12-12 08:26:29.664','2022-12-12 08:26:29.664',NULL),
	 ('d1a06faa-c0c2-4af3-91bb-70d7d835dd33','adinata user','pratama','adiuser','$2a$10$pquF3z9g8M4s5nHR/bLvp.gPa92DjHvLODhneoRyknzaS55tXSgKS','2cbc1d37-974e-4afd-9ba1-5d510e96227d','2022-12-10 11:34:35.343','2022-12-12 17:11:37.623',NULL),
	 ('4e4e162f-4942-4b8c-9e69-9d56b0af494d','adinata','pratama','adi2','$2a$10$fd7hhRpQ6FWKG/DMvr8iuOImcFrnztrR2qtLS/chapk0DTQrIhcE6','df7a6762-1d8d-4433-9735-365725abefec','2022-12-12 10:19:38.025','2022-12-12 17:27:40.648','2022-12-12 10:27:40.648');
