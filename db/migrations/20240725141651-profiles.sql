-- +migrate Up
CREATE TABLE
    `profiles` (
        `id` bigint NOT NULL AUTO_INCREMENT,
        `user_id` bigint NOT NULL,
        `fisrt_name` varchar(100) NOT NULL,
        `last_name` varchar(100) NOT NULL,
        `bio` text NOT NULL,
        `image_url` varchar(255) NOT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        FOREIGN KEY (`user_id`) REFERENCES users (`id`)
    );

-- +migrate Down
drop table if exists profiles;