-- Drop index "idx_users_batches_deleted_at" from table: "users_batches"
DROP INDEX "public"."idx_users_batches_deleted_at";
-- Modify "users_batches" table
ALTER TABLE "public"."users_batches" DROP CONSTRAINT "users_batches_pkey", DROP COLUMN "id", ADD PRIMARY KEY ("user_id", "batch_id");
-- Create index "idx_users_batches_deleted_at" to table: "users_batches"
CREATE INDEX "idx_users_batches_deleted_at" ON "public"."users_batches" ("deleted_at");
