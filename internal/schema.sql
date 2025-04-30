-- Create enum types
CREATE TYPE content_type AS ENUM ('notes', 'dpp', 'video');
CREATE TYPE language AS ENUM ('en', 'hi');
CREATE TYPE visibility AS ENUM ('public', 'private');
CREATE TYPE user_role AS ENUM ('admin', 'instructor', 'student');
