# Postgres Test

1. Soal 1:

   ```bash
    SELECT u.name, SUM(o.amount) AS total_spent
    FROM users u
    JOIN orders o ON u.id = o.user_id
    WHERE o.created_at >= '2022-01-01'
    GROUP BY u.name
    HAVING SUM(o.amount) >= 1000;
    ```

2. Soal 2:

   ```bash
    CREATE EXTENSION IF NOT EXISTS postgres_fdw;

    CREATE SERVER second_db FOREIGN DATA WRAPPER postgres_fdw OPTIONS (dbname 'second', host 'localhost', port '5432');

    --make sure db second has the corelated user
    CREATE USER MAPPING FOR current_user SERVER second_db OPTIONS (user 'username', password 'password');

    CREATE FOREIGN TABLE orders_remote (
      id SERIAL PRIMARY KEY,
      user_id INT,
      amount NUMERIC(10,2),
      created_at TIMESTAMP
    ) SERVER second_db OPTIONS (schema_name 'public', table_name 'orders');

    SELECT u.name, o.amount, o.created_at
    FROM users u
    JOIN orders_remote o ON u.id = o.user_id;
    ```