CREATE TABLE IF NOT EXISTS "category"
(
    id           uuid PRIMARY KEY,
    created_at   timestamp(6) NOT NULL,
    updated_at   timestamp(6) NOT NULL,
    slug         varchar(120) NOT NULL UNIQUE
);

ALTER TABLE "category" ADD COLUMN IF NOT EXISTS "icon_url" varchar(255);
