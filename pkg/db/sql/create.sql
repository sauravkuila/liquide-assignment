--drop tables for fresh schema
DROP TYPE IF EXISTS UserTypes;
DROP TABLE IF EXISTS user_detail;

--create types
CREATE TYPE UserTypes AS ENUM('USER','ADMIN');

-- create a function for timestamp
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--create tables
CREATE TABLE user_detail(
   id serial,
   user_name text not null unique,
   password text not null,
   user_type UserTypes not null,
   email text not null, 
   mobile text not null,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY(id)
);

-- create a trigger for timestamp
CREATE TRIGGER set_timestamp
AFTER UPDATE ON user_detail
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();