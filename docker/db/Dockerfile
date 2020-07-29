FROM mariadb:10.5

ENV DIR /docker-entrypoint-initdb.d

ADD create.sql "$DIR"/001_create.sql
ADD insert.sql "$DIR"/002_insert.sql
