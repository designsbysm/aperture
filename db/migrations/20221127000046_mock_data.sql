-- migrate:up
DO $$
DECLARE
  user_id uuid = '32cc62f0-292e-4349-9919-db2088ff22da';
BEGIN
  INSERT INTO authentication.users (id, updated_at, first_name, last_name, email, PASSWORD)
  VALUES(
    user_id,
    now(),
    'Scott',
    'Matthews',
    'scott@designsbysm.com',
    '$2a$10$fUkJsgycKsIN.VZgoR0zWOhL3QNEiTSrT3bNeoVYl6RHk5H3mDgF.'
  );

  INSERT INTO authentication.roles (updated_at, user_id, ROLE)
  VALUES(
    now(),
    user_id,
    'admin'
  );

  INSERT INTO authentication.refresh_tokens (id, updated_at, user_id)
  VALUES(
    '5ae3bba3-b5ab-4b94-a768-9a61fc4a4c50',
    now(),
    user_id
  );
END $$;

-- migrate:down
TRUNCATE authentication.users CASCADE;
