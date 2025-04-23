-- Create enum type "content_type"
CREATE TYPE "public"."content_type" AS ENUM ('notes', 'dpp', 'video');
-- Create enum type "language"
CREATE TYPE "public"."language" AS ENUM ('en', 'hi');
-- Create enum type "visibility"
CREATE TYPE "public"."visibility" AS ENUM ('public', 'private');
-- Create "contents" table
CREATE TABLE "public"."contents" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "chapter_id" text NOT NULL,
  "type" "public"."content_type" NULL DEFAULT 'notes',
  "name" text NOT NULL,
  "description" text NULL,
  "language" "public"."language" NOT NULL DEFAULT 'en',
  "visibility" "public"."visibility" NOT NULL DEFAULT 'public',
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_contents_chapter" FOREIGN KEY ("chapter_id") REFERENCES "public"."chapters" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_contents_deleted_at" to table: "contents"
CREATE INDEX "idx_contents_deleted_at" ON "public"."contents" ("deleted_at");
