CREATE TABLE "users" (
  "user_id" serial UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar NOT NULL DEFAULT (now()),
  "total_orders" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "services" (
  "service_id" serial PRIMARY KEY NOT NULL,
  "service_name" varchar UNIQUE NOT NULL,
  "service_price" bigint NOT NULL
);

CREATE TABLE "orders" (
  "order_id" serial UNIQUE PRIMARY KEY NOT NULL,
  "customer_id" serial NOT NULL,
  "service_id" serial NOT NULL,
  "order_status" varchar NOT NULL,
  "order_started" timestamptz NOT NULL DEFAULT (now()),
  "order_delivered" timestamptz
);

CREATE INDEX ON "users" ("user_id");

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "orders" ("order_id");

CREATE INDEX ON "orders" ("order_status");

CREATE INDEX ON "orders" ("order_started", "order_status");

COMMENT ON COLUMN "users"."user_id" IS 'this will consist of unique user_id';

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("service_id") REFERENCES "services" ("service_id");
