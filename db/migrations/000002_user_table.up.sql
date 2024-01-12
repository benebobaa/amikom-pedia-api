-- Enable uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Main table
CREATE TABLE "user" (
                        uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                        email VARCHAR(255) UNIQUE NOT NULL,
                        nim VARCHAR(255) UNIQUE NOT NULL,
                        name VARCHAR(255) NOT NULL,
                        username VARCHAR(255) UNIQUE NOT NULL,
                        bio TEXT,
                        password VARCHAR(255) NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Image table
CREATE TABLE "image" (
                         id SERIAL PRIMARY KEY,
                         user_uuid UUID REFERENCES "user"(uuid),
                         image_type VARCHAR(50),
                         image_url VARCHAR(255),
                         created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
