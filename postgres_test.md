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
    SELECT u.name, o.amount, o.created_at
    FROM users u
    JOIN orders o ON u.id = o.user_id;
    ```
