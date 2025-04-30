-- Modify "users" table
ALTER TABLE "public"."users" DROP CONSTRAINT "fk_users_organization", ADD CONSTRAINT "fk_users_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
