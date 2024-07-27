-- +migrate Up
CREATE TABLE
    `articles` (
        `id` bigint NOT NULL AUTO_INCREMENT,
        `title` varchar(255) NOT NULL,
        `content` text NOT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`)
    );

-- +migrate Down
DROP TABLE IF EXISTS articles;