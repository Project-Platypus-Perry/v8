-- Create enum type "user_role"
CREATE TYPE "public"."user_role" AS ENUM ('admin', 'instructor', 'student');
-- Modify "roles" table
ALTER TABLE "public"."roles" ALTER COLUMN "role" TYPE "public"."user_role" USING CASE WHEN "role" IN ('admin', 'instructor', 'student') THEN "role"::"public"."user_role" ELSE NULL END, ALTER COLUMN "role" DROP NOT NULL;
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "organization_id" text NOT NULL, ADD COLUMN "role" "public"."user_role" NULL, ADD CONSTRAINT "fk_users_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
