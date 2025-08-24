# Репликация в PostgreSQL

бд настроена локально, поэтому настроим репликацию

1. Создадим юзера для репликации

`CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'pass';`

2. Настраиваем `postgresql.conf`

```
wal_level = replica
max_wal_senders = 2
synchronous_commit = on
synchronous_standby_names = 'FIRST 1 (replica1, replica2)'
```

3. Настроим `pg_hba.conf` для реплик

```
host    replication     replicator      192.168.1.0/24          md5
```

4. Перезапускаем postgres

```
sudo systemctl restart postgresql
```

## Настройка реплики 1

1. Создаем бекап мастера для реплики 1

```
pg_basebackup -h localhost -U replicator -D replica1 -Fp -Xs -P -R
```

2. Настроим `postgresql.conf`

```
primary_conninfo = 'host=127.0.0.1 port=5432 user=replicator password=pass application_name=replica1'
port = 5433
```

## Настройка реплики 2

1. Создаем бекап мастера для реплики 2

```
pg_basebackup -h localhost -U replicator -D replica2 -Fp -Xs -P -R
```

2. Настроим `postgresql.conf`

```
primary_conninfo = 'host=127.0.0.1 port=5432 user=replicator password=pass application_name=replica2'
port = 5434
```

## Проверки

1. Проверим состояние репликации на мастере

```
psql -h 127.0.0.1 -U otus_user -d otus -c "SELECT application_name, sync_state FROM pg_stat_replication;"
```

2. Проверим кол-во записей в таблице users на репликах

```
psql -h 127.0.0.1 -p 5433 -U otus_user -d otus -c "SELECT COUNT(*) FROM users;"
psql -h 127.0.0.1 -p 5434 -U otus_user -d otus -c "SELECT COUNT(*) FROM users;"
output: 1000013
```
