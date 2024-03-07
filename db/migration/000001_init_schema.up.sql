CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "product_pack_sizes" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "product_line" varchar NOT NULL,
  "pack_size" bigint NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
