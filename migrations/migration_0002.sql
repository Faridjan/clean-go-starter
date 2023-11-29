CREATE TABLE IF NOT EXISTS "page"
(
    id           uuid PRIMARY KEY,
    created_at   timestamp(6) NOT NULL,
    updated_at   timestamp(6) NOT NULL,
    category_id  uuid NOT NULL,
    title        varchar(250) NOT NULL,
    text         text NOT NULL,

    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "page_title_idx" ON page (text);