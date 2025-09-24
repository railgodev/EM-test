-- +goose Down
DELETE FROM subscriptions
WHERE service_name IN ('Netflix', 'Spotify', 'Amazon Prime', 'Disney+'),
    ('Spotify', 999, gen_random_uuid(), '2025-03-15', NULL),
  ('Amazon Prime', 1200, gen_random_uuid(), '2025-05-01', '2026-05-01'),
  ('Disney+', 800, gen_random_uuid(), '2025-06-10', NULL);
