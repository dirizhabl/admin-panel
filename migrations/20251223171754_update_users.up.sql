ALTER TABLE users ADD first_name varchar null CHECK (first_name <> '');
ALTER TABLE users ADD last_name varchar null CHECK (last_name <> '');
ALTER TABLE users ADD age smallint null CHECK (age > 0);