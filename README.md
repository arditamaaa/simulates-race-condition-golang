### Penjelasan
1. `sync.Mutex`: Objek mutex digunakan untuk mengamankan akses ke sumber daya bersama (saldo rekening).
2. `b.mutex.Lock()` dan `defer b.mutex.Unlock()`:
   - `Lock()` mengunci mutex sebelum mengakses sumber daya
   - `defer Unlock()` memastikan bahwa mutex dilepaskan setelah fungsi selesai dieksekusi.
4. `Deposit()` dan `Withdraw()`: Kedua fungsi ini menggunakan mutex untuk memastikan keamanan goroutine.
5. Penanganan Error: Fungsi `Withdraw()` mengembalikan error jika penarikan gagal (saldo tidak cukup).
6. Contoh Penggunaan: Contoh menunjukkan bagaimana beberapa goroutine dapat melakukan transaksi secara bersamaan tanpa menimbulkan kondisi persaingan.
7. `sync.WaitGroup`: WaitGroup digunakan untuk menunggu semua goroutine selesai sebelum program berakhir.

### Goroutine
Goroutine adalah salah satu fitur kunci yang membedakan Go (Golang) dari bahasa pemrograman lainnya. Mereka adalah unit konkuren eksekusi yang ringan dan memungkinkan untuk menangani ribuan proses secara bersamaan.

### Point
1. Gunakan defer untuk memastikan unlock selalu terpanggil. defer dijalankan setelah fungsi selesai.
2. Dengan implementasi ini, sistem rekening bank Anda akan aman untuk digunakan dalam lingkungan goroutine.
