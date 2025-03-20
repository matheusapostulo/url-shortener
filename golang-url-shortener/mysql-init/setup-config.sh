#!/bin/bash

echo "⏳ Waiting for MySQL Master to be ready..."
until mysqladmin ping -h "mysql-master" --silent; do
    sleep 2
done

echo "🛠 Configuring Innodb Cluster and MySQL Router..."

mysqlsh --uri root:root@mysql-master:3306 --js -e "
  dba.configureInstance('root:root@mysql-master:3306');
  cluster = dba.createCluster('myCluster');
  cluster.addInstance('mysql-replica1', { recoveryMethod: 'clone' });
  cluster.addInstance('mysql-replica2', { recoveryMethod: 'clone' });
"

echo "✅ InnoDB Cluster and MySQL Router are ready!"
