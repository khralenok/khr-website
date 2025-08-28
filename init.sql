CREATE TABLE users (
 id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
 email varchar(255) NOT NULL UNIQUE,
 display_name varchar(255) NOT NULL,
 pwd_hash char(60) NOT NULL,
 role varchar(5) DEFAULT 'user',
 avatar_filename varchar(255) DEFAULT 'avatar-placeholder.jpg',
 created_at TIMESTAMP DEFAULT now(),
 CHECK(role IN ('user','admin'))
);

CREATE TABLE sessions (
 id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
 user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
 token_hash bytea NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT now(),
 last_seen_at TIMESTAMP NOT NULL DEFAULT now(),
 expires_at TIMESTAMP NOT NULL,
 revoked_at TIMESTAMP,
 ip inet,
 user_agent text
);

CREATE TABLE posts (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content TEXT NOT NULL,
  attachment_type varchar(8) DEFAULT 'none',
  created_at TIMESTAMP DEFAULT now(),
  CHECK(attachment_type IN ('none','image','carousel','youtube'))
);

CREATE TABLE deleted_posts (
  id INT PRIMARY KEY REFERENCES posts(id),
  reason TEXT DEFAULT 'Deleted by admin',
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE attachment_images(
 id INT PRIMARY KEY REFERENCES posts(id),
 img_filename varchar(255) NOT NULL UNIQUE,
 created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE attachment_carousels(
  id INT PRIMARY KEY REFERENCES posts(id),
  last_element_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE attachment_youtube_vids(
  id INT PRIMARY KEY REFERENCES posts(id),
  video_id varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE comments (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content TEXT NOT NULL,
  post_id INT REFERENCES posts(id),
  commentator_id INT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_comments (
  id INT PRIMARY KEY REFERENCES comments(id),
  reason TEXT DEFAULT 'Cascade deletion with post',
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE replies (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content TEXT NOT NULL,
  comment_id INT REFERENCES comments(id),
  commentator_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_replies (
  id INT PRIMARY KEY REFERENCES replies(id),
  reason TEXT DEFAULT 'Cascade deletion with comment',
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE likes (
  post_id INT REFERENCES posts(id),
  user_id INT REFERENCES users(id),
  is_unliked BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (post_id, user_id)
)
