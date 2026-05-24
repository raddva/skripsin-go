# Sistem Informasi Inventaris Dokumen Skripsi (SkripsIn)

## Deskripsi Proyek
Sistem Informasi Inventaris Dokumen Skripsi (SkripsIn) adalah aplikasi berbasis CLI (Command Line Interface) yang dirancang untuk mengelola arsip tugas akhir mahasiswa secara digital. Aplikasi ini memungkinkan staf administrasi program studi atau petugas perpustakaan untuk melakukan manajemen data judul skripsi, penulis, dan tahun kelulusan secara efisien.

Sistem ini dikembangkan menggunakan bahasa pemrograman Go (Golang) dengan fokus pada kemudahan pengelolaan data inventaris dokumen akademik.

## Identitas Kelompok
| Nama | NIM | Kelas |
| :--- | :--- | :--- |
| Nadya Auradiva | 103012500282 | IF-49-12 |
| Rafil Junior | 103012530012 | IF-49-12 |

## Spesifikasi dan Fitur Sistem
Berdasarkan kebutuhan fungsional, sistem ini mencakup poin-poin berikut:
- Manajemen Data (CRUD): Pengguna dapat menambah, mengubah, dan menghapus data dokumen skripsi secara dinamis.  
- Detail Informasi: Pencatatan informasi mendalam mengenai pembimbing, topik penelitian, dan status kelulusan mahasiswa.  
- Pencarian Data (Searching): Fitur pencarian data skripsi berdasarkan nama mahasiswa atau judul penelitian menggunakan algoritma Sequential Search dan Binary Search.  
- Pengurutan Data (Sorting): Fitur mengurutkan data skripsi berdasarkan nama penulis atau tahun lulus menggunakan algoritma Selection Sort dan Insertion Sort.  
- Statistik: Sistem mampu menampilkan statistik jumlah judul skripsi per tahun serta total dokumen yang tersimpan dalam database

## Struktur Data Utama
Aplikasi mengelola beberapa entitas data penting, antara lain:
- Informasi Skripsi: Judul penelitian atau topik penelitian.
- Informasi Penulis: Nama mahasiswa dan status kelulusan.
- Akademik: Nama pembimbing dan tahun lulus.

## Cara Menjalankan Aplikasi
1. Pastikan Anda telah menginstal Go di perangkat Anda.
2. Clone atau unduh repositori ini.
3. Buka terminal di direktori proyek.
4. Jalankan perintah berikut:
```bash
go run main.go
```

## Peranan Pengguna
* Staf Administrasi: Bertanggung jawab mengelola input data mahasiswa dan status kelulusan.  
* Petugas Perpustakaan: Mengelola inventaris judul penelitian dan melakukan pencarian referensi untuk kebutuhan akademik.

___

_Dokumen ini disusun sebagai bagian dari pemenuhan tugas mata kuliah Algoritma dan Pemrograman._