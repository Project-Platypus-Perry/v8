-- Modify "batches" table
ALTER TABLE "public"."batches" DROP CONSTRAINT "fk_batches_organization", DROP COLUMN "created_at", DROP COLUMN "updated_at", DROP COLUMN "deleted_at";
-- Drop index "idx_users_batches_deleted_at" from table: "users_batches"
DROP INDEX "public"."idx_users_batches_deleted_at";
-- Create index "idx_users_batches_deleted_at" to table: "users_batches"
CREATE INDEX "idx_users_batches_deleted_at" ON "public"."users_batches" ("deleted_at", "deleted_at");
