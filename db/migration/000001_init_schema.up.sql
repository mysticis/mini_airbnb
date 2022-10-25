CREATE TYPE "home_size" AS ENUM (
  '15m_squared_or_more',
  '30m_squared_or_more',
  '60m_squared_or_more',
  '90m_squared_or_more'
);

CREATE TABLE "tenant" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "hashed_password" varchar NOT NULL
);

CREATE TABLE "landlord" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL
);

CREATE TABLE "rooms" (
  "id" SERIAL PRIMARY KEY,
  "email_id" varchar not null,
  "owner_id" int NOT NULL,
  "home_type" varchar[] NOT NULL,
  "home_size" home_size NOT NULL,
  "furnished" boolean NOT NULL,
  "private_bathroom" boolean NOT NULL,
  "balcony" boolean NOT NULL,
  "garden" boolean NOT NULL,
  "kitchen" boolean NOT NULL,
  "pets_allowed" boolean NOT NULL,
  "parking" boolean NOT NULL,
  "wheelchair_accessible" boolean NOT NULL,
  "basement" boolean NOT NULL,
  "amenities" text[] NOT NULL,
  "suitable_for" text[] NOT NULL,
  "published_at" timestamptz NOT NULL DEFAULT (now()),
  "price" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "longitude" float NOT NULL,
  "latitude" float NOT NULL
);

CREATE TABLE "reservations" (
   "id" SERIAL PRIMARY KEY,
   "email_id" varchar NOT NULL,
  "tenant_id" int NOT NULL,
  "room_id" int NOT NULL,
  "duration" tsrange,
  "price" float NOT NULL,
  "total" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  EXCLUDE USING GIST (duration WITH &&)
);


CREATE TABLE "reviews" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "room_id" int NOT NULL,
  "comment" varchar NOT NULL,
  "rating" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "admin" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL
);

ALTER TABLE "rooms" ADD FOREIGN KEY ("owner_id") REFERENCES "landlord" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("email_id") REFERENCES "tenant" ("email") ON DELETE CASCADE;
ALTER TABLE "reservations" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;
ALTER TABLE "rooms" ADD FOREIGN KEY ("email_id") REFERENCES "landlord" ("email") ON DELETE CASCADE;
ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "tenant" ("id") ON DELETE CASCADE;

ALTER TABLE "reviews" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;
