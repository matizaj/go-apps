

create table public.test(id integer primary key, role_name varchar(50));
ALTER Table test owner to postgres;

INSERT INTO test("id", "role_name") VALUES (1, E'admin');