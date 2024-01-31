
-- Add foreign key reference to comment table
ALTER TABLE image
    ADD COLUMN file_path VARCHAR(255);
