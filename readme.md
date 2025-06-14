# Movie API Backend

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)
![Gin Framework](https://img.shields.io/badge/Gin-1.9%2B-lightgrey)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15%2B-blue)
![Redis](https://img.shields.io/badge/Redis-7%2B-red)
![Swagger](https://img.shields.io/badge/Swagger-2.0%2B-green)
![License](https://img.shields.io/badge/License-MIT-yellow)

## Daftar Isi

- [Fitur](#-fitur)
- [Persyaratan](#-persyaratan)
- [Environment](#-environment)
- [Instalasi](#-instalasi)
- [Konfigurasi](#-konfigurasi)
- [Dokumentasi API](#-dokumentasi-api)
- [Struktur Proyek](#-struktur-proyek)
- [Pengembangan](#-pengembangan)
- [Deploy ke Produksi](#-deploy-ke-produksi)
- [Kontribusi](#-kontribusi)
- [Lisensi](#-lisensi)
- [Kontak](#-kontak)

## 🌎 Environment

```sh
// see env.example
DBNAME=<YOUR_DB_NAME>
DBUSER=<YOUR_DB_USER>
DBHOST=<YOUR_DB_HOST>
DBPORT=<YOUR_DB_PORT>
DBPASS=<YOUR_DB_PASS>

JWT_SECRET=<YOUR_JWT_SECRET>
JWT_ISSUER=<YOUR_JWT_ISSUER>

RDSHOST=<YOUR_REDIS_HOST>
RDSPORT=<YOUR_REDIS_PORT>
```

## 🌟 Fitur

- Operasi CRUD untuk data Movie
- Autentikasi JWT
- Caching dengan Redis
- Dokumentasi API otomatis dengan Swagger
- Dukungan Docker
- Paginasi data
- Pencarian film
- Validasi request

## 📋 Persyaratan

- Go versi 1.20 atau lebih baru
- PostgreSQL 17+
- Redis 7+

## 🛠 Instalasi

### Menggunakan Docker (Direkomendasikan)

```bash
# Clone repository
git clone https://github.com/username/movie-api.git
cd movie-api

# Salin file env
cp .env.example .env

# Edit konfigurasi
nano .env

# Jalankan container
docker-compose up -d --build
```
