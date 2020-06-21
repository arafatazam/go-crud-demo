BEGIN;
CREATE TABLE "users" (
    "id" character(40) NOT NULL,
    "email" character(100) NOT NULL,
    "first_name" character(50) NOT NULL,
    "last_name" character(50),
    "address" character(50) NOT NULL,
    "phone" character(20),
    CONSTRAINT "users_email" UNIQUE ("email"),
    CONSTRAINT "users_id" PRIMARY KEY ("id")
) WITH (oids = false);
COMMIT;