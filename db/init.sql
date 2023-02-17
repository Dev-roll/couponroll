DROP DATABASE IF EXISTS `couponroll`;
CREATE DATABASE `couponroll` DEFAULT CHARACTER SET utf8mb4;

USE `couponroll`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `stores`;
DROP TABLE IF EXISTS `coupons`;
DROP TABLE IF EXISTS `user_stores`;
DROP TABLE IF EXISTS `favorite_stores`;
DROP TABLE IF EXISTS `user_coupons`;
DROP TABLE IF EXISTS `user_points`;
DROP TABLE IF EXISTS `items`;

CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` char(128) NOT NULL DEFAULT '',
  -- `icon` char(36) DEFAULT NULL,
  `role` varchar(36) NOT NULL DEFAULT 'user',
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `stores` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  `is_public` tinyint(1) NOT NULL DEFAULT 1,
  -- `icon` char(36) DEFAULT NULL,
  `creator_id` char(36) NOT NULL,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`)
)

CREATE TABLE `coupons` (
  `id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  `creator_id` char(36) NOT NULL,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `expires_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
)

CREATE TABLE `user_stores` (
  `user_id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `joined_at` datetime(6) NOT NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`, `store_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
)

CREATE TABLE `favorite_stores` (
  `user_id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`, `store_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
)

CREATE TABLE `user_coupons` (
  `user_id` char(36) NOT NULL,
  `coupon_id` char(36) NOT NULL,
  `is_used` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`, `coupon_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`coupon_id`) REFERENCES `coupons` (`id`)
)

CREATE TABLE `user_points` (
  `user_id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `points` int NOT NULL DEFAULT 0,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  `expires_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`user_id`, `store_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
)

CREATE TABLE `items` (
  `id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  `price` int NOT NULL DEFAULT 0,
  `stock` int NOT NULL DEFAULT 0,
  `is_public` tinyint(1) NOT NULL DEFAULT 1,
  `creator_id` char(36) NOT NULL,
  `created_at` datetime(6) NOT NULL DEFAULT NULL,
  `updated_at` datetime(6) NOT NULL DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`)
)