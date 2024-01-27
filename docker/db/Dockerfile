FROM mysql/mysql-server:8.0.32

COPY init_db.sql /docker-entrypoint-initdb.d/
COPY my.cnf /usr/etc/my.cnf

ENV MYSQL_ALLOW_EMPTY_PASSWORD=true

CMD ["mysqld"]