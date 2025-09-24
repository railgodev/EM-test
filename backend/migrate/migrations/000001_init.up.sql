-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name    TEXT NOT NULL,
    price           INT NOT NULL,
    user_id         UUID NOT NULL,
    start_date      DATE NOT NULL,
    end_date        DATE, 
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT price_nonnegative CHECK (price >= 0),
    CONSTRAINT end_date_valid CHECK (end_date IS NULL OR end_date > start_date)
);

-- index on service_name column for faster lookup
CREATE INDEX IF NOT EXISTS idx_subscriptions_service_name
  ON subscriptions(service_name);

-- index on user_id column for faster lookup by user
CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id
  ON subscriptions(user_id);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER trg_set_updated_at
BEFORE UPDATE ON subscriptions
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
