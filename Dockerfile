FROM mysql:8.0
EXPOSE 3306
COPY ./mock_sql_data/ardeo.sql /docker-entrypoint-initdb.d/
# docker run -p 3306:3306 --name mockmysql_c -e MYSQL_ROOT_PASSWORD=root -e MYSQL_ROOT_HOST=% mockmysql
# docker run --name phpmyadmin_c -d --link mockmysql_c:db -p 8080:80 phpmyadmin/phpmyadmin