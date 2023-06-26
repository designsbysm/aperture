-- migrate:up
DO $$
DECLARE
  user_id uuid = '0188f677-d7e4-7518-9a29-a6697234d61c';
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

  INSERT INTO authentication.refresh_tokens (id, updated_at, user_id)
  VALUES(
    '0188f678-4f36-7518-a307-fbfb4ee5188a',
    now(),
    user_id
  );
END $$;

-- migrate:down
TRUNCATE authentication.users CASCADE;
