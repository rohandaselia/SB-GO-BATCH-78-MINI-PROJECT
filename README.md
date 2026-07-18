# Car Rental API

## Kegunaan (Project Overview)
Platform Car Rental adalah backend API yang mempertemukan pemilik mobil (Owner) dengan penyewa (Customer) secara langsung. 
Sistem ini memfasilitasi manajemen inventaris mobil, pencarian jadwal kendaraan yang presisi tanpa risiko *double booking* (menggunakan *Exclusion Constraint* pada PostgreSQL), integrasi pembayaran menggunakan Xendit (Invoices & Webhooks), serta sistem ulasan untuk menyelesaikan siklus penyewaan.

## Teknologi Utama
- **Bahasa & Framework:** Golang, Gin Framework
- **Database:** PostgreSQL (driver: `github.com/lib/pq`)
- **Migrasi Database:** `github.com/rubenv/sql-migrate`
- **Keamanan:** JWT Token & Bcrypt Hashing
- **Payment Gateway:** Xendit

---

## Cara Penggunaan (Cara Menjalankan Server Lokal)

### 1. Persiapan Database & File `.env`
Pastikan Anda memiliki PostgreSQL yang menyala. Buat database kosong, misalnya `p2p_car_rental`. 
Kemudian, sesuaikan konfigurasi environment di file `.env` yang berada di root direktori:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password_database_anda
DB_NAME=p2p_car_rental

JWT_SECRET=super-secret-key-p2p-rent

# Didapatkan dari Dashboard Xendit
XENDIT_SECRET_KEY=xnd_development_...
XENDIT_CALLBACK_TOKEN=token_webhook_anda...
```

### 2. Migrasi Skema Database
Sistem ini menggunakan `sql-migrate`. Untuk mengaplikasikan skema dan constraint keamanan (anti-double booking):
1. Install tool migrasi jika belum ada: `go install github.com/rubenv/sql-migrate/...@latest`
2. Eksekusi migrasi menggunakan perintah: `sql-migrate up`
*(Pastikan Anda memiliki file konfigurasi dbconfig.yml atau mengeksekusi script SQL di folder migrations/ secara manual jika tidak menggunakan CLI).*

### 3. Menjalankan Aplikasi
Install semua dependensi (jika belum):
```bash
go mod tidy
```
Jalankan server Gin:
```bash
go run main.go
```
Server akan menyala di port `http://localhost:8080`.

---

## Daftar Path API Tersedia (API Endpoints)

### A. Endpoint Publik (Public Routes)
| Method | Endpoint | Kegunaan |
|--------|----------|----------|
| GET | `/health` | Memeriksa status server API |
| POST | `/api/v1/auth/register` | Mendaftarkan akun (Role: `customer` atau `owner`) |
| POST | `/api/v1/auth/login` | Login untuk mendapatkan token JWT |
| GET | `/api/v1/cars/search` | Mencari mobil dengan filter `start_date` & `end_date` |
| POST | `/api/v1/webhooks/xendit` | Webhook Xendit untuk update status pembayaran otomatis |
| GET | `/public/images/*filepath` | Mengakses file/foto mobil statis yang diunggah |

### B. Endpoint Owner (Membutuhkan JWT Role: Owner)
*Header: `Authorization: Bearer <token>`*
| Method | Endpoint | Kegunaan |
|--------|----------|----------|
| POST | `/api/v1/owner/cars` | Menambahkan data mobil baru |
| GET | `/api/v1/owner/cars` | Melihat daftar mobil yang dimiliki oleh owner login |
| PUT | `/api/v1/owner/cars/:id` | Mengubah detail data mobil (harga, deskripsi, dll) |
| DELETE | `/api/v1/owner/cars/:id` | Menghapus data mobil dari platform |
| POST | `/api/v1/owner/cars/:id/images`| Mengunggah foto mobil fisik (`multipart/form-data`) |
| PUT | `/api/v1/owner/bookings/:id/status`| Memajukan status sewa (`paid` -> `active` -> `completed`) |

### C. Endpoint Customer (Membutuhkan JWT Role: Customer)
*Header: `Authorization: Bearer <token>`*
| Method | Endpoint | Kegunaan |
|--------|----------|----------|
| POST | `/api/v1/customer/bookings` | Melakukan checkout/booking mobil, men-generate Xendit Invoice |
| POST | `/api/v1/customer/bookings/:id/reviews`| Memberikan rating (1-5) & ulasan jika sewa berstatus `completed` |

---
