CREATE TYPE "users_role_enum" AS ENUM ('ROLE_ADMIN', 'ROLE_AGENT', 'ROLE_READER');

CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "role" users_role_enum DEFAULT 'ROLE_READER',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "email" ON "users" ("email");