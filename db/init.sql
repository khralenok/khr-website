CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  image_url TEXT DEFAULT '',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_deleted BOOLEAN DEFAULT FALSE,
  deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO posts(content, image_url) VALUES
('My first month on Threads has wrapped up.\n
The numbers are modest — but hey, I’ve still got something to celebrate.\n
At least I’ve figured out what this channel is about:\n
This is where I share solo dev notes — with a pinch of self-irony and some hand-drawn illustrations.\n
Glad you’re here, folks. Let’s keep going!', '/img/example.webp')
('MVP = first-night Minecraft shelter.\n
No bed, no torch, dirt walls.\n
Ugly? Sure. But it does the job — keeps you alive till sunrise.')