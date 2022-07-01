CREATE DATABASE IF NOT EXISTS `go_web_server`;
USE go_web_server;

CREATE TABLE IF NOT EXISTS `countries` (
                                           `id` BINARY(16) PRIMARY KEY,
                                           `country_name` VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `provinces` (
                                           `id` BINARY(16) PRIMARY KEY,
                                           `province_name` VARCHAR(255) NOT NULL UNIQUE,
                                           `country_id` BINARY(16) NOT NULL,
                                           FOREIGN KEY (country_id) REFERENCES countries(id)
);

CREATE TABLE IF NOT EXISTS `localities` (
                                            `id` VARCHAR(255) PRIMARY KEY,
                                            `locality_name` VARCHAR(255) NOT NULL,
                                            `province_id` BINARY(16) NOT NULL,
                                            FOREIGN KEY (province_id) REFERENCES provinces(id)
);

CREATE TABLE IF NOT EXISTS `sellers` (
                                         `id` BINARY(16) PRIMARY KEY,
                                         `cid` VARCHAR(255) NOT NULL UNIQUE,
                                         `company_name` VARCHAR(255) NOT NULL,
                                         `address` VARCHAR(255) NOT NULL,
                                         `telephone` VARCHAR(255) NOT NULL,
                                         `locality_id` VARCHAR(255) NOT NULL,
                                         FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE IF NOT EXISTS `warehouses` (
                                            `id` BINARY(16) PRIMARY KEY,
                                            `address` VARCHAR(255) NOT NULL,
                                            `telephone` VARCHAR(255) NOT NULL,
                                            `warehouse_code` VARCHAR(255) NOT NULL UNIQUE,
                                            `locality_id` VARCHAR(255) NOT NULL,
                                            FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE IF NOT EXISTS `carries` (
                                         `id` BINARY(16) PRIMARY KEY,
                                         `cid` VARCHAR(255) NOT NULL UNIQUE,
                                         `company_name` VARCHAR(255) NOT NULL,
                                         `address` VARCHAR(255) NOT NULL,
                                         `telephone` VARCHAR(255) NOT NULL,
                                         `locality_id` VARCHAR(255) NOT NULL,
                                         FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE IF NOT EXISTS `product_types` (
                                               `id` BINARY(16) PRIMARY KEY,
                                               `description` VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS `sections` (
                                          `id` BINARY(16) PRIMARY KEY,
                                          `section_number` VARCHAR(255) NOT NULL UNIQUE,
                                          `current_capacity` INT NOT NULL,
                                          `current_temperature` DECIMAL(19,2) NOT NULL,
                                          `maximum_capacity` INT NOT NULL,
                                          `minimun_capacity` INT NOT NULL,
                                          `minimum_temperature` DECIMAL(19,2) NOT NULL,
                                          `product_type` DECIMAL(16) NOT NULL,
                                          FOREIGN KEY (product_type) REFERENCES product_types(id)
);

CREATE TABLE IF NOT EXISTS `products` (
                                          `id` BINARY(16) PRIMARY KEY,
                                          `description` VARCHAR(255) NOT NULL,
                                          `expiration_rate` DECIMAL(19,2) NOT NULL,
                                          `freezing_rate` DECIMAL(19,2) NOT NULL,
                                          `height` DECIMAL(19,2) NOT NULL,
                                          `length` DECIMAL(19,2) NOT NULL,
                                          `net_weight` DECIMAL(19,2) NOT NULL,
                                          `product_code` VARCHAR(255) NOT NULL UNIQUE,
                                          `recommended_freezing_temperature` DECIMAL(19,2) NOT NULL,
                                          `width` DECIMAL(19,2) NOT NULL,
                                          `product_type` BINARY(16) NOT NULL,
                                          `seller_id` BINARY(16) NOT NULL,
                                          FOREIGN KEY (product_type) REFERENCES product_types(id),
                                          FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

CREATE TABLE IF NOT EXISTS `product_batches` (
                                                 `id` BINARY(16),
                                                 `batch_number` VARCHAR(255) NOT NULL UNIQUE,
                                                 `current_quantity` INT NOT NULL,
                                                 `current_temperature` DECIMAL(19,2) NOT NULL,
                                                 `due_date` DATETIME(6),
                                                 `initial_quantity` INT NOT NULL,
                                                 `manufactoring_date`DATETIME(6) NOT NULL,
                                                 `manufactoring_hour` DATETIME(6) NOT NULL,
                                                 `minimum_temperature` DECIMAL(19,2) NOT NULL,
                                                 `product_id` BINARY(16) NOT NULL,
                                                 `section_id` BINARY(16) NOT NULL,
                                                 FOREIGN KEY (product_id) REFERENCES products(id),
                                                 FOREIGN KEY (section_id) REFERENCES sellers(id)
);

CREATE TABLE IF NOT EXISTS `product_records` (
                                                 `id` BINARY(16),
                                                 `last_update_date` DATETIME(6) NOT NULL,
                                                 `purchase_price` DECIMAL(19,2) NOT NULL,
                                                 `sale_price` DECIMAL(19,2) NOT NULL,
                                                 `product_id` BINARY(16) NOT NULL,
                                                 FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS `employees` (
                                           `id` BINARY(16) PRIMARY KEY,
                                           `id_card_number` VARCHAR(255) NOT NULL UNIQUE,
                                           `first_name` VARCHAR(255) NOT NULL,
                                           `last_name` VARCHAR(255) NOT NULL,
                                           `warehouse_id` BINARY(16) NOT NULL,
                                           FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

CREATE TABLE IF NOT EXISTS `buyers` (
                                        `id` BINARY(16) PRIMARY KEY,
                                        `id_card_number` VARCHAR(255) NOT NULL UNIQUE,
                                        `first_name` VARCHAR(255) NOT NULL,
                                        `last_name` VARCHAR(255) NOT NULL
);
