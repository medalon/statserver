CREATE TABLE IF NOT EXISTS `prerolls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `preroll_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `date` date NOT NULL,
  `show_kg` int(11) NOT NULL,
  `show_wr` int(11) NOT NULL,
  `click_kg` int(11) NOT NULL,
  `click_wr` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `preroll`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `preroll`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
