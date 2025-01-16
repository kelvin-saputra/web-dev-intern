CREATE TABLE comments (
    user_id INTEGER NOT NULL,
    id SERIAL PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL 
)