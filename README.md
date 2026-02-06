## 🏓 Padel Booking CLI Application

Padel Booking CLI Application adalah aplikasi Command Line Interface (CLI) berbasis Golang untuk mengelola sistem pemesanan lapangan padel.
Aplikasi ini dirancang untuk mensimulasikan sistem booking nyata dengan fitur role-based access (Admin & User), terintegrasi langsung dengan database MySQL.

## 📌 Business Case

Banyak tempat olahraga masih mengelola pemesanan lapangan secara manual (WhatsApp, catatan kertas, atau Excel), yang sering menyebabkan:

1. Double booking
2. Kesalahan jadwal
3. Sulit membuat laporan pendapatan

## 🎯 Objectives

Membangun aplikasi CLI menggunakan Golang

Mengimplementasikan CRUD database MySQL

Menerapkan clean folder structure

Menggunakan konsep role (Admin & User)

Menghindari double booking dengan database constraint

## 👥 User Roles

### Admin

Login sebagai admin

Melihat data booking

Mengelola laporan (daily & monthly)

### User

Registrasi & login

Melihat lapangan yang tersedia

Melakukan booking lapangan

Melihat riwayat booking sendiri

## ✨ Features
1. Authentication
2. Register user
3. Login user
4. Role-based menu (Admin / User)
6. View available courts
7. Create booking
8. View my bookings
9. View all bookings
10. Generate reports (daily & monthly)

## 🚀 How to Run

### 1. Clone Repository

git clone https://github.com/username/Paddle-Booking.git
cd Paddle-Booking

### 2. Setup Database

Buat database MySQL

Jalankan file SQL DDL

Isi data awal (roles, courts, time slots)

### 3. Setup Environment

Buat file .env:

DB_DSN=username:password@tcp(localhost:3306)/padel_booking

### 4. Run Application
go run main.go🚀 How to Run
#### 1. Clone Repository
git clone (https://github.com/andrewidianto12/Booking-Paddle)
cd Paddle-Booking

#### 2. Setup Database

Buat database MySQL

Jalankan file SQL DDL

Isi data awal (roles, courts, time slots)

#### 3. Setup Environment

Buat file .env:

DB_DSN=username:password@tcp(localhost:3306)/Database

#### 4. Run Application
go run main.go
