-- Drop constraints
ALTER TABLE users DROP CONSTRAINT IF EXISTS chk_hospital_id_for_role;

-- Drop indexes
DROP INDEX IF EXISTS idx_users_hospital_id;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_email;

-- Drop table
DROP TABLE IF EXISTS users;