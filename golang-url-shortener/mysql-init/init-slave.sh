#!/bin/bash

until mysql -h mysql-master -u root -proot -e "SHOW DATABASES;" > /dev/null 2>&1; do
    echo "Waiting for MySQL Master to be online..."
    sleep 5
done

MASTER_LOG_FILE=$(mysql -h mysql-master -u root -proot -e "SHOW BINARY LOG STATUS\G" | grep File | awk '{print $2}')
MASTER_LOG_POS=$(mysql -h mysql-master -u root -proot -e "SHOW BINARY LOG STATUS\G" | grep Position | awk '{print $2}')

mysql -h localhost -u root -proot -e "STOP REPLICA; CHANGE REPLICATION SOURCE TO SOURCE_HOST='mysql-master', SOURCE_USER='root', SOURCE_PASSWORD='root', SOURCE_LOG_FILE='$MASTER_LOG_FILE', SOURCE_LOG_POS=$MASTER_LOG_POS; START REPLICA;"

echo "MySQL Slave is up and running!"
