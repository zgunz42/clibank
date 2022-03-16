# 🏦 clibank

An CLI base bank using mysql as database

Resources
1. [ERD Database](https://drawsql.app/next-bigit/diagrams/cli-bank)
2. [Project Instruction](https://docs.google.com/document/d/15zOXYvx6ln4zrvs9mnbCvrL_DFTMg5wWvbCVrDsHzns/edit)


## Cara menjalankan Program
Untuk menjalankan atau mengembangkan program harus melakukan copy `.env.example` ke `.env`
pengaturan bisa di buat sama atau berbeda, Jika dilihat pada `.env.example` maka akan ada seperti
```
DB_USERNAME=test
DB_PASSWORD=test
DB_NAME=test_db
MYSQL_ROOT_PASSWORD=root
# msql DB url
DB_URL="$DB_USERNAME:$DB_PASSWORD@tcp(127.0.0.1:3306)/$DB_NAME?charset=utf8mb4&parseTime=True&loc=Local"
```
**Penjelasan:**
- DB_USERNAME : nama pengguna mysql
- DB_PASSWORD : password pengguna mysql
- MYSQL_ROOT_PASSWORD: root mysql password (untuk akses phpmyadmin)
- DB_URL : alamat database

## Menjalankan Menggunakan Docker Compose (Recomended)

Docker compose ini nanti akan menjadi software yang akan meyiapkan database, dan phpmyadmin
untuk keperluan development. Langkah awal sebelum menginstal docker-compose adalah dengan pertama kali menginstall docker, docker-compose hanya membantu kita memanajement dari semua docker container (virtual machine) ini. untuk menginstall docker bisa melakukan langkah di bawah
- [Install docker di Komputer](https://docs.docker.com/engine/install/) pastikan untuk memilih mendownload versi stabel.
- [Install docker compose](https://docs.docker.com/compose/install/) ikuti instruski sesuai jenis sistem operasi yang di gunakan
- Pindah ke directory folder
- Jalankan perintah `docker-compose up -d`
- Untuk membuat ulang virtual server `docker-compose down`
- Untuk menghentikan container `docker-compose stop`

>Pastikan perintah docker-compose berada di dalam root project directory
>Jika ada masalah bisa lihat pada bagian Troubleshooting pada bagian bawah dokumen ini

## Menjalankan Tanpa Docker
Untuk menjalankan tanpa docker pertama dengan telah menginstall mysql server dan kemudian pada
`.env` file ubah value `DB_URL` ke url database mysql

## Developing Flow
Project ini menggunakan TRUNK Workflow dimana setiap update feature harus di buat seminimal mungkin dan tidak boleh melakukan pull langsung ke `main` branch. Langkah untuk melakukan develop project ini, untuk setiap feature yang akan di buat harus memiliki nama branch seperti `feature/<name>`. Untuk setiap fitur yang di tambahkan minimal di tambahkan test baik unit test
, integration test, atau lainnya. 
Flow membuat feature:
    - lihat pada bagian app, jika nama module tidak dan feature di luar module di sarankan untuk
      membuat module baru
    - Di dalam module ada 3 file pada umumnya `service`, `repository`, `entity` dan `dto`. Semua buisness logic berada di dalam service, interaksi crud umumnya terjadi di dalam repository dan dto merupakan data transfer object yang akan nanti di saling passing antara 
    aplikasi ke service
    - Jika feature merupakan bagian dari cmd promp (tampil di console) maka harus di taruh di dalam folder `cmd` dalam root directory, dan registrasi cmd di dalam `app/application.go`


## Troubeshooting
Beberapa masalah umum dalam pengembangan aplikasi ini saat develop
1. Port alerady in use saat menjalankan `docker-compose up` ?
   ini bisa di atasi dengan `docker stop $(docker ps -aq)` ini akan menghentikan semua
   container yang jalan, sehingga port bisa di pakai
2. File `.env` tidak di tempukan saat sistem berjalan ?
   Pastikan `.env` berada pada root project directory
3. Gagal saat melakukan migrasi yang di lakukan oleh GORM:?
   Pertama export data di database ke sql file dan drop semua
   tabel di dalam database, dan coba menjalankan program lagi