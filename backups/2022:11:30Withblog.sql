-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: sql.freedb.tech
-- Generation Time: Nov 30, 2022 at 02:51 PM
-- Server version: 8.0.28
-- PHP Version: 8.0.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `freedb_techniques`
--

-- --------------------------------------------------------

--
-- Table structure for table `blog_items`
--

CREATE TABLE `blog_items` (
  `id` int NOT NULL,
  `title` text NOT NULL,
  `date` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `blog_pool_answers`
--

CREATE TABLE `blog_pool_answers` (
  `id` int NOT NULL,
  `blog_item_id` int NOT NULL,
  `name` text NOT NULL,
  `votes` int NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `techniques`
--

CREATE TABLE `techniques` (
  `id` int NOT NULL,
  `name` varchar(128) NOT NULL,
  `belt` varchar(255) NOT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `image_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `techniques`
--

INSERT INTO `techniques` (`id`, `name`, `belt`, `image_url`, `type`, `image_id`) VALUES
(1, 'O-soto-otoshi', 'yellow', 'https://drive.google.com/file/d/1qsICsg1ZkiVHVCq7S5N9GLvSfQhv6lqm/view?usp=sharing', 'Nage-Waza', '1qsICsg1ZkiVHVCq7S5N9GLvSfQhv6lqm'),
(2, 'O-uchi-gari', 'yellow', 'https://drive.google.com/file/d/1SNbow_JB3BunQl5sK8RZZuMakFXJWlvu/view?usp=sharing', 'Nage-Waza', '1SNbow_JB3BunQl5sK8RZZuMakFXJWlvu'),
(3, 'O-goshi', 'yellow', 'https://drive.google.com/file/d/1qWipY1hFzEUQjZc_TL8z3ac0ps1GZJtf/view?usp=sharing', 'Nage-Waza', '1qWipY1hFzEUQjZc_TL8z3ac0ps1GZJtf'),
(4, 'Mune-gatame', 'yellow', 'https://drive.google.com/file/d/1pEKJv8t369Qn5mRsxuiWcQ9fu3eiK2Yz/view?usp=sharing', 'Osaekomi-Waza', '1pEKJv8t369Qn5mRsxuiWcQ9fu3eiK2Yz'),
(9, 'Kuzure-kesa-gatame', 'yellow', 'https://drive.google.com/file/d/10yZnW5XG2a6P8c5A90w-15PIhSalSgrZ/view?usp=sharing', 'Osaekomi-Waza', '10yZnW5XG2a6P8c5A90w-15PIhSalSgrZ'),
(10, 'Ko-soto-gari', 'yellow', 'https://drive.google.com/file/d/1ISq8P2jNvbHIHYlwynkPiq3r89sop0iA/view?usp=sharing', 'Nage-Waza', '1ISq8P2jNvbHIHYlwynkPiq3r89sop0iA'),
(11, 'Ko-uchi-gari', 'yellow', 'https://drive.google.com/file/d/15rPi-D3X6uCwBazZAi-9NneVjgV91MMz/view?usp=sharing', 'Nage-Waza', '15rPi-D3X6uCwBazZAi-9NneVjgV91MMz'),
(12, 'Hiza-guruma', 'yellow', 'https://drive.google.com/file/d/16jwcEta9nsCRQIZqrgxGtQam0mX31EoD/view?usp=sharing', 'Nage-Waza', '16jwcEta9nsCRQIZqrgxGtQam0mX31EoD'),
(13, 'Eri-seoi-nage', 'yellow', 'https://drive.google.com/file/d/1FGhdWu8itdWb6GHR99ac9R17RXoLhCJ6/view?usp=sharing', 'Nage-Waza', '1FGhdWu8itdWb6GHR99ac9R17RXoLhCJ6'),
(14, 'Koshi-guruma', 'yellow', 'https://drive.google.com/file/d/1kEiVAg6RXvnPEDR8TkqEn7UXZZjULY5F/view?usp=sharing', 'Nage-Waza', '1kEiVAg6RXvnPEDR8TkqEn7UXZZjULY5F'),
(15, 'Kami-shiho-gatame', 'yellow', 'https://drive.google.com/file/d/1uNEryIWSpkX8BBKDTAgQyFkgCqyX_7sO/view?usp=sharing', 'Osaekomi-Waza', '1uNEryIWSpkX8BBKDTAgQyFkgCqyX_7sO'),
(16, 'Tate-shiho-gatame', 'yellow', 'https://drive.google.com/file/d/1vOA1HejU7HMf3_iBMYre4gaVTp6W6eW4/view?usp=sharing', 'Osaekomi-Waza', '1vOA1HejU7HMf3_iBMYre4gaVTp6W6eW4'),
(17, 'De-ashi-harai', 'orange', 'https://drive.google.com/file/d/1DBMv0FlGz929TKMCaFTR-eiOavaGM5Rb/view?usp=sharing', 'Nage-Waza', '1DBMv0FlGz929TKMCaFTR-eiOavaGM5Rb'),
(18, 'Ko-soto-gake', 'orange', 'https://drive.google.com/file/d/1LThq0v77RYY4MMzp4K-lR0k3YkiA8rqO/view?usp=sharing', 'Nage-Waza', '1LThq0v77RYY4MMzp4K-lR0k3YkiA8rqO'),
(19, 'Ippon-seoi-nage', 'orange', 'https://drive.google.com/file/d/1TvnUyAXTxeunJF1XflN_yW_dDl2fwRJB/view?usp=sharing', 'Nage-Waza', '1TvnUyAXTxeunJF1XflN_yW_dDl2fwRJB'),
(20, 'Morote-seoi-nage', 'orange', 'https://drive.google.com/file/d/1ufSl4n22UXMNQkSqBxbudvHJZLDkwbTr/view?usp=sharing', 'Nage-Waza', '1ufSl4n22UXMNQkSqBxbudvHJZLDkwbTr'),
(21, 'Tsuri-komi-goshi', 'orange', 'https://drive.google.com/file/d/1YdF5tioCh0Riz3vPiAkqy7PtdMzgt5Wr/view?usp=sharing', 'Nage-Waza', '1YdF5tioCh0Riz3vPiAkqy7PtdMzgt5Wr'),
(22, 'Kubi-nage', 'orange', 'https://drive.google.com/file/d/10v5XIdURJXmAFAQC-Mg65RbsFfHPKLWT/view?usp=sharing', 'Nage-Waza', '10v5XIdURJXmAFAQC-Mg65RbsFfHPKLWT'),
(23, 'Ushiro-kesa-gatame', 'orange', 'https://drive.google.com/file/d/16tJ6byvyVvTYA_puG5yVrI9BOT1_jXV9/view?usp=sharing', 'Osaekomi-Waza', '16tJ6byvyVvTYA_puG5yVrI9BOT1_jXV9'),
(24, 'Kuzure-yoko-shiho-gatame', 'orange', 'https://drive.google.com/file/d/1FZp0THEJ47OBC-e8RPXPav5mLcDoBSce/view?usp=sharing', 'Osaekomi-Waza', '1FZp0THEJ47OBC-e8RPXPav5mLcDoBSce'),
(25, 'O-soto-gari', 'orange', 'https://drive.google.com/file/d/10I2O8sPkylJEFHx35z8G30WzSDgY3Yz_/view?usp=sharing', 'Nage-Waza', '10I2O8sPkylJEFHx35z8G30WzSDgY3Yz_'),
(26, 'Tsuri-goshi', 'orange', 'https://drive.google.com/file/d/1WQ7aOMpjwckzlPz3xYRsc0zfvAbp-GYT/view?usp=sharing', 'Nage-Waza', '1WQ7aOMpjwckzlPz3xYRsc0zfvAbp-GYT'),
(27, 'Tai-otoshi', 'orange', 'https://drive.google.com/file/d/1vs7G7rDe5UYdVJO2W0hgN6RM74d1Je33/view?usp=sharing', 'Nage-Waza', '1vs7G7rDe5UYdVJO2W0hgN6RM74d1Je33'),
(28, 'Uchi-mata', 'orange', 'https://drive.google.com/file/d/16XWRbPj2cj4MtNY3pIuJdzChfq2twDNF/view?usp=sharing', 'Nage-Waza', '16XWRbPj2cj4MtNY3pIuJdzChfq2twDNF'),
(29, 'Harai-goshi', 'orange', 'https://drive.google.com/file/d/1zD67DXdh8AqLsdXXuuoiTcU7z0fkn3Gz/view?usp=sharing', 'Nage-Waza', '1zD67DXdh8AqLsdXXuuoiTcU7z0fkn3Gz'),
(30, 'Yoko-shiho-gatame', 'orange', 'https://drive.google.com/file/d/1Q8dFGgVE8hRR5gsozj0L1oEuUWlL0g5C/view?usp=sharing', 'Osaekomi-Waza', '1Q8dFGgVE8hRR5gsozj0L1oEuUWlL0g5C'),
(31, 'Kesa-gatame', 'orange', 'https://drive.google.com/file/d/1MZE17XwQBuL4o7XK5FFU-2zX_AP4am-0/view?usp=sharing', 'Osaekomi-Waza', '1MZE17XwQBuL4o7XK5FFU-2zX_AP4am-0'),
(32, 'Gyaku-juji-jime', 'orange', 'https://drive.google.com/file/d/1cCYo3E84m_vOneQbuLKJcxH6wcTeaNhw/view?usp=sharing', 'Shime-Waza', '1cCYo3E84m_vOneQbuLKJcxH6wcTeaNhw'),
(33, 'Nami-juji-jime', 'orange', 'https://drive.google.com/file/d/18ri-ALP525z8eCJLkmF-ZKLtpj7J9OeI/view?usp=sharing', 'Shime-Waza', '18ri-ALP525z8eCJLkmF-ZKLtpj7J9OeI'),
(34, 'Kata-juji-jime', 'orange', 'https://drive.google.com/file/d/1Iv-GnvYlUJ1F3j0bJY-Mo9Aybk4tkFNO/view?usp=sharing', 'Shime-Waza', '1Iv-GnvYlUJ1F3j0bJY-Mo9Aybk4tkFNO'),
(35, 'Ude-garami', 'orange', 'https://drive.google.com/file/d/12fWDqXD8cyTtrhkA6PmOcV0DvzZfO7Cm/view?usp=sharing', 'Kansetsu-Waza', '12fWDqXD8cyTtrhkA6PmOcV0DvzZfO7Cm'),
(36, 'Juji-gatame', 'orange', 'https://drive.google.com/file/d/1tiu0rL79egqINJHBa315l_zV2ntpH_FW/view?usp=sharing', 'Kansetsu-Waza', '1tiu0rL79egqINJHBa315l_zV2ntpH_FW'),
(37, 'Sasae-tsuri-komi-ashi', 'green', 'https://drive.google.com/file/d/1kwT2TrvX7SwrpwmdfzDypJgvqVaeqm3t/view?usp=sharing', 'Nage-waza', '1kwT2TrvX7SwrpwmdfzDypJgvqVaeqm3t'),
(38, 'Sode-tsuri-komi-goshi', 'green', 'https://drive.google.com/file/d/1MmwKkINptunQ9EBzxtDoOfUcN1lqQYAP/view?usp=sharing', 'Nage-waza', '1MmwKkINptunQ9EBzxtDoOfUcN1lqQYAP'),
(39, 'Uki-goshi', 'green', 'https://drive.google.com/file/d/1c8WnN3dRwEhf-91ZXy6ziPfac2KEJEO1/view?usp=sharing', 'Nage-waza', '1c8WnN3dRwEhf-91ZXy6ziPfac2KEJEO1'),
(40, 'Seoi-otoshi', 'green', 'https://drive.google.com/file/d/13v1PXsibNJObwqIM2rMtKFfCTID1hmxk/view?usp=sharing', 'Nage-waza', '13v1PXsibNJObwqIM2rMtKFfCTID1hmxk'),
(41, 'Okuri-ashi-harai(barai)', 'green', 'https://drive.google.com/file/d/1ZlpIzCHfWdjMOaensrHgQwbs4Q0E1k9u/view?usp=sharing', 'Nage-waza', '1ZlpIzCHfWdjMOaensrHgQwbs4Q0E1k9u'),
(42, 'Kuzure-kami-shiho-gatame', 'green', 'https://drive.google.com/file/d/1NXleroQoEx6eCIrYhaxfp0zlZvWZifTF/view?usp=sharing', 'Osaekomi-waza', '1NXleroQoEx6eCIrYhaxfp0zlZvWZifTF'),
(43, 'Hadaka-jime', 'green', 'https://drive.google.com/file/d/1kiW8LleNqjFKSmZQQz57VK-br9e8OIan/view?usp=sharing', 'Shime-waza', '1kiW8LleNqjFKSmZQQz57VK-br9e8OIan'),
(44, 'Okuri-eri-jime', 'green', 'https://drive.google.com/file/d/1_A5eYn96gND_5VLPLjYmW26aPUWWN3-_/view?usp=sharing', 'Shime-waza', '1_A5eYn96gND_5VLPLjYmW26aPUWWN3-_'),
(45, 'Ude-gatame', 'green', 'https://drive.google.com/file/d/1gwgXUYF16Ghx_dvpDUeR9Ju-Gk9sSmx4/view?usp=sharing', 'Kansetsu-waza', '1gwgXUYF16Ghx_dvpDUeR9Ju-Gk9sSmx4'),
(47, 'Harai-tsuri-komi-ashi', 'green', 'https://drive.google.com/file/d/1P5ZWhRSRFqmzUCp26X5861vCqgzz0ppg/view?usp=sharing', 'Nage-waza', '1P5ZWhRSRFqmzUCp26X5861vCqgzz0ppg'),
(48, 'Tani-otoshi', 'green', 'https://drive.google.com/file/d/1WJrsR3-KaDJVgWKWwxo_j689blTFCuc6/view?usp=sharing', 'Nage-waza', '1WJrsR3-KaDJVgWKWwxo_j689blTFCuc6'),
(49, 'Yoko-sumi-gaeshi', 'green', 'https://drive.google.com/file/d/10u9_f3tn7agWfLKqrZiKMqX1CaHTTzxU/view?usp=sharing', 'Nage-waza', '10u9_f3tn7agWfLKqrZiKMqX1CaHTTzxU'),
(50, 'Tomoe-nage', 'green', 'https://drive.google.com/file/d/1FMmUpRW8KwsRNwhsPjVbYX4U_1vdof9y/view?usp=sharing', 'Nage-waza', '1FMmUpRW8KwsRNwhsPjVbYX4U_1vdof9y'),
(51, 'Yoko-tomoe-nage', 'green', 'https://drive.google.com/file/d/1ztOX-Qn_BawALYtQvEp0nJ8UMBB4kmxD/view?usp=sharing', 'Nage-waza', '1ztOX-Qn_BawALYtQvEp0nJ8UMBB4kmxD'),
(52, 'Kuzure-tate-shiho-gatame', 'green', 'https://drive.google.com/file/d/1GnGsbTrIgqPX2ASpOnbks2UI7SjroXo5/view?usp=sharing', 'Osaekomi-waza', '1GnGsbTrIgqPX2ASpOnbks2UI7SjroXo5'),
(53, 'Kata-ha-jime', 'green', 'https://drive.google.com/file/d/19SJDOmaOcl2OGqabnE3VwmJoAm8UzD7w/view?usp=sharing', 'Shime-waza', '19SJDOmaOcl2OGqabnE3VwmJoAm8UzD7w'),
(54, 'Ryote-jime', 'green', 'https://drive.google.com/file/d/10fL66PnoOhe9DGlJQksvwOmsVfoV2vi1/view?usp=sharing', 'Shime-waza', '10fL66PnoOhe9DGlJQksvwOmsVfoV2vi1'),
(55, 'Kesa-garami', 'green', 'https://drive.google.com/file/d/1yNkJZ_5ob9fowsUVLDvm9T9J_khJo1vE/view?usp=sharing', 'Kansetsu-waza', '1yNkJZ_5ob9fowsUVLDvm9T9J_khJo1vE'),
(56, 'Ude-hishigi-ashi-gatame', 'green', 'https://drive.google.com/file/d/1srRQB7RK1dtKcG5HPo-9T7NSlnPlK-pn/view?usp=sharing', 'Kansetsu-waza', '1srRQB7RK1dtKcG5HPo-9T7NSlnPlK-pn'),
(58, 'O-guruma', 'blue', 'https://drive.google.com/file/d/1gvkoLmlsiK1GzDSQ3aJhRkWxrMKPIFE5/view?usp=sharing', 'Nage-waza', '1gvkoLmlsiK1GzDSQ3aJhRkWxrMKPIFE5'),
(59, 'Ashi-guruma', 'blue', 'https://drive.google.com/file/d/1DOcNzKdfY4cwk8Frfc82ShdsmO6wjnn3/view?usp=sharing', 'Nage-waza', '1DOcNzKdfY4cwk8Frfc82ShdsmO6wjnn3'),
(60, 'Sumi-gaeshi', 'blue', 'https://drive.google.com/file/d/1QcyCvl5ozDIddo64Uy6JBqMyCDY0SywN/view?usp=sharing', 'Nage-waza', '1QcyCvl5ozDIddo64Uy6JBqMyCDY0SywN'),
(61, 'Uki-waza', 'blue', 'https://drive.google.com/file/d/11Ufc9pWl_a47g9wzrp7x7FpZRFqrIZCS/view?usp=sharing', 'Nage-waza', '11Ufc9pWl_a47g9wzrp7x7FpZRFqrIZCS'),
(62, 'Yoko-otoshi', 'blue', 'https://drive.google.com/file/d/1qkw-4Cy_N6GoUp4LvX1DNqP1-2RPfSmX/view?usp=sharing', 'Nage-waza', '1qkw-4Cy_N6GoUp4LvX1DNqP1-2RPfSmX'),
(63, 'Ushiro-goshi', 'blue', 'https://drive.google.com/file/d/1WnhUwUmg1ab9N8t1EWRT4C54D_VaXmHs/view?usp=sharing', 'Nage-waza', '1WnhUwUmg1ab9N8t1EWRT4C54D_VaXmHs'),
(64, 'Soto-makikomi', 'blue', 'https://drive.google.com/file/d/1F8GLo2JTCUa5ZpxYHv-30069NTh2DjbQ/view?usp=sharing', 'Nage-waza', '1F8GLo2JTCUa5ZpxYHv-30069NTh2DjbQ'),
(65, 'Makura-kesa-gatame', 'blue', 'https://drive.google.com/file/d/1MO-2R4vi6pbJ72CaLrvn4W3foOvIV1Im/view?usp=sharing', 'Osaekomi-waza', '1MO-2R4vi6pbJ72CaLrvn4W3foOvIV1Im'),
(66, 'Sankaku-gatame', 'blue', 'https://drive.google.com/file/d/1PYPiSlE-w7kWTZWiWQi5Tu0eQ70JUUS8/view?usp=sharing', 'Osaekomi-waza', '1PYPiSlE-w7kWTZWiWQi5Tu0eQ70JUUS8'),
(67, 'Kata-te-jime', 'blue', 'https://drive.google.com/file/d/1zJ4D8A4yR_AzzqM1NVP3epeEnyeyZxxE/view?usp=sharing', 'Shime-waza', '1zJ4D8A4yR_AzzqM1NVP3epeEnyeyZxxE'),
(68, 'Sankaku-jime', 'blue', 'https://drive.google.com/file/d/18rH2hpxaznjjB1d8ryi_y0PGYPvZR7Si/view?usp=sharing', 'Osaekomi-waza', '18rH2hpxaznjjB1d8ryi_y0PGYPvZR7Si'),
(69, 'Hara-gatame', 'blue', 'https://drive.google.com/file/d/1jVBszx9ztGcTarxDyDRjkfUN1rbqA67_/view?usp=sharing', 'Kansetsu-waza', '1jVBszx9ztGcTarxDyDRjkfUN1rbqA67_'),
(70, 'Waki-gatame', 'blue', 'https://drive.google.com/file/d/1ItzY1XneAhDoG8VUDNDkw_69Hr1jWmXW/view?usp=sharing', 'Kansetsu-waza', '1ItzY1XneAhDoG8VUDNDkw_69Hr1jWmXW'),
(71, 'Hane-goshi', 'blue', 'https://drive.google.com/file/d/1fLUkAI8pPUoyUbiItFdSKOHb-DlyYPoZ/view?usp=sharing', 'Nage-waza', '1fLUkAI8pPUoyUbiItFdSKOHb-DlyYPoZ'),
(72, 'Tsubame-geashi', 'blue', 'https://drive.google.com/file/d/1V0ahp6oSUVPtO52NH9cucA5s37kSUYrz/view?usp=sharing', 'Nage-waza', '1V0ahp6oSUVPtO52NH9cucA5s37kSUYrz'),
(73, 'Ura-nage', 'blue', 'https://drive.google.com/file/d/1bRFHSSzqGs55CM2xceWJPCbZje2F06bI/view?usp=sharing', 'Nage-waza', '1bRFHSSzqGs55CM2xceWJPCbZje2F06bI'),
(74, 'Yoko-guruma', 'blue', 'https://drive.google.com/file/d/1qET3l2kLlpfg6UOB1h_LF5H-XmQWINKg/view?usp=sharing', 'Nage-waza', '1qET3l2kLlpfg6UOB1h_LF5H-XmQWINKg'),
(75, 'Yoko-gake', 'blue', 'https://drive.google.com/file/d/1hUb4hOBuf-ttwfWTQpS3B4uB4AXRkzzo/view?usp=sharing', 'Nage-waza', '1hUb4hOBuf-ttwfWTQpS3B4uB4AXRkzzo'),
(76, 'Kata-gatame', 'blue', 'https://drive.google.com/file/d/1CEdHmWTfvQDncZhRCZWzF1haarxfMPxk/view?usp=sharing', 'Osaekomi-waza', '1CEdHmWTfvQDncZhRCZWzF1haarxfMPxk'),
(77, 'Uki-gatame', 'blue', 'https://drive.google.com/file/d/1OarSJklvbhU9YWpXE9RDeUKGP5O1sBU1/view?usp=sharing', 'Osaekomi-waza', '1OarSJklvbhU9YWpXE9RDeUKGP5O1sBU1'),
(78, 'Sode-guruma-jime', 'blue', 'https://drive.google.com/file/d/1bOfHH9slVYdfo1O3FtTYoxM257DjHk57/view?usp=sharing', 'Shime-waza', '1bOfHH9slVYdfo1O3FtTYoxM257DjHk57'),
(79, 'Tsukomi-jime', 'blue', 'https://drive.google.com/file/d/16yJT2JY3JTbkBkgEukBwvOxJ7FbKHKbs/view?usp=sharing', 'Shime-waza', '16yJT2JY3JTbkBkgEukBwvOxJ7FbKHKbs'),
(80, 'Kannuki-gatame', 'blue', 'https://drive.google.com/file/d/1kxgRXpPKXiU3frFbpr5y9QvWwKXdo-l4/view?usp=sharing', 'Kansetsu-waza', '1kxgRXpPKXiU3frFbpr5y9QvWwKXdo-l4'),
(81, 'Hiza-gatame', 'blue', 'https://drive.google.com/file/d/1GiV02hv39KucJXxly6pzGG57wP762Ll2/view?usp=sharing', 'Kansetsu-waza', '1GiV02hv39KucJXxly6pzGG57wP762Ll2'),
(82, 'Utsuri-goshi', 'brown', 'https://drive.google.com/file/d/1IudtIOUYl9cbUY6316j9o7Jn9Ehu3Scz/view?usp=sharing', 'Nage-waza', '1IudtIOUYl9cbUY6316j9o7Jn9Ehu3Scz'),
(83, 'O-soto-guruma', 'brown', 'https://drive.google.com/file/d/1ToCy6kbZOo1BDprVsqsTseOqtTYvE_9Z/view?usp=sharing', 'Nage-waza', '1ToCy6kbZOo1BDprVsqsTseOqtTYvE_9Z'),
(84, 'Uki-otoshi', 'brown', 'https://drive.google.com/file/d/1OZMG9avL8sMlcRuN12Mais1C8UZf6_K0/view?usp=sharing', 'Nage-waza', '1OZMG9avL8sMlcRuN12Mais1C8UZf6_K0'),
(85, 'Sumi-otoshi', 'brown', 'https://drive.google.com/file/d/1sfWRSPDm3LQ3NoL_V-UM2SL5CsTQJG_r/view?usp=sharing', 'Nage-waza', '1sfWRSPDm3LQ3NoL_V-UM2SL5CsTQJG_r'),
(86, 'Yoko-wakare', 'brown', 'https://drive.google.com/file/d/1TwLqgThhvEtqjukcyNfbEYhKqSSnZ3Ha/view?usp=sharing', 'Nage-waza', '1TwLqgThhvEtqjukcyNfbEYhKqSSnZ3Ha'),
(87, 'Kata-guruma', 'brown', 'https://drive.google.com/file/d/1LUpIwkmH63XTynb4yLtZKPjbmaYb5bpo/view?usp=sharing', 'Nage-waza', '1LUpIwkmH63XTynb4yLtZKPjbmaYb5bpo');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `blog_items`
--
ALTER TABLE `blog_items`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `blog_pool_answers`
--
ALTER TABLE `blog_pool_answers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `techniques`
--
ALTER TABLE `techniques`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `blog_items`
--
ALTER TABLE `blog_items`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `blog_pool_answers`
--
ALTER TABLE `blog_pool_answers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `techniques`
--
ALTER TABLE `techniques`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=88;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
