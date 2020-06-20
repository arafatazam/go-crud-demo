BEGIN;
CREATE TABLE "users" (
    "id" character(40) NOT NULL,
    "email" character(100) NOT NULL,
    "first_name" character(50) NOT NULL,
    "last_name" character(50),
    "date_of_birth" date NOT NULL,
    "address1" character(50) NOT NULL,
    "address2" character(50),
    "phone" character(20),
    CONSTRAINT "users_email" UNIQUE ("email"),
    CONSTRAINT "users_id" PRIMARY KEY ("id")
) WITH (oids = false);
COMMIT;