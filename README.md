Sport Booking CLI (Padel Court Booking)

Sport Booking CLI adalah aplikasi Command Line Interface (CLI) berbasis Golang untuk melakukan pemesanan lapangan olahraga (Padel Court). Aplikasi ini mendukung role Admin dan User, terhubung ke MySQL, dan menerapkan struktur project yang terpisah antara CLI, business logic, dan entity.

Fitur Utama
👤 User

Register & Login

Melihat daftar lapangan yang tersedia

Melakukan booking lapangan

Melihat riwayat booking pribadi

Logout

🛠 Admin

Login sebagai admin

Mengelola data lapangan

Melihat data booking

Membuat laporan (Daily & Monthly Report)

🏗️ Struktur Folder
SPORT-BOOKING
│
├── cli
│   ├── admin        # Menu & interaksi CLI untuk admin
│   └── user         # Menu & interaksi CLI untuk user
│
├── config           # Konfigurasi database & environment
│
├── database         # SQL schema & query database
│
├── entity           # Struct Golang (mapping tabel database)
│
├── handler          # Business logic & query database
│
├── .env             # Environment variable (DSN database)
├── go.mod
├── go.sum
├── main.go          # Entry point aplikasi
└── README.md

🗄️ Database Design

Aplikasi menggunakan beberapa tabel utama:

roles

users

courts

time_slots

bookings

reports

Relasi database dirancang dengan foreign key untuk menjaga integritas data dan mencegah double booking pada lapangan dan waktu yang sama.

⚙️ Tech Stack

Golang

MySQL

Github

CLI (Terminal Based App)

🔧 Setup & Installation
1️⃣ Clone Repository
git clone https://github.com/username/sport-booking.git
cd sport-booking

2️⃣ Install Dependency
go mod tidy

3️⃣ Setup Database

Import SQL schema dari folder database

Pastikan MySQL sudah berjalan

4️⃣ Setup Environment Variable

Buat file .env:

DB_DSN=username:password@tcp(localhost:3306)/sport_booking

5️⃣ Run Application
go run main.go