
# AIRBNB CLONE API KELOMPOK 1

Selamat datang di projek kami, projek ini bertujuan untuk mereplikasi beberapa fitur-fitur pada Airbnb, seperti register, login, mengelola homestays, melakukan reservation, dan melihat daftar reservation yang pernah dilakukan ke dalam bentuk HTTP Restful API yang dapat dikonsumsi oleh frontend.

## 👀 Environment Variables

Untuk menjalankan proyek ini kamu membutuhkan beberapa environment variable yang dapat kamu contoh di .env.example setelah itu kamu dapat  mengeksport dengan menggunakan perintah source .env.

Berikut adalah beberapa environment variabel yang diperlukan:

`DBUSER`
`DBPASS`
`DBHOST`
`DBPORT`
`DBNAME`
`JWTSECRET`
`CLOUDINARY_CLOUD_NAME`
`CLOUDINARY_API_KEY`
`CLOUDINARY_API_SECRET`
`CLOUDINARY_UPLOAD_FOLDER`

Untuk mengetahaui environment terkait cloudinary kamu dapat mengujungi ini

https://cloudinary.com/documentation/go_quick_start

Untuk membuat dan folder agar bisa diassign ke environment `CLOUDINARY_UPLOAD_FOLDER` kamu dapat mengikuti ini

https://cloudinary.com/documentation/dam_folders_collections_sharing

## 🛠️ Installation

Untuk menjalankan program ini pertama kamu harus mengclone repository ini dengan menggunakan perintah

```bash
git clone https://github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB.git
```

masuk ke folder

```bash
cd BE-AIRBNB
```

pastikan golang dan mysql kamu sudah terinstall, jika belum silahkan kunjungi :

https://go.dev/doc/install
https://dev.mysql.com/doc/refman/8.3/en/windows-installation.html

jika sudah silahkan jalankan

```bash
go mod tidy
go run .
```

Dan selamat kamu berhasil menjalankan projek ini 🎊🎊.
Untuk mengetahui endpoint-endpoint yang dapat digunakan kamu dapat melihat pada folder docs.
## 🙋‍♂️ Authors

- [@anggraanutomo](https://www.github.com/anggraanutomo)
- [@ezabintangr](https://github.com/ezabintangr)


## 👨‍💻 Tech Stack

**Server:** Golang, Echo, MySQL

