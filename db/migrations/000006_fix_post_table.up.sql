CREATE TABLE "like" (
                        id SERIAL PRIMARY KEY,
                        user_id UUID REFERENCES "user"(uuid),
                        post_id UUID REFERENCES "post"(id),
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);