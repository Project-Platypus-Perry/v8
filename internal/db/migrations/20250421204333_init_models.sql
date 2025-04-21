-- Modify "batches" table
ALTER TABLE "public"."batches" ADD COLUMN "name" text NOT NULL, ADD COLUMN "description" text NULL;
-- Create "classrooms" table
CREATE TABLE "public"."classrooms" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" text NOT NULL,
  "batch_id" text NOT NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "settings" jsonb NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_classrooms_batch" FOREIGN KEY ("batch_id") REFERENCES "public"."batches" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_classrooms_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_classrooms_deleted_at" to table: "classrooms"
CREATE INDEX "idx_classrooms_deleted_at" ON "public"."classrooms" ("deleted_at", "deleted_at");
-- Create "chapters" table
CREATE TABLE "public"."chapters" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "classroom_id" text NOT NULL,
  "name" text NOT NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_chapters_classroom" FOREIGN KEY ("classroom_id") REFERENCES "public"."classrooms" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_chapters_deleted_at" to table: "chapters"
CREATE INDEX "idx_chapters_deleted_at" ON "public"."chapters" ("deleted_at", "deleted_at");
-- Drop enum type "content_type"
DROP TYPE "public"."content_type";
-- Drop enum type "language"
DROP TYPE "public"."language";
-- Drop enum type "visibility"
DROP TYPE "public"."visibility";
