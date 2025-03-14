#!/bin/bash

until mysql -h mysql-master -u root -proot -e "SHOW DATABASES;" > /dev/null 2>&1; do
    echo "Waiting for MySQL Master to be online..."
    sleep 5
done

mysql -h localhost -u root -proot -e "STOP REPLICA; CHANGE REPLICATION SOURCE TO SOURCE_HOST='mysql-master', SOURCE_USER='root', SOURCE_PASSWORD='root', SOURCE_AUTO_POSITION=1; START REPLICA;"

echo "MySQL Slave is up and running!"
