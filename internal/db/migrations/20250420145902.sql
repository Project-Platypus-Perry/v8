-- Drop index "idx_organizations_deleted_at" from table: "organizations"
DROP INDEX "public"."idx_organizations_deleted_at";
-- Create index "idx_organizations_deleted_at" to table: "organizations"
CREATE INDEX "idx_organizations_deleted_at" ON "public"."organizations" ("deleted_at", "deleted_at");
