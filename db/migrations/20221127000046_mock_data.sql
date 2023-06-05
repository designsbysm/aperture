-- migrate:up
DO $$
DECLARE
  user_id uuid = '32cc62f0-292e-4349-9919-db2088ff22da';
BEGIN
  INSERT INTO authentication.users (id, updated_at, first_name, last_name, mobile, email, PASSWORD)
  VALUES(
    user_id,
    now(),
    'Scott',
    'Matthews',
    '+16025381318',
    'scott@designsbysm.com',
    '$2a$10$fUkJsgycKsIN.VZgoR0zWOhL3QNEiTSrT3bNeoVYl6RHk5H3mDgF.'
  );

  INSERT INTO authentication.roles (updated_at, user_id, ROLE)
  VALUES(
    now(),
    user_id,
    'admin'
  );

  INSERT INTO authentication.refresh_tokens (updated_at, user_id, refresh_token)
  VALUES(
    now(),
    user_id,
    '5ae3bba3-b5ab-4b94-a768-9a61fc4a4c50'
  );
END $$;

-- migrate:down
TRUNCATE authentication.users CASCADE;
