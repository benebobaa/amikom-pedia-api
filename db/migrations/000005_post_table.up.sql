

-- Create Post table with user_id and ref_post_id
CREATE TABLE "post" (
                        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                        content TEXT NOT NULL,
                        user_id UUID REFERENCES "user"(uuid), -- Reference to the user table
                        ref_post_id UUID, -- Self-referencing column (nullable)
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key reference to comment table
ALTER TABLE image
    ADD COLUMN post_id UUID REFERENCES post(id);
