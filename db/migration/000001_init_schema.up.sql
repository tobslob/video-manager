CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "videos" (
  "id" uuid PRIMARY KEY NOT NULL,
  "url" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "duration" varchar NOT NULL,
  "title" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "metadatas" (
  "id" uuid PRIMARY KEY NOT NULL,
  "video_id" uuid UNIQUE NOT NULL,
  "width" int NOT NULL,
  "height" int NOT NULL,
  "file_type" varchar NOT NULL,
  "file_size" varchar,
  "last_modify" timestamptz NOT NULL,
  "accessed_date" timestamptz NOT NULL,
  "resolutions" int NOT NULL,
  "keywords" varchar,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "annotations" (
  "id" uuid PRIMARY KEY NOT NULL,
  "video_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "type" varchar NOT NULL,
  "note" varchar NOT NULL,
  "title" varchar NOT NULL,
  "label" varchar NOT NULL,
  "pause" boolean NOT NULL,
  "start_time" varchar NOT NULL,
  "end_time" varchar NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "videos" ("id");

CREATE INDEX ON "videos" ("user_id");

CREATE INDEX ON "metadatas" ("video_id");

CREATE INDEX ON "annotations" ("video_id");

CREATE INDEX ON "annotations" ("user_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "videos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "metadatas" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id");

ALTER TABLE "annotations" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id");
