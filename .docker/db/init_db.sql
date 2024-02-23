# DBの作成
CREATE DATABASE sauna CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

# ユーザーの作成
CREATE USER sauna@'%';
GRANT ALL PRIVILEGES ON *.* TO sauna;

# rootユーザーのリモート接続を許可
RENAME USER root@'localhost' to root@'%';
