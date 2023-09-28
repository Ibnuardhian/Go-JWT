# Go-JWT

## Deskripsi

Repository ini berisi implementasi JSON Web Token (JWT) untuk autentikasi pengguna dalam aplikasi yang dibangun menggunakan bahasa pemrograman Go. JWT adalah sebuah standar (RFC 7519) yang mendefinisikan cara untuk mengirimkan informasi antara dua pihak yang terpercaya menggunakan objek JSON.

## Fitur

- Registrasi Pengguna
- Login Pengguna
- Autentikasi dan Otorisasi menggunakan JWT

## Persyaratan

- Go versi 1.16 atau lebih baru
- MySQL

## Cara Penggunaan

1. **Clone Repository:**
   ```sh
   git clone https://github.com/Ibnuardhian/Go-JWT.git
   cd Go-JWT
   ```

2. **Setup Database:**
   Buat database MySQL dan konfigurasikan kredensial database di file `gojwt`.

3. **Install Dependencies:**
4. 
   ```sh
   go mod tidy
   ```

5. **Run Aplikasi:**
   ```sh
   go run main.go
   ```

## Endpoints

- `POST /signup`: Registrasi pengguna baru
- `POST /login`: Login pengguna dan mendapatkan token JWT
- `GET /validate`: Mendapatkan informasi pengguna yang telah login
