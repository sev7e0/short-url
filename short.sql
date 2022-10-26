CREATE TABLE `short_link` (
                              `id` bigint NOT NULL DEFAULT '0' COMMENT '主键',
                              `url` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '链接',
                              `short_url` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '短链',
                              `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='短链'

