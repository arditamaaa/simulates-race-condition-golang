### Penjelasan
1. `sync.Mutex`: Objek mutex digunakan untuk mengamankan akses ke sumber daya bersama (saldo rekening).
2. `b.mutex.Lock()` dan `defer b.mutex.Unlock()`: `Lock()` mengunci mutex sebelum mengakses sumber daya, dan `defer Unlock()` memastikan bahwa mutex dilepaskan setelah fungsi selesai dieksekusi.
3. `Deposit()` dan `Withdraw()`: Kedua fungsi ini menggunakan mutex untuk memastikan keamanan goroutine.
4. Penanganan Error: Fungsi `Withdraw()` mengembalikan error jika penarikan gagal (saldo tidak cukup).
5. Contoh Penggunaan: Contoh menunjukkan bagaimana beberapa goroutine dapat melakukan transaksi secara bersamaan tanpa menimbulkan kondisi persaingan.
6. `sync.WaitGroup`: WaitGroup digunakan untuk menunggu semua goroutine selesai sebelum program berakhir.
### Point Penting
1. Penggunaan mutex adalah cara umum untuk menangani kondisi persaingan dalam aplikasi goroutine.
2. Pastikan untuk selalu melepaskan mutex setelah selesai menggunakan sumber daya bersama.
3. Pertimbangkan untuk menggunakan mekanisme sinkronisasi yang lebih canggih jika aplikasi Anda memiliki kebutuhan yang lebih kompleks.
4. Gunakan defer untuk memastikan unlock selalu terpanggil.
5. Dengan implementasi ini, sistem rekening bank Anda akan aman untuk digunakan dalam lingkungan goroutine.
