CREATE TABLE IF NOT EXISTS "user" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" TEXT UNIQUE NOT NULL,
    "password" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "cpf" TEXT NOT NULL,
    "birth_date" DATE NOT NULL,
    "type" TEXT NOT NULL,
    "gender" TEXT,
    "phone_1" TEXT NOT NULL,
    "phone_2" TEXT,
    "create_at" DATE NOT NULL DEFAULT(now()),
    "updated_at" DATE NOT NULL DEFAULT(now()),
    "deleted_at" DATE
);

CREATE TABLE IF NOT EXISTS "appointment" (
    "id" BIGSERIAL PRIMARY KEY,
    "reserved" DATE,
    "day" INTEGER NOT NULL,
    "month" INTEGER NOT NULL,
    "year" INTEGER NOT NULL,
    "day_shift" TEXT NOT NULL,
    "animal_type" TEXT NOT NULL,
    "animal_name" TEXT,
    "animal_age" INTEGER,
    "user_id" BIGINT,
    "create_at" DATE NOT NULL DEFAULT(now()),
    "updated_at" DATE NOT NULL DEFAULT(now()),
    "deleted_at" DATE
);

ALTER TABLE "appointment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");