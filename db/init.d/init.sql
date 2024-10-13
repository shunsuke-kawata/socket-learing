-- データベースの作成
CREATE DATABASE IF NOT EXISTS task_management;

-- データベースを使用
USE task_management;

-- タスクステータステーブル
CREATE TABLE IF NOT EXISTS statuses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- デバイステーブル
CREATE TABLE IF NOT EXISTS devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    mac_address VARCHAR(17) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- タスクテーブル
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    color_code CHAR(7) NOT NULL,
    status_id INT NOT NULL, -- タスクの現在のステータス
    device_id INT NOT NULL,
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (status_id) REFERENCES task_statuses(id),
    FOREIGN KEY (device_id) REFERENCES devices(id)
);

-- タスクステータスのデフォルトデータの追加
INSERT INTO task_statuses (status_name) VALUES ('Pending'), ('InProgress'), ('Done');

-- デバイステーブルへのテストデータの挿入
INSERT INTO devices (mac_address)
VALUES 
    ('00:14:22:01:23:45'),
    ('00:16:17:AB:CD:EF'),
    ('00:1A:79:12:34:56'),
    ('00:11:22:33:44:55'),
    ('00:AA:BB:CC:DD:EE');

-- タスクテーブルへのテストデータの挿入
INSERT INTO tasks (title, description, color_code, status_id, device_id, due_date)
VALUES 
    ('Task 1', 'Description for Task 1', '#FF5733', 1, 1, '2024-10-20'),  -- Pending
    ('Task 2', 'Description for Task 2', '#33FF57', 2, 2, '2024-10-25'),  -- InProgress
    ('Task 3', 'Description for Task 3', '#3357FF', 3, 3, '2024-11-01'),  -- Done
    ('Task 4', 'Description for Task 4', '#F1C40F', 1, 4, '2024-11-05'),  -- Pending
    ('Task 5', 'Description for Task 5', '#E74C3C', 2, 5, '2024-11-10');  -- InProgress
