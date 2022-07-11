-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 26, 2022 at 01:45 PM
-- Server version: 8.0.13-4
-- PHP Version: 7.2.24-0ubuntu0.18.04.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `hzhf7kfMUy`
--

-- --------------------------------------------------------

--
-- Table structure for table `techniques`
--

CREATE TABLE `techniques` (
  `id` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `belt` varchar(255) NOT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `techniques`
--

INSERT INTO `techniques` (`id`, `name`, `belt`, `image_url`, `type`) VALUES
(1, 'O-soto-otoshi', 'yellow', 'https://drive.google.com/file/d/1p3JuGkd533lBv4lrzzUD4BibXy4Jjinz/view?usp=sharing', 'Nage-Waza'),
(2, 'O-uchi-gari', 'yellow', 'https://drive.google.com/file/d/1APzyii2ibDU82E_OZXilW3fl7_KZMFSr/view?usp=sharing', 'Nage-Waza'),
(3, 'O-goshi', 'yellow', 'https://drive.google.com/file/d/1V_CtDZUFY3T8Q3Zt4tKqXpBc9YV4G-j-/view?usp=sharing', 'Nage-Waza'),
(4, 'Mune-gatame', 'yellow', 'https://drive.google.com/file/d/1zoDxMR04Zi6RudTyegS_m39GHNbt49kZ/view?usp=sharing', 'Osaekomi-Waza'),
(9, 'Kuzure-kesa-gatame', 'yellow', 'https://drive.google.com/file/d/1fcfVAvjw6Ec8J3qnPvPpGVDfck2j5fn8/view?usp=sharing', 'Osaekomi-Waza'),
(10, 'Ko-soto-gari', 'yellow', 'https://drive.google.com/file/d/1zN2obD93q6xNl2bCpmDYXitb5LwysA7q/view?usp=sharing', 'Nage-Waza'),
(11, 'Ko-uchi-gari', 'yellow', 'https://drive.google.com/file/d/13Au37XbZXHxEcvqEbf_yFDhNGN9oCgiD/view?usp=sharing', 'Nage-Waza'),
(12, 'Hiza-guruma', 'yellow', 'https://drive.google.com/file/d/1_TcaJVLmIha2-q8uh3wa1xHoLxHtx3Tn/view?usp=sharing', 'Nage-Waza'),
(13, 'Eri-seoi-nage', 'yellow', 'https://drive.google.com/file/d/1sE9ACTQch9nTm_OZN6OQKPk66ZotIA4W/view?usp=sharing', 'Nage-Waza'),
(14, 'Koshi-guruma', 'yellow', 'https://drive.google.com/file/d/1TkOP0wbnN0CWQ-r8yVUzhDGSzKlD0XEm/view?usp=sharing', 'Nage-Waza'),
(15, 'Kami-shiho-gatame', 'yellow', 'https://drive.google.com/file/d/1vypc8zfwcNwAYk3CzD2QvPGf2y_8Hdtq/view?usp=sharing', 'Osaekomi-Waza'),
(16, 'Tate-shiho-gatame', 'yellow', 'https://drive.google.com/file/d/1JOAKswZmPIH2BVMKUHWClvTKsGrQgiOT/view?usp=sharing', 'Osaekomi-Waza'),
(17, 'De-ashi-harai', 'orange', 'https://drive.google.com/file/d/1DBMv0FlGz929TKMCaFTR-eiOavaGM5Rb/view?usp=sharing', 'Nage-Waza'),
(18, 'Ko-soto-gake', 'orange', 'https://drive.google.com/file/d/1LThq0v77RYY4MMzp4K-lR0k3YkiA8rqO/view?usp=sharing', 'Nage-Waza'),
(19, 'Ippon-seoi-nage', 'orange', 'https://drive.google.com/file/d/1TvnUyAXTxeunJF1XflN_yW_dDl2fwRJB/view?usp=sharing', 'Nage-Waza'),
(20, 'Morote-seoi-nage', 'orange', 'https://drive.google.com/file/d/1ufSl4n22UXMNQkSqBxbudvHJZLDkwbTr/view?usp=sharing', 'Nage-Waza'),
(21, 'Tsuri-komi-goshi', 'orange', 'https://drive.google.com/file/d/1YdF5tioCh0Riz3vPiAkqy7PtdMzgt5Wr/view?usp=sharing', 'Nage-Waza'),
(22, 'Kubi-nage', 'orange', 'https://drive.google.com/file/d/10v5XIdURJXmAFAQC-Mg65RbsFfHPKLWT/view?usp=sharing', 'Nage-Waza'),
(23, 'Ushiro-kesa-gatame', 'orange', 'https://drive.google.com/file/d/16tJ6byvyVvTYA_puG5yVrI9BOT1_jXV9/view?usp=sharing', 'Osaekomi-Waza'),
(24, 'Kuzure-yoko-shiho-gatame', 'orange', 'https://drive.google.com/file/d/1FZp0THEJ47OBC-e8RPXPav5mLcDoBSce/view?usp=sharing', 'Osaekomi-Waza'),
(25, 'O-soto-gari', 'orange', 'https://drive.google.com/file/d/10I2O8sPkylJEFHx35z8G30WzSDgY3Yz_/view?usp=sharing', 'Nage-Waza'),
(26, 'Tsuri-goshi', 'orange', 'https://drive.google.com/file/d/1WQ7aOMpjwckzlPz3xYRsc0zfvAbp-GYT/view?usp=sharing', 'Nage-Waza'),
(27, 'Tai-otoshi', 'orange', 'https://drive.google.com/file/d/1vs7G7rDe5UYdVJO2W0hgN6RM74d1Je33/view?usp=sharing', 'Nage-Waza'),
(28, 'Uchi-mata', 'orange', 'https://drive.google.com/file/d/16XWRbPj2cj4MtNY3pIuJdzChfq2twDNF/view?usp=sharing', 'Nage-Waza'),
(29, 'Harai-goshi', 'orange', 'https://drive.google.com/file/d/1zD67DXdh8AqLsdXXuuoiTcU7z0fkn3Gz/view?usp=sharing', 'Nage-Waza'),
(30, 'Yoko-shiho-gatame', 'orange', 'https://drive.google.com/file/d/1Q8dFGgVE8hRR5gsozj0L1oEuUWlL0g5C/view?usp=sharing', 'Osaekomi-Waza'),
(31, 'Kesa-gatame', 'orange', 'https://drive.google.com/file/d/1MZE17XwQBuL4o7XK5FFU-2zX_AP4am-0/view?usp=sharing', 'Osaekomi-Waza'),
(32, 'Gyaku-juji-jime', 'orange', 'https://drive.google.com/file/d/1cCYo3E84m_vOneQbuLKJcxH6wcTeaNhw/view?usp=sharing', 'Shime-Waza'),
(33, 'Nami-juji-jime', 'orange', 'https://drive.google.com/file/d/18ri-ALP525z8eCJLkmF-ZKLtpj7J9OeI/view?usp=sharing', 'Shime-Waza'),
(34, 'Kata-juji-jime', 'orange', 'https://drive.google.com/file/d/1Iv-GnvYlUJ1F3j0bJY-Mo9Aybk4tkFNO/view?usp=sharing', 'Shime-Waza'),
(35, 'Ude-garami', 'orange', 'https://drive.google.com/file/d/12fWDqXD8cyTtrhkA6PmOcV0DvzZfO7Cm/view?usp=sharing', 'Kansetsu-Waza'),
(36, 'Juji-gatame', 'orange', 'https://drive.google.com/file/d/1tiu0rL79egqINJHBa315l_zV2ntpH_FW/view?usp=sharing', 'Kansetsu-Waza'),
(37, 'Sasae-tsuri-komi-ashi', 'green', 'https://drive.google.com/file/d/1kwT2TrvX7SwrpwmdfzDypJgvqVaeqm3t/view?usp=sharing', 'Nage-waza'),
(38, 'Sode-tsuri-komi-goshi', 'green', 'https://drive.google.com/file/d/1MmwKkINptunQ9EBzxtDoOfUcN1lqQYAP/view?usp=sharing', 'Nage-waza'),
(39, 'Uki-goshi', 'green', 'https://drive.google.com/file/d/1c8WnN3dRwEhf-91ZXy6ziPfac2KEJEO1/view?usp=sharing', 'Nage-waza'),
(40, 'Seoi-otoshi', 'green', 'https://drive.google.com/file/d/13v1PXsibNJObwqIM2rMtKFfCTID1hmxk/view?usp=sharing', 'Nage-waza'),
(41, 'Okuri-ashi-harai(barai)', 'green', 'https://drive.google.com/file/d/1ZlpIzCHfWdjMOaensrHgQwbs4Q0E1k9u/view?usp=sharing', 'Nage-waza'),
(42, 'Kuzure-kami-shiho-gatame', 'green', 'https://drive.google.com/file/d/1NXleroQoEx6eCIrYhaxfp0zlZvWZifTF/view?usp=sharing', 'Osaekomi-waza'),
(43, 'Hadaka-jime', 'green', 'https://drive.google.com/file/d/1kiW8LleNqjFKSmZQQz57VK-br9e8OIan/view?usp=sharing', 'Shime-waza'),
(44, 'Okuri-eri-jime', 'green', 'https://drive.google.com/file/d/1_A5eYn96gND_5VLPLjYmW26aPUWWN3-_/view?usp=sharing', 'Shime-waza'),
(45, 'Ude-gatame', 'green', 'https://drive.google.com/file/d/1gwgXUYF16Ghx_dvpDUeR9Ju-Gk9sSmx4/view?usp=sharing', 'Kansetsu-waza'),
(47, 'Harai-tsuri-komi-ashi', 'green', 'https://drive.google.com/file/d/1P5ZWhRSRFqmzUCp26X5861vCqgzz0ppg/view?usp=sharing', 'Nage-waza'),
(48, 'Tani-otoshi', 'green', 'https://drive.google.com/file/d/1WJrsR3-KaDJVgWKWwxo_j689blTFCuc6/view?usp=sharing', 'Nage-waza'),
(49, 'Yoko-sumi-gaeshi', 'green', 'https://drive.google.com/file/d/10u9_f3tn7agWfLKqrZiKMqX1CaHTTzxU/view?usp=sharing', 'Nage-waza'),
(50, 'Tomoe-nage', 'green', 'https://drive.google.com/file/d/1FMmUpRW8KwsRNwhsPjVbYX4U_1vdof9y/view?usp=sharing', 'Nage-waza'),
(51, 'Yoko-tomoe-nage', 'green', 'https://drive.google.com/file/d/1ztOX-Qn_BawALYtQvEp0nJ8UMBB4kmxD/view?usp=sharing', 'Nage-waza'),
(52, 'Kuzure-tate-shiho-gatame', 'green', 'https://drive.google.com/file/d/1GnGsbTrIgqPX2ASpOnbks2UI7SjroXo5/view?usp=sharing', 'Osaekomi-waza'),
(53, 'Kata-ha-jime', 'green', 'https://drive.google.com/file/d/19SJDOmaOcl2OGqabnE3VwmJoAm8UzD7w/view?usp=sharing', 'Shime-waza'),
(54, 'Ryote-jime', 'green', 'https://drive.google.com/file/d/10fL66PnoOhe9DGlJQksvwOmsVfoV2vi1/view?usp=sharing', 'Shime-waza'),
(55, 'Kesa-garami', 'green', 'https://drive.google.com/file/d/1yNkJZ_5ob9fowsUVLDvm9T9J_khJo1vE/view?usp=sharing', 'Kansetsu-waza'),
(56, 'Ude-hishigi-ashi-gatame', 'green', 'https://drive.google.com/file/d/1srRQB7RK1dtKcG5HPo-9T7NSlnPlK-pn/view?usp=sharing', 'Kansetsu-waza'),
(58, 'O-guruma', 'blue', 'https://drive.google.com/file/d/1gvkoLmlsiK1GzDSQ3aJhRkWxrMKPIFE5/view?usp=sharing', 'Nage-waza'),
(59, 'Ashi-guruma', 'blue', 'https://drive.google.com/file/d/1DOcNzKdfY4cwk8Frfc82ShdsmO6wjnn3/view?usp=sharing', 'Nage-waza'),
(60, 'Sumi-gaeshi', 'blue', 'https://drive.google.com/file/d/1QcyCvl5ozDIddo64Uy6JBqMyCDY0SywN/view?usp=sharing', 'Nage-waza'),
(61, 'Uki-waza', 'blue', 'https://drive.google.com/file/d/11Ufc9pWl_a47g9wzrp7x7FpZRFqrIZCS/view?usp=sharing', 'Nage-waza'),
(62, 'Yoko-otoshi', 'blue', 'https://drive.google.com/file/d/1qkw-4Cy_N6GoUp4LvX1DNqP1-2RPfSmX/view?usp=sharing', 'Nage-waza'),
(63, 'Ushiro-goshi', 'blue', 'https://drive.google.com/file/d/1WnhUwUmg1ab9N8t1EWRT4C54D_VaXmHs/view?usp=sharing', 'Nage-waza'),
(64, 'Soto-makikomi', 'blue', 'https://drive.google.com/file/d/1F8GLo2JTCUa5ZpxYHv-30069NTh2DjbQ/view?usp=sharing', 'Nage-waza'),
(65, 'Makura-kesa-gatame', 'blue', 'https://drive.google.com/file/d/1MO-2R4vi6pbJ72CaLrvn4W3foOvIV1Im/view?usp=sharing', 'Osaekomi-waza'),
(66, 'Sankaku-gatame', 'blue', 'https://drive.google.com/file/d/1PYPiSlE-w7kWTZWiWQi5Tu0eQ70JUUS8/view?usp=sharing', 'Osaekomi-waza'),
(67, 'Kata-te-jime', 'blue', 'https://drive.google.com/file/d/1zJ4D8A4yR_AzzqM1NVP3epeEnyeyZxxE/view?usp=sharing', 'Shime-waza'),
(68, 'Sankaku-jime', 'blue', 'https://drive.google.com/file/d/18rH2hpxaznjjB1d8ryi_y0PGYPvZR7Si/view?usp=sharing', 'Osaekomi-waza'),
(69, 'Hara-gatame', 'blue', 'https://drive.google.com/file/d/1jVBszx9ztGcTarxDyDRjkfUN1rbqA67_/view?usp=sharing', 'Kansetsu-waza'),
(70, 'Waki-gatame', 'blue', 'https://drive.google.com/file/d/1ItzY1XneAhDoG8VUDNDkw_69Hr1jWmXW/view?usp=sharing', 'Kansetsu-waza'),
(71, 'Hane-goshi', 'blue', 'https://drive.google.com/file/d/1fLUkAI8pPUoyUbiItFdSKOHb-DlyYPoZ/view?usp=sharing', 'Nage-waza'),
(72, 'Tsubame-geashi', 'blue', 'https://drive.google.com/file/d/1V0ahp6oSUVPtO52NH9cucA5s37kSUYrz/view?usp=sharing', 'Nage-waza'),
(73, 'Ura-nage', 'blue', 'https://drive.google.com/file/d/1bRFHSSzqGs55CM2xceWJPCbZje2F06bI/view?usp=sharing', 'Nage-waza'),
(74, 'Yoko-guruma', 'blue', 'https://drive.google.com/file/d/1qET3l2kLlpfg6UOB1h_LF5H-XmQWINKg/view?usp=sharing', 'Nage-waza'),
(75, 'Yoko-gake', 'blue', 'https://drive.google.com/file/d/1hUb4hOBuf-ttwfWTQpS3B4uB4AXRkzzo/view?usp=sharing', 'Nage-waza'),
(76, 'Kata-gatame', 'blue', 'https://drive.google.com/file/d/1CEdHmWTfvQDncZhRCZWzF1haarxfMPxk/view?usp=sharing', 'Osaekomi-waza'),
(77, 'Uki-gatame', 'blue', 'https://drive.google.com/file/d/1OarSJklvbhU9YWpXE9RDeUKGP5O1sBU1/view?usp=sharing', 'Osaekomi-waza'),
(78, 'Sode-guruma-jime', 'blue', 'https://drive.google.com/file/d/1bOfHH9slVYdfo1O3FtTYoxM257DjHk57/view?usp=sharing', 'Shime-waza'),
(79, 'Tsukomi-jime', 'blue', 'https://drive.google.com/file/d/16yJT2JY3JTbkBkgEukBwvOxJ7FbKHKbs/view?usp=sharing', 'Shime-waza'),
(80, 'Kannuki-gatame', 'blue', 'https://drive.google.com/file/d/1kxgRXpPKXiU3frFbpr5y9QvWwKXdo-l4/view?usp=sharing', 'Kansetsu-waza'),
(81, 'Hiza-gatame', 'blue', 'https://drive.google.com/file/d/1GiV02hv39KucJXxly6pzGG57wP762Ll2/view?usp=sharing', 'Kansetsu-waza');
(82, 'Kata-guruma', 'brown', 'https://drive.google.com/file/d/1LUpIwkmH63XTynb4yLtZKPjbmaYb5bpo/view?usp=sharing', 'Kansetsu-waza');
(83, 'Yoko-wakare', 'brown', 'https://drive.google.com/file/d/1TwLqgThhvEtqjukcyNfbEYhKqSSnZ3Ha/view?usp=sharing', 'Kansetsu-waza');
(84, 'Sumi-otoshi', 'brown', 'https://drive.google.com/file/d/1sfWRSPDm3LQ3NoL_V-UM2SL5CsTQJG_r/view?usp=sharing', 'Kansetsu-waza');
(85, 'Uki-otoshi', 'brown', 'https://drive.google.com/file/d/1OZMG9avL8sMlcRuN12Mais1C8UZf6_K0/view?usp=sharing', 'Kansetsu-waza');
(86, 'O-soto-guruma', 'brown', 'https://drive.google.com/file/d/1ToCy6kbZOo1BDprVsqsTseOqtTYvE_9Z/view?usp=sharing', 'Kansetsu-waza');
(87, 'Utsuri-gosh', 'brown', 'https://drive.google.com/file/d/1IudtIOUYl9cbUY6316j9o7Jn9Ehu3Scz/view?usp=sharing', 'Kansetsu-waza');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `techniques`
--
ALTER TABLE `techniques`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `techniques`
--
ALTER TABLE `techniques`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=87;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
