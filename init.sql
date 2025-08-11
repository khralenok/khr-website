CREATE TABLE users (
 id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
 username varchar(255) NOT NULL UNIQUE,
 pwd_hash char(60) NOT NULL,
 role varchar(5) DEFAULT 'user',
 created_at TIMESTAMP DEFAULT now(),
 CHECK(role IN ('user','admin'))
);

CREATE TABLE posts (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content TEXT NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_posts (
  id INT PRIMARY KEY REFERENCES posts(id),
  deleted_at TIMESTAMP DEFAULT now()
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
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE likes (
  post_id INT REFERENCES posts(id),
  user_id INT REFERENCES users(id),
  is_unliked BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (post_id, user_id)
)
