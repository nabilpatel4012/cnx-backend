CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL,
  "refresh_token" varchar UNIQUE NOT NULL,
  "user_agent" varchar NOT NULL,
  "cleint_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");