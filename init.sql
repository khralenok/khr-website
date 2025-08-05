CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_posts (
  id INT PRIMARY KEY REFERENCES posts(id),
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  post_id INT REFERENCES posts(id),
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_comments (
  id INT PRIMARY KEY REFERENCES comments(id),
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE replies (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  comment_id INT REFERENCES comments(id),
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE deleted_replies (
  id INT PRIMARY KEY REFERENCES replies(id),
  deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE likes (
  post_id INT REFERENCES posts(id),
  user_id INT,
  is_unliked BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (post_id, user_id)
)
