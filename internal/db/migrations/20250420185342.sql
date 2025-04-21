-- Drop index "idx_roles_deleted_at" from table: "roles"
DROP INDEX "public"."idx_roles_deleted_at";
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "public"."roles" ("deleted_at", "deleted_at");
