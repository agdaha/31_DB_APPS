DROP TABLE IF EXISTS public.task_labels;
DROP TABLE IF EXISTS public.tasks;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.labels;

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL COLLATE pg_catalog."default"
);


CREATE TABLE IF NOT EXISTS public.labels
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL COLLATE pg_catalog."default"
);


CREATE TABLE IF NOT EXISTS public.tasks
(
    id SERIAL PRIMARY KEY,
    opened bigint NOT NULL DEFAULT EXTRACT(epoch FROM now()),
    closed bigint DEFAULT 0,
    author_id integer,
    assigned_id integer,
    title text COLLATE pg_catalog."default",
    content text COLLATE pg_catalog."default",
    CONSTRAINT assigned_id FOREIGN KEY (assigned_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT author_id FOREIGN KEY (author_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);


CREATE TABLE IF NOT EXISTS public.task_labels
(
    task_id integer NOT NULL,
    label_id integer NOT NULL,
    CONSTRAINT label_id FOREIGN KEY (label_id)
        REFERENCES public.labels (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT task_id FOREIGN KEY (task_id)
        REFERENCES public.tasks (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);


INSERT INTO public.users(id, name) VALUES (1, 'author_1');
INSERT INTO public.users(id, name) VALUES (2, 'author_2');

INSERT INTO public.labels(id, name) VALUES (1, 'IT');
INSERT INTO public.labels(id, name) VALUES (2, 'ML');
INSERT INTO public.labels(id, name) VALUES (3, 'DB');

INSERT INTO public.tasks(id, author_id, assigned_id, title, content)
	VALUES (1, 1, 2, 'Task 1', 'search info');
INSERT INTO public.tasks(id, author_id, assigned_id, title, content)
	VALUES (2, 2, 1, 'Task 2', 'schema db');

INSERT INTO public.task_labels(task_id, label_id) VALUES (1, 1);
INSERT INTO public.task_labels(task_id, label_id) VALUES (2, 2);
INSERT INTO public.task_labels(task_id, label_id) VALUES (2, 3);

SELECT setval('labels_id_seq', 4, true);
SELECT setval('users_id_seq', 3, true);
SELECT setval('tasks_id_seq', 3, true);