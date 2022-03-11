-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: db_mysql_reseller_v1
-- Generation Time: Mar 06, 2022 at 02:25 AM
-- Server version: 8.0.24
-- PHP Version: 7.4.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_c45_penjualan`
--

-- --------------------------------------------------------

--
-- Table structure for table `casbin_rule`
--

CREATE TABLE `casbin_rule` (
  `p_type` varchar(32) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `casbin_rule`
--

INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
('p', 'manager', '/api/manager', 'read', '', '', ''),
('p', 'manager', '/api/manager', 'write', '', '', ''),
('p', 'admin', '/api/users', 'read', '', '', ''),
('p', 'admin', '/api/users', 'write', '', '', ''),
('g', '33f1b115-b0d7-4582-a863-cd5a9d28014b', 'admin', '', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `username` text NOT NULL,
  `password` text NOT NULL,
  `role` text NOT NULL,
  `nama_lengkap` text NOT NULL,
  `no_telp` text NOT NULL,
  `email` text NOT NULL,
  `created_at` text NOT NULL,
  `updated_at` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `role`, `nama_lengkap`, `no_telp`, `email`, `created_at`, `updated_at`) VALUES
('33f1b115-b0d7-4582-a863-cd5a9d28014b', 'rhein', '$2a$10$JEGr.ClHzClxGZJC2fPFGOxCW.hJZ.2ePLvf.KGvx3Q2dsjcorf4u', 'admin', 'rhein ali', '082219193211', 'admin@admin.com', '03-06-2022 10:24:08', '03-06-2022 10:24:08');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `casbin_rule`
--
ALTER TABLE `casbin_rule`
  ADD KEY `idx_casbin_rule` (`p_type`,`v0`,`v1`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
