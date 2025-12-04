# Catatan Alur Pembuatan Auth JWT dengan Golang

Dokumen ini adalah panduan mental model untuk Anda. Jika nanti Anda diminta membuat aplikasi serupa dari nol, ikuti langkah-langkah logis ini.

## 1. Persiapan Pondasi (Infrastructure Layer)
Sebelum masuk ke logika bisnis (coding fitur), siapkan dulu "wadah" dan alat-alatnya.

*   **Config (`internal/config`)**:
    *   **Kenapa?** Aplikasi butuh settingan (DB user, password, JWT secret) yang tidak boleh di-hardcode.
    *   **Apa yang dilakukan?** Buat struct `Config`, load file `.env`.
*   **Database (`internal/database`)**:
    *   **Kenapa?** Kita butuh koneksi ke tempat penyimpanan data.
    *   **Apa yang dilakukan?** Buat fungsi `InitDB` yang membuka koneksi ke MySQL menggunakan config tadi.

## 2. Pola Pikir Clean Architecture (Layering)
Saat membuat fitur (misal: Auth atau User), jangan tulis semua kode di satu file. Pecah menjadi 3 lapisan utama. Hafalkan urutan ini:

**Repository → Service → Handler**

### A. Repository (Urusan Database)
*   **Tugasnya**: Hanya boleh ngobrol sama Database (SQL). Tidak boleh ada logika bisnis disini.
*   **Contoh**: "Simpan user ini ke tabel", "Cari user dengan email X".
*   **File**: `repository.go`

### B. Service (Urusan Bisnis Logic)
*   **Tugasnya**: Otak dari aplikasi. Menggabungkan data dari Repository dan alat lain (seperti Hashing atau JWT).
*   **Contoh**:
    *   *Register*: Cek email duplikat (panggil Repo) → Hash password (panggil Utils) → Simpan user (panggil Repo).
    *   *Login*: Cari user (panggil Repo) → Cek password cocok gak? (panggil Utils) → Bikin Token (panggil JWT Utils).
*   **File**: `service.go`

### C. Handler (Urusan HTTP/Web)
*   **Tugasnya**: Pintu gerbang. Menerima request dari user (JSON), validasi input dikit, lempar ke Service, lalu balikin response JSON.
*   **Contoh**: "Ada request Login nih, email & passwordnya ada gak di body? Oke ada, tolong Service proses dong. Hasilnya apa? Oke sukses, nih user saya kasih JSON 200 OK".
*   **File**: `handler.go`

---

## 3. Alur Implementasi Fitur Auth

Saat mengerjakan fitur Auth, kerjakan dengan urutan ini agar tidak bingung:

1.  **Model & DTO**: Tentukan dulu bentuk datanya.
    *   `User` (Model database): ID, Name, Email, Password.
    *   `LoginRequest` (DTO input): Email, Password.
2.  **Repository**: Bikin fungsi `FindByEmail` dan `Save`.
3.  **Utilities (Bantuan)**:
    *   **Hashing**: Bikin fungsi `HashPassword` dan `CheckPasswordHash` (pakai bcrypt). Password gak boleh plain text!
    *   **Token**: Bikin fungsi `GenerateToken` dan `ValidateToken` (pakai jwt-go).
4.  **Service**: Gabungkan semuanya.
    *   *Login*: Repo cari user -> Cek Hash -> Generate Token.
5.  **Handler**: Bikin endpoint HTTP-nya.

## 4. Middleware (Satpam)
Setelah fitur Login jadi, Anda butuh "Satpam" untuk menjaga halaman rahasia (misal: Profile).

*   **Logika Satpam**:
    1.  Cegat setiap request.
    2.  Cek header `Authorization`. Ada isinya gak?
    3.  Kalau ada, validasi tokennya (panggil fungsi `ValidateToken` tadi). Asli gak tanda tangannya? Belum expired kan?
    4.  Kalau aman, ambil `user_id` dari token, tempel ke `Context` (saku request), lalu persilakan masuk.
    5.  Kalau gak aman, tendang (Return 401 Unauthorized).

## 5. Wiring (Menyambungkan Kabel)
Terakhir, di `main.go`, sambungkan semua komponen yang sudah dibuat tadi. Ingat **Dependency Injection**:

1.  Buka Koneksi DB.
2.  Init **Repository** (butuh DB).
3.  Init **Service** (butuh Repository).
4.  Init **Handler** (butuh Service).
5.  Init **Router**, daftarkan URL (`/login`, `/register`) ke Handler yang sesuai.
6.  Jalankan Server.

---

## Ringkasan Cheat Sheet

| Komponen | Pertanyaan Mental |
| :--- | :--- |
| **Config** | "Settingan aplikasi ngambil dari mana?" |
| **Database** | "Konek ke MySQL gimana?" |
| **Repository** | "SQL Query-nya apa?" |
| **Service** | "Logika bisnisnya gimana? (Hash pass, Cek duplikat, Bikin token)" |
| **Handler** | "Baca JSON input gimana? Balikin JSON response gimana?" |
| **Middleware** | "Siapa yang boleh masuk sini? Cek tokennya dulu." |
| **Main** | "Rakit semua komponen jadi satu." |
