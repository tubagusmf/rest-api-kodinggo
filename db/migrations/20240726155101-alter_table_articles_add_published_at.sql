-- +migrate Up
ALTER TABLE `articles` ADD `published_at` TIMESTAMP NULL AFTER `content`;

-- +migrate Down
ALTER TABLE `articles`
DROP `published_at`;