-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 03, 2024 at 03:21 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

CREATE DATABASE IF NOT EXISTS chores_mgt;

USE chores_mgt;


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `chores_mgt`
--

-- --------------------------------------------------------

--
-- Table structure for table `Assigned_people`
--

CREATE TABLE `Assigned_people` (
  `pid` int(11) NOT NULL,
  `assignmentid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Assigned_people`
--

INSERT INTO `Assigned_people` (`pid`, `assignmentid`) VALUES
(10, 53),
(10, 54),
(10, 55),
(10, 56),
(10, 57),
(10, 58),
(10, 59);

-- --------------------------------------------------------

--
-- Table structure for table `Assignment`
--

CREATE TABLE `Assignment` (
  `assignmentid` int(11) NOT NULL,
  `cid` int(11) NOT NULL,
  `sid` int(11) NOT NULL,
  `date_assign` date NOT NULL,
  `date_due` date NOT NULL,
  `last_updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `date_completed` date DEFAULT NULL,
  `img` varchar(100) DEFAULT NULL,
  `who_assigned` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Assignment`
--

INSERT INTO `Assignment` (`assignmentid`, `cid`, `sid`, `date_assign`, `date_due`, `last_updated`, `date_completed`, `img`, `who_assigned`) VALUES
(53, 12, 2, '2024-03-29', '2024-03-30', '2024-03-29 21:05:50', NULL, NULL, 11),
(54, 13, 1, '2024-03-29', '2024-03-30', '2024-03-29 17:02:51', NULL, NULL, 11),
(55, 14, 4, '2024-03-29', '2024-03-31', '2024-03-29 21:05:30', NULL, NULL, 11),
(56, 15, 1, '2024-03-29', '2024-04-05', '2024-03-29 17:03:02', NULL, NULL, 11),
(57, 16, 3, '2024-03-29', '2024-04-03', '2024-03-29 21:05:39', '2024-03-29', NULL, 11),
(58, 17, 3, '2024-03-29', '2024-04-04', '2024-03-29 21:05:38', '2024-03-29', NULL, 11),
(59, 22, 2, '2024-03-29', '2024-04-05', '2024-03-29 21:05:53', NULL, NULL, 11);

-- --------------------------------------------------------

--
-- Table structure for table `Chores`
--

CREATE TABLE `Chores` (
  `cid` int(11) NOT NULL,
  `chorename` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Chores`
--

INSERT INTO `Chores` (`cid`, `chorename`) VALUES
(12, 'Feed the cats'),
(13, 'Prepare breakfast'),
(14, 'Clean your room'),
(15, 'Make your bed'),
(16, 'Go buy new slippers'),
(17, 'Go buy new shoes'),
(22, 'Go and shower');

-- --------------------------------------------------------

--
-- Table structure for table `Family_name`
--

CREATE TABLE `Family_name` (
  `fid` int(11) NOT NULL,
  `fam_name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Family_name`
--

INSERT INTO `Family_name` (`fid`, `fam_name`) VALUES
(1, 'Father'),
(2, 'Mother'),
(3, 'Son'),
(4, 'Daughter'),
(5, 'Sister'),
(6, 'Brother');

-- --------------------------------------------------------

--
-- Table structure for table `People`
--

CREATE TABLE `People` (
  `pid` int(11) NOT NULL,
  `rid` int(11) NOT NULL,
  `fid` int(11) NOT NULL,
  `fname` varchar(50) NOT NULL,
  `lname` varchar(50) NOT NULL,
  `gender` int(11) NOT NULL,
  `dob` date NOT NULL,
  `tel` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `passwd` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `People`
--

INSERT INTO `People` (`pid`, `rid`, `fid`, `fname`, `lname`, `gender`, `dob`, `tel`, `email`, `passwd`) VALUES
(10, 3, 6, 'Remedios', 'Peck', 1, '2001-02-18', '+1 (461) 772-5951', 'someone@mail.com', '$2y$10$/UcpQ08nEQUzU3eMXUGwc.eqOMP9us3p3TgCrLJlrnYZOXxrcyQZ.'),
(11, 2, 2, 'Martha', 'Patrick', 1, '1996-08-21', '+1 (539) 821-7792', 'another@mail.com', '$2y$10$V7GdpBbsl3EYDdxVbqLeH.S3craU7xW.1aK.D6F8qUjevSXZyaNLO'),
(12, 2, 2, 'John', 'Doe', 0, '1990-01-01', '1234567890', 'email@mail.com', '$2a$10$ImNpkgSZjJQ1ykooDDt5RemGKKMq2Z82mlReETSUQEhR5Ek2LJTIu'),
(13, 3, 4, 'Ulric', 'Sanchez', 1, '1996-04-06', '+1 (965) 648-4829', 'she@mail.com', '$2y$10$JvLRj8TQVOSOvbzpphY9TeNUacQMd3RGozlEk0ggbx5RC3pcTQ4ay');

-- --------------------------------------------------------

--
-- Table structure for table `Role`
--

CREATE TABLE `Role` (
  `rid` int(11) NOT NULL,
  `rname` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Role`
--

INSERT INTO `Role` (`rid`, `rname`) VALUES
(1, 'superadmin'),
(2, 'admin'),
(3, 'standard');

-- --------------------------------------------------------

--
-- Table structure for table `Status`
--

CREATE TABLE `Status` (
  `sid` int(11) NOT NULL,
  `sname` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `Status`
--

INSERT INTO `Status` (`sid`, `sname`) VALUES
(1, 'Assigned'),
(2, 'InProgress'),
(3, 'Completed'),
(4, 'Incomplete');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Assigned_people`
--
ALTER TABLE `Assigned_people`
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `pid` (`pid`);

--
-- Indexes for table `Assignment`
--
ALTER TABLE `Assignment`
  ADD PRIMARY KEY (`assignmentid`),
  ADD KEY `cid` (`cid`),
  ADD KEY `sid` (`sid`),
  ADD KEY `who_assigned` (`who_assigned`);

--
-- Indexes for table `Chores`
--
ALTER TABLE `Chores`
  ADD PRIMARY KEY (`cid`);

--
-- Indexes for table `Family_name`
--
ALTER TABLE `Family_name`
  ADD PRIMARY KEY (`fid`);

--
-- Indexes for table `People`
--
ALTER TABLE `People`
  ADD PRIMARY KEY (`pid`),
  ADD KEY `rid` (`rid`),
  ADD KEY `fid` (`fid`);

--
-- Indexes for table `Role`
--
ALTER TABLE `Role`
  ADD PRIMARY KEY (`rid`);

--
-- Indexes for table `Status`
--
ALTER TABLE `Status`
  ADD PRIMARY KEY (`sid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Assignment`
--
ALTER TABLE `Assignment`
  MODIFY `assignmentid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=60;

--
-- AUTO_INCREMENT for table `Chores`
--
ALTER TABLE `Chores`
  MODIFY `cid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `Family_name`
--
ALTER TABLE `Family_name`
  MODIFY `fid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `People`
--
ALTER TABLE `People`
  MODIFY `pid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `Role`
--
ALTER TABLE `Role`
  MODIFY `rid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `Status`
--
ALTER TABLE `Status`
  MODIFY `sid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Assigned_people`
--
ALTER TABLE `Assigned_people`
  ADD CONSTRAINT `assigned_people_ibfk_1` FOREIGN KEY (`assignmentid`) REFERENCES `Assignment` (`assignmentid`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `assigned_people_ibfk_2` FOREIGN KEY (`pid`) REFERENCES `People` (`pid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Assignment`
--
ALTER TABLE `Assignment`
  ADD CONSTRAINT `assignment_ibfk_1` FOREIGN KEY (`cid`) REFERENCES `Chores` (`cid`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `assignment_ibfk_2` FOREIGN KEY (`sid`) REFERENCES `Status` (`sid`) ON DELETE NO ACTION ON UPDATE CASCADE,
  ADD CONSTRAINT `assignment_ibfk_3` FOREIGN KEY (`who_assigned`) REFERENCES `People` (`pid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `People`
--
ALTER TABLE `People`
  ADD CONSTRAINT `people_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `Role` (`rid`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `people_ibfk_2` FOREIGN KEY (`fid`) REFERENCES `Family_name` (`fid`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
