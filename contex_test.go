package golang_context

import (
	"fmt"
	"context"
	"testing"
)

func TestContext(t *testing.T){
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextWithValue(t *testing.T){
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f")) 	// Dapat
	fmt.Println(contextF.Value("c"))	// Dapat Milik parent
	fmt.Println(contextF.Value("b"))	// Tidak dapat, beda parent
	fmt.Println(contextA.Value("b"))	// Tidak bisa mengambil data child
}

/**
 * PENGENALAN CONTEXT
 * Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout
 * dan sinyal deadline
 * Context biasanya dinuat per-request (misalnya setiap ada request masuk ke server web
 * melalui http request)
 * Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar 
 * proses(goroutine)
 * 
 * * Package Context
 * Context direpresentasikan di dalam sebuah interface Context
 * interface context terdapat dalam package context
 * 
 * * Membuat Context
 * Karena context adalah sebuah interface, untuk membuat context kita butuh 
 * sebuah struct yang sesuai dengan kontrak interface Context
 * Namum kita tidak perlu membuatnya secara manual karena di package context
 * terdapat function yang bisa digunakan untuk membuat context
 * ** Function Membuat context
 * *** context.Background()
 * digunakan untuk membuat context kosong, tidak pernah dibatalkan, tidak pernah
 * timeout, tidak memiliki value apapun. Biasanya digunakan di main function 
 * atau dalam test, atau dalam awal proses request terjadi
 * *** context.TODO()
 * digunakan untuk membuat context kosong seperti Background, namum biasanya 
 * menggunakan ini ketika belum jelas context apa yang ingin digunakan
 * 
 * PARENT DAN CHILD CONTEXT
 * Contetx menganut konsep parent dan child yang artinya pada saat kita membuat 
 * kita bisa membuat child context dari context yang sudah dibuat
 * Parent context dapat memiliki banyak child, namun child hanya bisa memiliki satu parent
 * Child akan memiliki semua fitur yang dibuat pada parent
 * * Hubungan Antara Parent dan Child Context
 * Parent dan Child akan selalu terhubung 
 * Misalnya ketika melakukan pembatalan Context A maka semua child dan sub-child dari context
 * A akan ikut dibatalkan
 * Namum jika membatalakan context B maka, hanya context B dan semua child dan sub-childnya 
 * yang dibatalkan sedangkan Parent (Context A) tidak ikut dibatalkan
 * * Immutable 
 * Context merupakan object yang Immutable, artinya setelah context dibuat, dia tidak bisa
 * dirubah lagi, Misalnya ketika menambahkan value atau menambahkan pengaturan timeout dan
 * yang lainnya, secara otomatis akan membentuk child context baru, bukan merubah context 
 * tersebut
 * 
 * CONTEXT WITH VALUE
 * Pada saat awal membuat context, context tidak memiliki value 
 * Kita bisa menambahkan sebuah value dengan data pair (key - value) ke dalam context
 * Saat kita menambahkan value context, secara otomatis akan tercipta child context baru
 * artinya original contextnya tidak akan berubah sama sekali
 * Untuk membuat menambahkan value ke context, kita bisa menggunakan function 
 * context.WithValue(parent, key, value)
 * *Context Get Value
 * Saat mengambil value sebuah context, maka child dapat mengambil value dari parentnya
 * namum parent tidak bisa mengambil value dari child
 */