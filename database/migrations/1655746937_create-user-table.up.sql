CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

CREATE SEQUENCE users_type_seq;

CREATE TABLE users_type(
  id smallint NOT NULL DEFAULT NEXTVAL ('users_type_seq'),
  name varchar(80) DEFAULT NULL,
  created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE SEQUENCE users_seq;

CREATE TABLE users (
  id bigint NOT NULL DEFAULT NEXTVAL ('users_seq'),
  email varchar(160) NOT NULL,
  name varchar(200) NOT NULL,
  created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  type_id smallint NOT NULL,
  ldap_uuid varchar(36) NULL,
  PRIMARY KEY (id),
  CONSTRAINT users_type_fk FOREIGN KEY (type_id) REFERENCES users_type (id)
);

CREATE TRIGGER users_timestamp BEFORE UPDATE ON users 
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();