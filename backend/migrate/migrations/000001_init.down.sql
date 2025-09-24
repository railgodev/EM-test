-- +goose Down
DROP TRIGGER IF EXISTS trg_set_updated_at ON subscriptions;
DROP FUNCTION IF EXISTS set_updated_at();
DROP EXTENSION IF EXISTS pgcrypto;
DROP TABLE IF EXISTS subscriptions;
DROP INDEX IF EXISTS idx_subscriptions_service_name;
DROP INDEX IF EXISTS idx_subscriptions_user_id;