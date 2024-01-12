-- User Registration Table
CREATE TABLE "user_registration" (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     email VARCHAR(255) NOT NULL,
                                     nim VARCHAR(10) NOT NULL,
                                     password VARCHAR(255) NOT NULL,
                                     is_verified BOOLEAN DEFAULT FALSE,
                                     email_verified_at TIMESTAMPTZ,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- OTP Table
CREATE TABLE "otp" (
                       id SERIAL PRIMARY KEY,
                       user_rid INT REFERENCES "user_registration"(id),
                       user_id UUID REFERENCES "user"(uuid),
                       otp_value VARCHAR(6) NOT NULL,
                       expired_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       ref_code VARCHAR(16) NOT NULL
);
