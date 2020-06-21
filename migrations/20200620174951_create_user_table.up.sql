BEGIN;
CREATE TABLE "users" (
    "id" varchar(40) NOT NULL,
    "email" varchar(100) NOT NULL,
    "first_name" varchar(50) NOT NULL,
    "last_name" varchar(50),
    "address" varchar(50) NOT NULL,
    "phone" varchar(20),
    CONSTRAINT "users_email" UNIQUE ("email"),
    CONSTRAINT "users_id" PRIMARY KEY ("id")
) WITH (oids = false);
COMMIT;