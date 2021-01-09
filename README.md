<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Sistem Informasi Perpustakaan Berbasis Rest API

Project sederhana membuat Sistem Informasi Perpustakaan berbasis Restful API menggunakan bahasa pemrograman Go dengan database MySql. Alur projectnya kurang lebihnya seperti [ini](https://core.ac.uk/download/pdf/12347733.pdf). Password hash menggunakan argon2id

### User Level :boy: :woman:
- Pegawai 
- Siswa

### ToDo List Pengerjaan Project :pushpin:
- [x] Create Migration table
- [x] Create Seed data akun admin
- [x] Register siswa
- [x] Login With JWT
- [ ] CRUDF Buku
    - [ ] CRUDF Detail buku
    - [x] CRUDF Jenis buku
    - [ ] CRUDF Penulis buku
    - [ ] CRUDF Penerbit buku
- [ ] CRUDF peminjaman
    - [ ] CRUDF Detail pinjam
    - [ ] CRUDF Denda
- [ ] <b> DONE!!</b> 

# How To Install
- Download Go terlebih dahulu [disini](https://golang.org/dl/)
- Download repo [ini](https://github.com/afrizal423/Golang-Perpustakaan-Restful-API/archive/master.zip)
- Masuk kedalam foldernya
- Rename file ```.env.sample``` menjadi ```.env```
- buka file ```.env```, silahkan ubah sesuai konfigurasi yang dimiliki
- Jalankan perintah ```go run main.go``` untuk menjalankan program