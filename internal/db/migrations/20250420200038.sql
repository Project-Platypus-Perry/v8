-- Create "batches" table
CREATE TABLE "public"."batches" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_batches_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_batches_deleted_at" to table: "batches"
CREATE INDEX "idx_batches_deleted_at" ON "public"."batches" ("deleted_at", "deleted_at");
