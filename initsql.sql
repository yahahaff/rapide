INSERT INTO `rapide`.`sys_role` (`id`, `num`, `pid`, `name`, `tips`, `version`) VALUES (1, 1, 1, 'administrator', '超级管理员', NULL);
INSERT INTO `rapide`.`sys_dept` (`id`, `num`, `p_id`, `pids`, `full_name`, `tips`, `created_at`, `updated_at`) VALUES (1, 1, 1, '1', 'devops', '运维', NULL, NULL);
# 创建admin账号 admin  admin
INSERT INTO `rapide`.`sys_user` (`id`, `name`, `email`, `phone`, `password`, `avatar`, `role_id`, `dept_id`, `status`, `otp_enabled`, `otp_verified`, `otp_secret`, `otp_auth_url`, `introduction`, `created_at`, `updated_at`) VALUES (1, 'admin', NULL, NULL, '$2a$14$mSZW8OoUALjrL3Oyn66mHORF2mYXYG1c1gitdwJbIMFFu4EmffgqO', '', 1, 1, 0, 0, 0, '', '', '', '2024-06-10 16:36:13.244', '2024-06-10 16:36:13.244');
