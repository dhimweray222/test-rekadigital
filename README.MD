# Attendance Apps

## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/dhimweray222/test-rekadigital
    cd Test-BE-UNINET
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```
3. Change Config in .env file
    ```bash
    # Server Settings:
    SERVER_URI=localhost
    SERVER_PORT=3000
    ENDPOINT_PREFIX=/transactions
    DEFAULT_LIMIT=10

    # Database Setting
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=test_be_rekadigital
    DB_USERNAME=postgres
    DB_PASSWORD=dimasslalu123
    DB_POOL_MIN=10
    DB_POOL_MAX=100
    DB_TIMEOUT=10
    DB_MAX_IDLE_TIME_SECOND=60
    DB_CONNECTION_TIMEOUT=10

    ```
4. Run the application:

    ```bash
    go run main.go
    ```

5. Open your web browser and visit [http://localhost:3000](http://localhost:3000) to see the app in action.

## Features
### Postman Documentation
```
https://documenter.getpostman.com/view/23663611/2s9Ykn9Mjr
```


## Answer 
```
1. github.com/go-playground/validator/v10 v10.16.0:
Ini adalah library untuk melakukan validasi pada struktur data dalam Go (Golang). Dengan menggunakan validator, kita dapat dengan mudah menetapkan aturan validasi untuk struktur data, seperti memastikan bahwa suatu nilai tidak kosong, memiliki panjang tertentu, atau memenuhi kriteria tertentu. Validator digunakan untuk memastikan bahwa data yang dimasukkan atau diproses dalam aplikasi memenuhi aturan atau kriteria tertentu. Hal ini membantu meningkatkan kualitas data dan mencegah kesalahan masukan yang tidak valid.

2. github.com/gofiber/fiber/v2 v2.52.0:
Fiber adalah framework web yang ringan untuk bahasa pemrograman Go. Fiber dikembangkan untuk memberikan kinerja yang sangat tinggi dengan sintaks yang mudah dan sederhana. Fiber dirancang agar cepat dan efisien, dan sering digunakan dalam pengembangan aplikasi web dengan Go.
Fiber digunakan sebagai framework web untuk membangun aplikasi web dengan Go. Dengan menggunakan Fiber, pengembang dapat dengan cepat membuat aplikasi web yang efisien dan mudah dipelihara.

3. github.com/google/uuid v1.5.0:
Package ini menyediakan implementasi UUID (Universally Unique Identifier) sesuai dengan spesifikasi RFC 4122. UUID sering digunakan untuk mengidentifikasi entitas secara unik dan secara acak di aplikasi atau sistem yang berbeda. Penggunaan UUID membantu menghindari konflik ID dan memudahkan integrasi sistem yang berbeda. UUID digunakan untuk memberikan identifikasi unik pada entitas dalam aplikasi. Ini membantu mencegah konflik ID dan memastikan keunikan identifikasi di seluruh sistem.

4. github.com/jackc/pgx/v5 v5.5.1:
Pgx adalah driver database PostgreSQL untuk Go. Dengan menggunakan Pgx, kita dapat berinteraksi dengan database PostgreSQL dari aplikasi Go. Ini memungkinkan eksekusi query, pengelolaan transaksi, dan lainnya. Pgx digunakan untuk berinteraksi dengan database PostgreSQL dari aplikasi Go. Dengan menggunakan driver ini, aplikasi dapat melakukan operasi database seperti eksekusi query, pengelolaan transaksi, dan lainnya.

5. github.com/joho/godotenv v1.5.1:
Godotenv adalah library yang membantu app membaca variabel lingkungan dari file .env. Variabel lingkungan sering digunakan untuk menyimpan konfigurasi aplikasi, seperti informasi koneksi database atau kredensial API, tanpa harus secara langsung menuliskan nilainya dalam kode.Godotenv membantu dalam manajemen konfigurasi aplikasi dengan membaca variabel lingkungan dari file .env. Ini memungkinkan pengguna untuk menyimpan konfigurasi, seperti kredensial atau pengaturan koneksi, secara terpisah dari kode aplikasi.

Penggunaan teknologi ini mempermudah pengembangan aplikasi Go, meningkatkan keamanan dan kualitas kode, serta memisahkan konfigurasi dari kode sumber, membuat aplikasi lebih mudah dikonfigurasi dan dikelola.




```




