SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `content_category`;
CREATE TABLE `content_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `slug` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `short_description` varchar(1024) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `content_category` (`id`, `slug`, `name`, `short_description`) VALUES
(1,	'go',	'Go',	'Resource for learning Go language'),
(2,	'php',	'PHP',	'Resource for learning PHP language'),
(3,	'python',	'Python',	'Resource for learning Phyton language'),
(4,	'rust',	'Rust',	'Resource for learning Rust language'),
(5,	'nodejs',	'Node.js',	'Resource for learning Node.js'),
(6,	'react',	'React',	'Resource for learning React'),
(7,	'erlang',	'Erlang',	'Resource for learning Erlang language'),
(8,	'dotnetcore',	'.NET Core',	'Resource for learning .NET Core'),
(9,	'ruby',	'Ruby',	'Resource for learning Ruby language'),
(10,'haskell',	'Haskell',	'Resource for learning Haskell language'),
(11,'scala',	'Scala',	'Resource for learning Scala language'),
(12,'symfony',	'Symfony',	'Resource for learning Symfony framework');

