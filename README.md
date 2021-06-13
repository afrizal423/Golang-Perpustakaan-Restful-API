<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Sistem Informasi Perpustakaan Berbasis Rest API

Project sederhana membuat Sistem Informasi Perpustakaan berbasis Restful API menggunakan bahasa pemrograman Go dengan database MySql. Alur projectnya kurang lebihnya seperti [ini](https://core.ac.uk/download/pdf/12347733.pdf). Password hash menggunakan argon2id

### User Level :boy: :woman:
- Pegawai 
- Siswa  <i>(soon)</i> 

### ToDo List Pengerjaan Project :pushpin:
- [x] Create Migration table
- [x] Create Seed data akun admin
- [x] Register siswa
- [x] Login With JWT
- [x] CRUDF Buku
    - [x] CRUDF Jenis buku
    - [x] CRUDF Penulis buku
    - [x] CRUDF Penerbit buku
- [ ] CRUDF peminjaman
    - [ ] CRUDF Detail pinjam
    - [ ] CRUDF Denda
- [ ] **DONE!!** 

# How To Install
- Download Go terlebih dahulu [disini](https://golang.org/dl/)
- Download repo [ini](https://github.com/afrizal423/Golang-Perpustakaan-Restful-API/archive/master.zip)
- Masuk kedalam foldernya
- Rename file ```.env.sample``` menjadi ```.env```
- buka file ```.env```, silahkan ubah sesuai konfigurasi yang dimiliki
**NOTE!!**
    - Jangan lupa hapus tagar (#) pada file ```.env```.
    - Pilih lihat variabel MIGRATE terdapat true atau false, pilih salah satu, jika true sistem akan migrate ketika pertama dijalankan **(akan drop data juga sebelum di migrate)**, jika false maka tidak akan migrate ke database **(tidak akan drop data)**.
    - contoh
        ```
        API_SECRET=989898kjashndflkajsndflkasndf
        MIGRATE=false
        DB_HOST=127.0.0.1
        DB_DRIVER=mysql 
        DB_USER=root
        DB_PASSWORD=
        DB_NAME=go_perpus
        DB_PORT=3306 
        ```
- Jalankan perintah ```go run main.go``` untuk menjalankan program