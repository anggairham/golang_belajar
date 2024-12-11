#

- Golang bukan OOP

## Interface kosong
tipe bool, int, string itu secara tidak langsung itu implementasi dari interface kosong
ada banyak contoh interface kosong di Go-Lang seperti :
fmt.Println(a ...interface{})
panic(v interface{})
recover()interface{}

## Nil
- biasanya di bahasa java, object belum diinisialisasi maka secara otomatis nilainya null atau nil
- berbeda dengan Go-lang, saat kita buat variabel dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
  contoh : jika bool false, int 0, string ""
- Namun di golang ada data nil, yaitu data kosong 
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice,pointer,dan channel

## error Interface
- golang memeiliki interface sebagai kontrak untuk membuat error, nama interfacenya adalah error
- int, error