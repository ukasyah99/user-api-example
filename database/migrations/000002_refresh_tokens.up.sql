CREATE TABLE `refresh_tokens` (
   `id` CHAR(32) PRIMARY KEY,
   `user_id` BIGINT NOT NULL,
   `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   `valid_until` TIMESTAMP NOT NULL,
   FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
