-- Reverting changes: Removing password_changed_at column from the users table
ALTER TABLE IF EXISTS "users"
DROP COLUMN IF EXISTS "password_changed_at";
