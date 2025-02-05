--drop tables for fresh schema
DROP TYPE IF EXISTS UserTypes;
DROP TABLE IF EXISTS user_detail;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS vote;

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

CREATE TABLE post (
    post_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user_detail(id),
    content TEXT NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comment (
    comment_id SERIAL PRIMARY KEY,
    post_id INT REFERENCES post(post_id),
    user_id INT REFERENCES user_detail(id),
    parent_comment_id INT REFERENCES comment(comment_id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vote (
    vote_id SERIAL PRIMARY KEY,
    post_id INT REFERENCES post(post_id),
    user_id INT REFERENCES user_detail(id),
    vote_type BOOLEAN NOT NULL, -- TRUE for upvote, FALSE for downvote
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (post_id, user_id)
);

-- create a trigger for timestamp
CREATE TRIGGER set_timestamp
AFTER UPDATE ON user_detail
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
AFTER UPDATE ON post
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
AFTER UPDATE ON comment
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
AFTER UPDATE ON vote
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();