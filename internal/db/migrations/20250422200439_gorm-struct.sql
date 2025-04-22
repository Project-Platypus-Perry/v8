-- Drop index "idx_batches_deleted_at" from table: "batches"
DROP INDEX "public"."idx_batches_deleted_at";
-- Create index "idx_batches_deleted_at" to table: "batches"
CREATE INDEX "idx_batches_deleted_at" ON "public"."batches" ("deleted_at");
-- Drop index "idx_chapters_deleted_at" from table: "chapters"
DROP INDEX "public"."idx_chapters_deleted_at";
-- Create index "idx_chapters_deleted_at" to table: "chapters"
CREATE INDEX "idx_chapters_deleted_at" ON "public"."chapters" ("deleted_at");
-- Drop index "idx_classrooms_deleted_at" from table: "classrooms"
DROP INDEX "public"."idx_classrooms_deleted_at";
-- Create index "idx_classrooms_deleted_at" to table: "classrooms"
CREATE INDEX "idx_classrooms_deleted_at" ON "public"."classrooms" ("deleted_at");
-- Drop index "idx_organizations_deleted_at" from table: "organizations"
DROP INDEX "public"."idx_organizations_deleted_at";
-- Create index "idx_organizations_deleted_at" to table: "organizations"
CREATE INDEX "idx_organizations_deleted_at" ON "public"."organizations" ("deleted_at");
-- Drop index "idx_roles_deleted_at" from table: "roles"
DROP INDEX "public"."idx_roles_deleted_at";
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "public"."roles" ("deleted_at");
