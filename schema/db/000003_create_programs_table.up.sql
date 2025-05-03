CREATE TABLE `programs` (
  `id` BINARY(16),
  `question` VARCHAR(255) NOT NULL,
  `answer` VARCHAR(255) NOT NULL,
  `content_id` BINARY(16),
  PRIMARY KEY (`id`),
  FOREIGN KEY (content_id) REFERENCES contents (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
