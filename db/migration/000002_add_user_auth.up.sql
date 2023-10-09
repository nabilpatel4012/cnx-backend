-- Adding hashed_password and password_changed_at columns to the users table
ALTER TABLE "users"
ADD COLUMN "hashed_password" varchar NOT NULL,
ADD COLUMN "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z';
