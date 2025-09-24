INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
VALUES 
  ('Netflix', 1500, '550e8400-e29b-41d4-a716-446655440000', '2025-01-01', '2026-01-01'),
  ('Spotify', 999, gen_random_uuid(), '2025-03-15', NULL),
  ('Amazon Prime', 1200, gen_random_uuid(), '2025-05-01', '2026-05-01'),
  ('Disney+', 800, gen_random_uuid(), '2025-06-10', NULL);

