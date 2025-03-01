-- Create an additional database (if needed)
CREATE DATABASE wordpress;

-- Create an additional user (if needed)
CREATE USER wpuser WITH ENCRYPTED PASSWORD 'wppassword';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE wordpress TO wpuser;
