FROM mariadb:12.2.1-noble-rc
COPY init.sql /docker-entrypoint-initdb.d/init.sql
RUN chmod 644 /docker-entrypoint-initdb.d/init.sql