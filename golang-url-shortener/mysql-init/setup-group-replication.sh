#!/bin/bash

# Configurando as variáveis globais para o Group Replication
echo "Configurando as variáveis do Group Replication no master..."

mysql -h mysql-master -u root -proot -e "
  SET GLOBAL group_replication_group_name = 'aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee';
  SET GLOBAL group_replication_local_address = 'mysql-master:33061';
  SET GLOBAL group_replication_group_seeds = 'mysql-master:33061,mysql-replica1:33061,mysql-replica2:33061';
  SET GLOBAL group_replication_start_on_boot = OFF;
  SET GLOBAL group_replication_enforce_update_everywhere_checks = OFF;
  SET GLOBAL group_replication_single_primary_mode = ON;
"

echo "Configurando as variáveis do Group Replication na réplica 1..."
mysql -h mysql-replica1 -u root -proot -e "
  INSTALL PLUGIN group_replication SONAME 'group_replication.so';
  SET GLOBAL group_replication_group_name = 'aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee';
  SET GLOBAL group_replication_local_address = 'mysql-replica1:33061';
  SET GLOBAL group_replication_group_seeds = 'mysql-master:33061,mysql-replica1:33061,mysql-replica2:33061';
  SET GLOBAL group_replication_start_on_boot = OFF;
  SET GLOBAL group_replication_enforce_update_everywhere_checks = OFF;
  SET GLOBAL group_replication_single_primary_mode = ON;
"
echo "Configuração do Group Replication na replica 1 finalizada!"

echo "Configurando as variáveis do Group Replication na réplica 2..."
mysql -h mysql-replica2 -u root -proot -e "
  INSTALL PLUGIN group_replication SONAME 'group_replication.so';
  SET GLOBAL group_replication_group_name = 'aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee';
  SET GLOBAL group_replication_local_address = 'mysql-replica2:33061';
  SET GLOBAL group_replication_group_seeds = 'mysql-master:33061,mysql-replica1:33061,mysql-replica2:33061';
  SET GLOBAL group_replication_start_on_boot = OFF;
  SET GLOBAL group_replication_enforce_update_everywhere_checks = OFF;
  SET GLOBAL group_replication_single_primary_mode = ON;
"
echo "Configuração do Group Replication na replica 2 finalizada!"
