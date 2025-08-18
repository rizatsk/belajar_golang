# Learner GO Database MySql
Materi GO Context
https://docs.google.com/presentation/d/15pvN3L3HTgA9aIMNkm03PzzIwlff0WDE6hOWWut9pg8/edit?slide=id.p#slide=id.p

## GO-LANG
### Pengenalan
Go (Golang), Belajar database untuk koneksi dan query ke database yang digunakan. dapat gunakan driver yang disediakan oleh GO untuk koneksi database yang diinginkan https://golang.org/s/sqldrivers

### Prepare Statement
Untuk prepare mengganti parameter, Misal ada ribuan data yang harus di-insert, tapi hanya parameternya (values) yang berubah.
→ PREPARE query sekali, lalu EXECUTE berkali-kali. Ini hemat CPU karena DB nggak perlu parse query ulang.

### Copy From
Untuk query insert lebih banyak, dengan logic seperti ini:
- Go kirim 1 perintah COPY users ... FROM STDIN ke PostgreSQL.
- PostgreSQL buka "pipeline" (kayak satu koneksi khusus).
- Go stream-kan semua row ke pipeline itu tanpa parsing SQL ulang per row.
- PostgreSQL langsung masukkan data ke tabel dengan sekali jalan.

