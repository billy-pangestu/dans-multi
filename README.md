# BE MAIN TEMPLATE

![Architecture Clean Code](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

First Step
 - create main.go
 - create .gitignore
 - create go module 

    `go mod init <project-name>`

### Ada 6 bagian kode di dalam design ini:

## 1. Model
Model dapat berisikan objek data dengan metode atau hanya berupa kumpulan struktur data.

## 2. Repository
Repository berisikan perintah untuk query ke database.

## 3. Usecase
Usecase berisi bisnis logic aplikasi, mengatur aliran data ke dan dari model, dan mengarahkan data tersebut untuk diolah.

## 4. Transport
Transport digunakan untuk mengolah request sebelum masuk ke usecase dan mengolah response setelah kembali dari usecase.

## 5. Routes
Routes untuk mengolah endpoint, menerapkan middleware, dan menjalankan service.

## 6. Adapter
Adapter berisikan fungsi-fungsi untuk membuat koneksi service ke database atau client lain.

# Helper
Helper berisikan middleware, response template, fungsi crypto, logger, dll. Pengguna bisa langsung memanggil fungsi didalam helper jika diperlukan.

# References
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html