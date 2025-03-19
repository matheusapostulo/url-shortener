#!/bin/bash

echo "Waiting for MySQL Master to be ready..."
until mysqladmin ping -h "mysql-master" --silent; do
    sleep 2
done
echo "✅ MySQL Master is ready!"

echo "Waiting for MySQL Replicas to be ready..."
REPLICAS=("mysql-replica1" "mysql-replica2") 

echo "⏳ Waiting for MySQL replicas to start..."
for REPLICA in "${REPLICAS[@]}"; do
    echo "🔍 Checking replica status on $REPLICA..."
    until mysql -h "$REPLICA" -u root -proot -e "SELECT 1" &>/dev/null; do
        echo "❌ Failed to connect to $REPLICA. Retrying..."
        sleep 2
    done
    echo "✅ Replica $REPLICA is ready!"
done

if [ -f /etc/mysqlrouter/mysqlrouter.cnf ]; then
    echo "🛠 Using custom MySQL Router configuration..."
else
    echo "❌ MySQL Router configuration file not found at /etc/mysqlrouter/mysqlrouter.conf"
    exit 1
fi

mysqlrouter --config /etc/mysqlrouter/mysqlrouter.cnf