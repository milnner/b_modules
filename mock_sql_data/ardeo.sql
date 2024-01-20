-- phpMyAdmin SQL Dump
-- version 5.1.1deb5ubuntu1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Tempo de geração: 20-Jan-2024 às 16:07
-- Versão do servidor: 8.0.35-0ubuntu0.22.04.1
-- versão do PHP: 8.1.2-1ubuntu2.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `ardeo`
--
CREATE DATABASE IF NOT EXISTS `ardeo` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `ardeo`;

-- --------------------------------------------------------

--
-- Estrutura da tabela `answer_n_to_one`
--

CREATE TABLE `answer_n_to_one` (
  `id` int NOT NULL,
  `area_id` int NOT NULL,
  `one_question_n_answer_activity_id` int NOT NULL,
  `correctness` tinyint NOT NULL,
  `answer` blob NOT NULL,
  `activated` tinyint NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `area`
--

CREATE TABLE `area` (
  `id` int NOT NULL,
  `title` varchar(35) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `owner_id` int NOT NULL,
  `creation_datetime` datetime DEFAULT NULL,
  `activated` tinyint NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `classes`
--

CREATE TABLE `classes` (
  `id` int NOT NULL,
  `title` varchar(75) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(375) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `creation_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `creator_user_id` int NOT NULL,
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP,
  `area_id` int NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `class_see_content`
--

CREATE TABLE `class_see_content` (
  `id` int NOT NULL,
  `class_id` int NOT NULL,
  `content_id` int NOT NULL,
  `position` int DEFAULT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `class_takes_user`
--

CREATE TABLE `class_takes_user` (
  `id` int NOT NULL,
  `entry_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` int NOT NULL,
  `class_id` int NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `contents`
--

CREATE TABLE `contents` (
  `id` int NOT NULL,
  `title` varchar(45) NOT NULL,
  `description` varchar(400) NOT NULL,
  `creation_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `area_id` int NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `content_see_activity`
--

CREATE TABLE `content_see_activity` (
  `id` int NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1',
  `position` int NOT NULL,
  `area_id` int NOT NULL,
  `content_id` int NOT NULL,
  `text_activity_id` int DEFAULT NULL,
  `one_question_n_answer_activity_id` int DEFAULT NULL,
  `image_activity_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `image_activities`
--

CREATE TABLE `image_activities` (
  `id` int NOT NULL,
  `area_id` int NOT NULL,
  `title` varchar(155) NOT NULL,
  `_blob` blob NOT NULL,
  `last_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `one_question_n_answer_activities`
--

CREATE TABLE `one_question_n_answer_activities` (
  `id` int NOT NULL,
  `area_id` int NOT NULL,
  `question` blob NOT NULL,
  `last_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `text_activities`
--

CREATE TABLE `text_activities` (
  `id` int NOT NULL,
  `area_id` int NOT NULL,
  `title` varchar(155) NOT NULL,
  `_blob` blob NOT NULL,
  `last_update` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `surname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) NOT NULL,
  `entry_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `bourn_date` datetime NOT NULL,
  `sex` enum('male','female','other') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `hash` varchar(128) NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Estrutura da tabela `user_has_area_access`
--

CREATE TABLE `user_has_area_access` (
  `id` int NOT NULL,
  `permission` enum('read','write') NOT NULL,
  `entry_datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `area_id` int NOT NULL,
  `user_id` int NOT NULL,
  `activated` tinyint UNSIGNED NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `answer_n_to_one`
--
ALTER TABLE `answer_n_to_one`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ans_n_to_one_area_fk` (`area_id`),
  ADD KEY `one_question_n_answer_activity_fk` (`one_question_n_answer_activity_id`);

--
-- Índices para tabela `area`
--
ALTER TABLE `area`
  ADD PRIMARY KEY (`id`),
  ADD KEY `owner_id` (`owner_id`);

--
-- Índices para tabela `classes`
--
ALTER TABLE `classes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_title_creator_user_id` (`title`,`creator_user_id`),
  ADD KEY `id_usuario_criador` (`creator_user_id`),
  ADD KEY `area_id` (`area_id`);

--
-- Índices para tabela `class_see_content`
--
ALTER TABLE `class_see_content`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `class_id` (`class_id`,`content_id`),
  ADD KEY `content_id` (`content_id`);

--
-- Índices para tabela `class_takes_user`
--
ALTER TABLE `class_takes_user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `class_has_user_ibfk_1` (`user_id`),
  ADD KEY `class_has_user_ibfk_2` (`class_id`);

--
-- Índices para tabela `contents`
--
ALTER TABLE `contents`
  ADD PRIMARY KEY (`id`),
  ADD KEY `area_id` (`area_id`);

--
-- Índices para tabela `content_see_activity`
--
ALTER TABLE `content_see_activity`
  ADD PRIMARY KEY (`id`),
  ADD KEY `content_s_act_area_fk` (`area_id`),
  ADD KEY `content_s_act_content_fk` (`content_id`),
  ADD KEY `content_s_act_text_act_fk` (`text_activity_id`),
  ADD KEY `content_s_act_image_act_fk` (`image_activity_id`),
  ADD KEY `content_s_act_one_q_n_asw_fk` (`one_question_n_answer_activity_id`);

--
-- Índices para tabela `image_activities`
--
ALTER TABLE `image_activities`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_img_act_area_id` (`area_id`);

--
-- Índices para tabela `one_question_n_answer_activities`
--
ALTER TABLE `one_question_n_answer_activities`
  ADD PRIMARY KEY (`id`),
  ADD KEY `o_que_n_ans_area_fk` (`area_id`);

--
-- Índices para tabela `text_activities`
--
ALTER TABLE `text_activities`
  ADD PRIMARY KEY (`id`),
  ADD KEY `txt_act_fk` (`area_id`);

--
-- Índices para tabela `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Índices para tabela `user_has_area_access`
--
ALTER TABLE `user_has_area_access`
  ADD PRIMARY KEY (`id`),
  ADD KEY `area_id` (`area_id`),
  ADD KEY `user_id` (`user_id`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `answer_n_to_one`
--
ALTER TABLE `answer_n_to_one`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `area`
--
ALTER TABLE `area`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `classes`
--
ALTER TABLE `classes`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `class_see_content`
--
ALTER TABLE `class_see_content`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `class_takes_user`
--
ALTER TABLE `class_takes_user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `contents`
--
ALTER TABLE `contents`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `content_see_activity`
--
ALTER TABLE `content_see_activity`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `image_activities`
--
ALTER TABLE `image_activities`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `one_question_n_answer_activities`
--
ALTER TABLE `one_question_n_answer_activities`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `text_activities`
--
ALTER TABLE `text_activities`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `user_has_area_access`
--
ALTER TABLE `user_has_area_access`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- Restrições para despejos de tabelas
--

--
-- Limitadores para a tabela `answer_n_to_one`
--
ALTER TABLE `answer_n_to_one`
  ADD CONSTRAINT `ans_n_to_one_area_fk` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`),
  ADD CONSTRAINT `one_question_n_answer_activity_fk` FOREIGN KEY (`one_question_n_answer_activity_id`) REFERENCES `one_question_n_answer_activities` (`id`);

--
-- Limitadores para a tabela `area`
--
ALTER TABLE `area`
  ADD CONSTRAINT `area_ibfk_1` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`);

--
-- Limitadores para a tabela `classes`
--
ALTER TABLE `classes`
  ADD CONSTRAINT `classes_ibfk_1` FOREIGN KEY (`creator_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `classes_ibfk_2` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`);

--
-- Limitadores para a tabela `class_see_content`
--
ALTER TABLE `class_see_content`
  ADD CONSTRAINT `class_see_content_ibfk_1` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  ADD CONSTRAINT `class_see_content_ibfk_2` FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`);

--
-- Limitadores para a tabela `class_takes_user`
--
ALTER TABLE `class_takes_user`
  ADD CONSTRAINT `class_takes_user_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `class_takes_user_ibfk_2` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Limitadores para a tabela `contents`
--
ALTER TABLE `contents`
  ADD CONSTRAINT `contents_ibfk_2` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`);

--
-- Limitadores para a tabela `content_see_activity`
--
ALTER TABLE `content_see_activity`
  ADD CONSTRAINT `content_s_act_area_fk` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`),
  ADD CONSTRAINT `content_s_act_content_fk` FOREIGN KEY (`content_id`) REFERENCES `contents` (`id`),
  ADD CONSTRAINT `content_s_act_image_act_fk` FOREIGN KEY (`image_activity_id`) REFERENCES `image_activities` (`id`),
  ADD CONSTRAINT `content_s_act_one_q_n_asw_fk` FOREIGN KEY (`one_question_n_answer_activity_id`) REFERENCES `one_question_n_answer_activities` (`id`),
  ADD CONSTRAINT `content_s_act_text_act_fk` FOREIGN KEY (`text_activity_id`) REFERENCES `text_activities` (`id`);

--
-- Limitadores para a tabela `image_activities`
--
ALTER TABLE `image_activities`
  ADD CONSTRAINT `fk_img_act_area_id` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`);

--
-- Limitadores para a tabela `one_question_n_answer_activities`
--
ALTER TABLE `one_question_n_answer_activities`
  ADD CONSTRAINT `o_que_n_ans_area_fk` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`);

--
-- Limitadores para a tabela `text_activities`
--
ALTER TABLE `text_activities`
  ADD CONSTRAINT `txt_act_fk` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`);

--
-- Limitadores para a tabela `user_has_area_access`
--
ALTER TABLE `user_has_area_access`
  ADD CONSTRAINT `user_has_area_access_ibfk_1` FOREIGN KEY (`area_id`) REFERENCES `area` (`id`),
  ADD CONSTRAINT `user_has_area_access_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;