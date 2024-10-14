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
CREATE TABLE IF NOT EXISTS addresses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ip_address VARCHAR(39) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- タスクテーブル
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    color_code CHAR(7) NOT NULL,
    status_id INT NOT NULL, -- タスクの現在のステータス
    address_id INT NOT NULL,
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (status_id) REFERENCES statuses(id),
    FOREIGN KEY (address_id) REFERENCES addresses(id)
);

-- タスクステータスのデフォルトデータの追加
INSERT INTO statuses (status_name) VALUES ('Pending'), ('InProgress'), ('Done');

-- デバイステーブルへのテストデータの挿入
INSERT INTO addresses (ip_address)
VALUES 
    ('192.168.0.1'),
    ('192.168.0.2'),
    ('192.168.0.3'),
    ('192.168.0.4'),
    ('192.168.0.5');

-- タスクテーブルへのテストデータの挿入
INSERT INTO tasks (title, description, color_code, status_id, address_id, due_date)
VALUES 
    ('Task 1', 'Description for Task 1', '#FF5733', 1, 1, '2024-10-20'),  -- Pending
    ('Task 2', 'Description for Task 2', '#33FF57', 2, 2, '2024-10-25'),  -- InProgress
    ('Task 3', 'Description for Task 3', '#3357FF', 3, 3, '2024-11-01'),  -- Done
    ('Task 4', 'Description for Task 4', '#F1C40F', 1, 4, '2024-11-05'),  -- Pending
    ('Task 5', 'Description for Task 5', '#E74C3C', 2, 5, '2024-11-10');  -- InProgress
