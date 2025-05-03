CREATE TABLE `courses` (
  `id` BINARY(16),
  `name` VARCHAR(50) NOT NULL,
  `level` INT(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
