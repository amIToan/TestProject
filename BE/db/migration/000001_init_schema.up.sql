CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL, 
  "email" varchar UNIQUE NOT NULL,
  "role" varchar(10) CHECK (role IN ('admin', 'reader')), 
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
); 

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "dob" DATE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "gender" varchar(10) NOT NULL 
);

CREATE TABLE "books" (
  "book_id" bigserial PRIMARY KEY, 
  "title" varchar NOT NULL, 
  "description" varchar NOT NULL,
  "image_links" TEXT[], 
  "price" NUMERIC(10, 2) NOT NULL,
  "publication_date" DATE NOT NULL, 
  "author_id" bigint NOT NULL,
  FOREIGN KEY ("author_id") REFERENCES "authors" ("id")
);

CREATE TABLE "genres" (
  "genre_id" bigserial PRIMARY KEY, 
  "name" varchar NOT NULL
);

CREATE TABLE "book_genres" ( 
  "book_id" bigint NOT NULL,
  "genre_id" bigint NOT NULL,
  FOREIGN KEY ("book_id") REFERENCES "books" ("book_id"),
  FOREIGN KEY ("genre_id") REFERENCES "genres" ("genre_id")
);


CREATE INDEX ON "books" ("title", "author_id");

CREATE INDEX ON "book_genres" ("book_id", "genre_id");

CREATE OR REPLACE FUNCTION update_password_changed_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.password_changed_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_password_trigger
BEFORE UPDATE OF hashed_password ON users
FOR EACH ROW
EXECUTE FUNCTION update_password_changed_at();

