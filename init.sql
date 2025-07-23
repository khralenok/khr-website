CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  image_url TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO posts(content, image_url) VALUES
($$ Last Friday, I nearly dropped a project I hadn't even started — just because it felt too complex.

Then I thought: with that mindset, I'd never build anything cooler than a calculator.

So I broke it into doable parts — and today, I finally started working on the AllWallets database.

It's the first step toward building the full app. $$, 
'/img/illustration-1.jpg'),
($$ Midweek AllWallets API update:

🟢 Refined the auth logic from my previous API — this time, I didn't miss a uniqueness constraint for usernames 😜.

🟢 Added wallet creation logic. Now, users can create wallets and become admins of them. It's a basic setup for features coming soon.

🟢 Built a profile endpoint that returns user info along with their wallets. It was my first time joining two tables in a single query — learned a bit more about SQL.

Hope your week is going great — see ya! $$,
 '/img/illustration-2.jpg'),
( $$ The first week of AllWallets API development went great! ✅ I added logic for creating wallets — and sharing them between users.

This wallet sharing feature makes my API stand out — most personal finance trackers are… well, personal. But what if finances are shared — between family or partners?

There's a gap between simple apps and pro accounting tools — and it's one I want to explore.

I'm just learning for now — not testing demand yet. But it's still worth solving a real problem. 🧩 $$,
 '/img/illustration-3.jpg'),
( $$ And of course, right after my previous post, I spent two hours debugging — all because of a wrong key in the request JSON… 🫩$$,
 '');