CREATE TABLE "users" (
                         "id" varchar(50) UNIQUE PRIMARY KEY NOT NULL,
                         "username" varchar(50) UNIQUE NOT NULL,
                         "display_name" varchar(50) NOT NULL,
                         "email" varchar(255) UNIQUE NOT NULL,
                         "password" varchar(255) NOT NULL,
                         created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
                         updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);
