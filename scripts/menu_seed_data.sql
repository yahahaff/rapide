-- Menu seed data for vue-vben-admin integration
-- Make sure we're using the right database
USE rapide;

-- Clear existing menu data if needed
-- TRUNCATE TABLE sys_menu;

-- Reset auto increment if needed
-- ALTER TABLE sys_menu AUTO_INCREMENT = 1;

-- Dashboard (Root menu)
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    1, NULL, 'Dashboard', '/dashboard', 'LAYOUT', '/dashboard/analysis',
    'Dashboard', 'ion:grid-outline', '', 0, -1,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Dashboard - Analysis
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    2, 1, 'Analysis', 'analysis', '/dashboard/analysis/index', '',
    'Analysis', 'ion:bar-chart-outline', 'dashboard:analysis', 1, 1,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Dashboard - Workbench
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    3, 1, 'Workbench', 'workbench', '/dashboard/workbench/index', '',
    'Workbench', 'ion:layers-outline', 'dashboard:workbench', 1, 2,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System Management (Root menu)
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    10, NULL, 'System', '/system', 'LAYOUT', '/system/user',
    'System Management', 'ion:settings-outline', '', 0, 1000,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System - User Management
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    11, 10, 'UserManagement', 'user', '/system/user/index', '',
    'User Management', 'ion:people-outline', 'system:user:list', 1, 1,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System - Role Management
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    12, 10, 'RoleManagement', 'role', '/system/role/index', '',
    'Role Management', 'ion:key-outline', 'system:role:list', 1, 2,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System - Menu Management
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    13, 10, 'MenuManagement', 'menu', '/system/menu/index', '',
    'Menu Management', 'ion:menu-outline', 'system:menu:list', 1, 3,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System - Department Management
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    14, 10, 'DeptManagement', 'dept', '/system/dept/index', '',
    'Department Management', 'ion:git-branch-outline', 'system:dept:list', 1, 4,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- System - Log Management
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    15, 10, 'LogManagement', 'log', '/system/log/index', '',
    'Log Management', 'ion:document-text-outline', 'system:log:list', 1, 5,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Account (Root menu)
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    20, NULL, 'Account', '/account', 'LAYOUT', '/account/settings',
    'Account', 'ion:person-outline', '', 0, 2000,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Account - Settings
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    21, 20, 'AccountSettings', 'settings', '/account/settings/index', '',
    'Account Settings', 'ion:settings-outline', '', 1, 1,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Account - Change Password
INSERT INTO `sys_menu` (
    `id`, `parent_id`, `name`, `path`, `component`, `redirect`, 
    `title`, `icon`, `permission`, `type`, `order_no`, 
    `hidden`, `keep_alive`, `ignore_keep_alive`, `hide_breadcrumb`, 
    `hide_children_in_menu`, `current_active_menu`, `status`, 
    `created_at`, `updated_at`
) VALUES (
    22, 20, 'ChangePassword', 'password', '/account/password/index', '',
    'Change Password', 'ion:lock-closed-outline', '', 1, 2,
    0, 1, 0, 0,
    0, '', 1,
    NOW(), NOW()
);

-- Create role-menu relationships for admin role (assuming role_id 1 is admin)
INSERT INTO `sys_role_menus` (`role_id`, `menu_id`) VALUES 
(1, 1), (1, 2), (1, 3),
(1, 10), (1, 11), (1, 12), (1, 13), (1, 14), (1, 15),
(1, 20), (1, 21), (1, 22);

-- Create a default department if it doesn't exist
INSERT INTO `sys_dept` (`id`, `p_code`, `p_codes`, `name`, `sort`, `level`, `tips`, `status`, `created_at`, `updated_at`)
SELECT 1, 1, '1', 'Default Department', 1, 1, 'Default department created by system', 1, NOW(), NOW()
FROM DUAL
WHERE NOT EXISTS (SELECT 1 FROM `sys_dept` WHERE `id` = 1);