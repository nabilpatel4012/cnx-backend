-- Adding service_image column to the services table
ALTER TABLE IF EXISTS "services"
ADD COLUMN "service_image" varchar NOT NULL DEFAULT '';