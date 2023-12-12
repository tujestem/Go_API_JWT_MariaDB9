-- phpMyAdmin SQL Dump
-- version 4.5.4.1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Czas generowania: 12 Gru 2023, 11:29
-- Version MYSQL: 5.7.11
-- Version PHP: 7.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Baza danych: `test_db`
--

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `tab1`
--

CREATE TABLE `tab1` (
  `NAME` text COLLATE utf8mb4_bin NOT NULL,
  `SURNAME` text COLLATE utf8mb4_bin NOT NULL,
  `AGE` text COLLATE utf8mb4_bin NOT NULL,
  `SEX` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Zrzut danych tabeli `tab1`
--

INSERT INTO `tab1` (`NAME`, `SURNAME`, `AGE`, `SEX`) VALUES
('MIKE', 'STOMPER', '55', 'M'),
('DIANA', 'STRUMBERG', '55', 'F'),
('TestUser', '20231212122651', '30', 'M');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
