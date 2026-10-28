DROP TABLE IF EXISTS cards CASCADE;
DROP TABLE IF EXISTS decks CASCADE;
DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL
);

CREATE TABLE decks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE (name, user_id)
);

CREATE TABLE cards (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  question TEXT NOT NULL UNIQUE,
  answer TEXT NOT NULL,
  deck_id UUID NOT NULL REFERENCES decks(id) ON DELETE CASCADE
);

INSERT INTO users(id, name, email, password_hash) VALUES 
('5af92237-164a-4bd5-93f2-db7ab8eb7c51', 'Den', 'den@den.com', '$2a$12$uiJ1ntPt/CejObRNibDBbuNPt5QVntKJ8RLzqU/kvkjE0xa0Izl8a');

INSERT INTO decks(name, user_id) VALUES
('Deck 1', '5af92237-164a-4bd5-93f2-db7ab8eb7c51'),
('Deck 2', '5af92237-164a-4bd5-93f2-db7ab8eb7c51'),
('Deck 3', '5af92237-164a-4bd5-93f2-db7ab8eb7c51'),
('Deck 4', '5af92237-164a-4bd5-93f2-db7ab8eb7c51'),
('Deck 5', '5af92237-164a-4bd5-93f2-db7ab8eb7c51'),
('Deck 6', '5af92237-164a-4bd5-93f2-db7ab8eb7c51');
