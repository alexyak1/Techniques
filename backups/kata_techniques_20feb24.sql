-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Хост: sql.freedb.tech
-- Время создания: Фев 20 2024 г., 15:03
-- Версия сервера: 8.0.36-0ubuntu0.22.04.1
-- Версия PHP: 8.2.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `freedb_techniques`
--

-- --------------------------------------------------------

--
-- Структура таблицы `kata_techniques`
--

CREATE TABLE `kata_techniques` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `kata_name` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `type` varchar(255) NOT NULL,
  `image_id` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `kata_techniques`
--

INSERT INTO `kata_techniques` (`id`, `name`, `kata_name`, `image_url`, `type`, `image_id`) VALUES
(1, 'Uki-otoshi', 'nage-no-kata', '', 'Te-waza', ''),
(2, 'Uki-goshi', 'nage-no-kata', '', 'Te-waza', ''),
(3, 'Kata-guruma', 'nage-no-kata', '', 'Te-waza', ''),
(4, 'Uki-goshi', 'nage-no-kata', '', 'Koshi-waza', ''),
(5, 'Harai-goshi', 'nage-no-kata', '', 'Koshi-waza', ''),
(6, 'Tsurikomi-goshi', 'nage-no-kata', '', 'Koshi-waza', ''),
(7, 'Okuri-ashi-harai', 'nage-no-kata', '', 'Ashi-Waza', ''),
(8, 'Sasae-tsurikomi-ashi', 'nage-no-kata', '', 'Ashi-Waza', ''),
(9, 'Uchi-mata', 'nage-no-kata', '', 'Ashi-Waza', ''),
(10, 'Tomoe-nage', 'nage-no-kata', '', 'Masutemi-Waza', ''),
(11, 'Ura-nage', 'nage-no-kata', '', 'Masutemi-Waza', ''),
(12, 'Sumi-gaeshi', 'nage-no-kata', '', 'Masutemi-Waza', ''),
(13, 'Yoko-gake', 'nage-no-kata', '', 'Yoko-stemi-Waza', ''),
(14, 'Yoko-guruma', 'nage-no-kata', '', 'Yoko-stemi-Waza', ''),
(16, 'Uki-waza', 'nage-no-kata', '', 'Yoko-stemi-Waza', '');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `kata_techniques`
--
ALTER TABLE `kata_techniques`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_2` (`id`),
  ADD KEY `id` (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `kata_techniques`
--
ALTER TABLE `kata_techniques`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
