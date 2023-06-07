# How to run

## Build Docker Image
```shell
docker build -t contoh .
```

## Run Docker Image
```shell
docker run contoh
```

## Docker Compose

Menjalankan semua service:

```shell
docker-compose up -d
```

Shutdown service:

```shell
docker-compose down
```


## Migrasi Data
### Export Database Menggunakan Mysqldump

1. Export database dengan `mysqldump`

```shell
mysqldump -u root namadatabase > namadatabase.sql
```

2. Copy file sql ke dalam container, misalnya di path `/home/namadatabase.sql`

```shell
docker compose cp namadatabase.sql db:/home/namadatabase.sql
```

3. Masuk ke shell DB:

```shell
docker compose exec -it db sh
```

4. Masuk ke sesi mysql:

```shell
mysql -u root -p
```
5. Jalankan perintah:
```shell
use namadatabase;
source /home/namadatabase.sql;
```