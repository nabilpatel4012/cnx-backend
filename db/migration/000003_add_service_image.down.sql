-- Reverting changes: Removing service_image column from the services table
ALTER TABLE IF EXISTS "services"
DROP COLUMN IF EXISTS "service_image";