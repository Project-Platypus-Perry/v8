-- Create "users_batches" table
CREATE TABLE "public"."users_batches" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" text NOT NULL,
  "batch_id" text NOT NULL,
  "organization_id" text NOT NULL,
  PRIMARY KEY ("id", "user_id", "batch_id"),
  CONSTRAINT "fk_users_batches_batch" FOREIGN KEY ("batch_id") REFERENCES "public"."batches" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_batches_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_batches_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_users_batches_deleted_at" to table: "users_batches"
CREATE INDEX "idx_users_batches_deleted_at" ON "public"."users_batches" ("deleted_at");
-- Create "users_classrooms" table
CREATE TABLE "public"."users_classrooms" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" text NOT NULL,
  "classroom_id" text NOT NULL,
  "organization_id" text NOT NULL,
  "batch_id" text NOT NULL,
  PRIMARY KEY ("id", "user_id", "classroom_id"),
  CONSTRAINT "fk_users_classrooms_batch" FOREIGN KEY ("batch_id") REFERENCES "public"."batches" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_classrooms_classroom" FOREIGN KEY ("classroom_id") REFERENCES "public"."classrooms" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_classrooms_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_users_classrooms_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_users_classrooms_deleted_at" to table: "users_classrooms"
CREATE INDEX "idx_users_classrooms_deleted_at" ON "public"."users_classrooms" ("deleted_at");
