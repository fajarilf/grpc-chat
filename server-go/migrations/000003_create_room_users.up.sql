-- Join table for the many-to-many between rooms and users (a room has many
-- members, a user belongs to many rooms).
CREATE TABLE IF NOT EXISTS room_users (
    room_id     INTEGER     NOT NULL REFERENCES rooms (id) ON DELETE CASCADE,
    user_id     INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    joined_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (room_id, user_id)
);

-- Speeds up "list the rooms a given user is in".
CREATE INDEX IF NOT EXISTS idx_room_users_user_id ON room_users (user_id);
