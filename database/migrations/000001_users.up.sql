CREATE TABLE `users` (
   `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
   `display_name` VARCHAR(50) NOT NULL,
   `name` VARCHAR(50) NOT NULL UNIQUE,
   `email` VARCHAR(50) NOT NULL UNIQUE,
   `password` VARCHAR(100) NOT NULL,
   `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO `users` VALUES (1, "Ukasyah", "ukasyah99", "hi.ukasyah@gmail.com", "rahasia", NOW(), NOW());
