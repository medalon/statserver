CREATE TABLE `banners` (
  `id` int(11) NOT NULL,
  `banner_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `date` date NOT NULL,
  `show_kg` int(11) NOT NULL,
  `show_wr` int(11) NOT NULL,
  `click_kg` int(11) NOT NULL,
  `click_wr` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `banners`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `banners`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
