-- DDL generated by Postico 1.5.17
-- Not all database features are supported. Do not use for backup.

-- Table Definition ----------------------------------------------

CREATE TABLE users (
    userid text PRIMARY KEY,
    name text
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX users_pkey ON users(userid text_ops);

---Insert Data-----

INSERT INTO "users"("userid", "name") VALUES ('01', 'Budi');
INSERT INTO "users"("userid", "name") VALUES ('02', 'Nano');