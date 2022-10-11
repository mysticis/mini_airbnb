CREATE TABLE "room_media" (
   "id" SERIAL PRIMARY KEY,
  "room_id" int NOT NULL,
  "room_images" varchar NOT NULL
);

ALTER TABLE "room_media" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;