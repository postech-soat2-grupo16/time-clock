CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    "name"       VARCHAR(255) NOT NULL,
    email        VARCHAR(255) UNIQUE NOT NULL,
    registration VARCHAR(255) UNIQUE NOT NULL,
    password     VARCHAR(255) NOT NULL,
    created_at   timestamptz  NULL,
    updated_at   timestamptz  NULL,
    deleted_at   timestamptz  NULL
);

CREATE TABLE IF NOT EXISTS time_clocks
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER     NOT NULL,
    clock_in   timestamptz NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_registration_users' AND tablename = 'users') THEN
            CREATE INDEX idx_registration_users ON users (registration);
        END IF;
END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_user_id_time_clocks' AND tablename = 'time_clocks') THEN
            CREATE INDEX idx_user_id_time_clocks ON time_clocks (user_id);
        END IF;
END $$;