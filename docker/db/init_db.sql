# DBの作成
CREATE DATABASE sau_na_dev CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

# ユーザーの作成
CREATE USER api@'%';
GRANT ALL PRIVILEGES ON *.* TO api;

# rootユーザーのリモート接続を許可
RENAME USER root@'localhost' to root@'%';
