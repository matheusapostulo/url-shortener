#!/bin/bash

echo "⏳ Waiting for MySQL Master to be online..."
until mysql -h mysql-master -u root -proot -e "SHOW DATABASES;" > /dev/null 2>&1; do
    echo "🔄 Still waiting for MySQL Master..."
    sleep 5
done
echo "✅ MySQL Master is online!"

echo "🔧 Configuring MySQL Replica..."
mysql -h localhost -u root -proot -e "STOP REPLICA; CHANGE REPLICATION SOURCE TO SOURCE_HOST='mysql-master', SOURCE_USER='root', SOURCE_PASSWORD='root', SOURCE_AUTO_POSITION=1; SET GLOBAL SQL_SLAVE_SKIP_COUNTER = 1; START REPLICA;"

echo "🎉 MySQL Replica is up and running!"