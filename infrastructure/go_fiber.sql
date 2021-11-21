CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username varchar(100),
                       password varchar(255),
                       name varchar(255),
                       last_name varchar(100),
                       email varchar(255),
                       remember_token varchar(100),
                       action_token varchar(255),
                       pin varchar(10),
                       verify varchar(25),
                       verify_date date,
                       mobile varchar(25),
                       type varchar(25),
                       status varchar(25),
                       user_group_id int,
                       gender varchar(25),
                       birthday date DEFAULT NULL,
                       created_at timestamp NULL DEFAULT NULL,
                       updated_at timestamp NULL DEFAULT NULL,
                       deleted_at timestamp NULL DEFAULT NULL
);

--
-- Dumping data for table `users`
--

INSERT INTO users (id, username, password, name, last_name, email, remember_token, action_token, pin, verify, verify_date, mobile, type, status, user_group_id, gender, birthday, created_at, updated_at, deleted_at) VALUES
(1, 'admin@admin.com', '$2y$10$NKE.pUShl4/JPqRpPgym1.fhXaumalps/lIrv2x0B6Iy9gM4GU0qG', 'Superadmin', 'administator', 'admin@admin.com', NULL, NULL, NULL, 'waiting', NULL, '1234567890', 'admin', 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL),
(2, 'user@user.com', '$2y$10$bK94hCGGApJA.RLEWZc/quVVpwceL0hD0OblbXr2QOJo4aeldfIfu', 'User', '', 'user@user.com', NULL, NULL, NULL, 'waiting', NULL, '', NULL, 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL),
(3, 'sirichai.jan@ascendcorp.com', '$2y$10$XEU4zOztvil8F5BKiCHMtueVVuFfMVhkg3MB3.36mxVB5AV2IJjSW', 'Dev Boy', 'Team P ches', 'sirichai.jan@ascendcorp.com', NULL, NULL, NULL, 'waiting', NULL, '1234567890', 'admin', 'active', 1, NULL, NULL, '2020-12-14 23:32:35', '2020-12-14 23:32:35', NULL);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name varchar(255),
                          price NUMERIC(5,2),
                          created_by int,
                          updated_by int,
                          created_at timestamp NULL DEFAULT NULL,
                          updated_at timestamp NULL DEFAULT NULL,
                          deleted_at timestamp NULL DEFAULT NULL
);