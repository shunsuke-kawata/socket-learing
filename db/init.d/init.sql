-- データベースの作成
CREATE DATABASE IF NOT EXISTS task_management;

-- データベースを使用
USE task_management;

-- タスクステータステーブル
CREATE TABLE IF NOT EXISTS statuses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- デバイステーブル
CREATE TABLE IF NOT EXISTS addresses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ip_address VARCHAR(39) NOT NULL UNIQUE,
    color_code CHAR(7) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- タスクテーブル
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status_id INT NOT NULL, -- タスクの現在のステータス
    address_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (status_id) REFERENCES statuses(id),
    FOREIGN KEY (address_id) REFERENCES addresses(id)
);

-- タスクステータスのデフォルトデータの追加
INSERT INTO statuses (status_name) VALUES ('Pending'), ('InProgress'), ('Done');

-- デバイステーブルへのテストデータの挿入
INSERT INTO addresses (ip_address, color_code)
VALUES 
    ('000.000.0.0', "#000000"),
    ('192.168.0.2', "#FF00FF"),
    ('192.168.0.3', "#FFFF00")
;

-- タスクテーブルへのテストデータの挿入
INSERT INTO tasks (title, description, status_id, address_id)
VALUES 
    ('Task 1', 'Description for Task 1', 1, 1),  -- Pending
    ('Task 2', 'Description for Task 2', 2, 2),  -- InProgress
    ('Task 3', 'Description for Task 3', 3, 3) -- Done
;