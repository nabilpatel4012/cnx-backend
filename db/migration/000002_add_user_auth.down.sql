-- Reverting changes: Removing hashed_password and password_changed_at columns from the users table
ALTER TABLE IF EXISTS "users"
DROP COLUMN IF EXISTS "hashed_password",
DROP COLUMN IF EXISTS "password_changed_at";
