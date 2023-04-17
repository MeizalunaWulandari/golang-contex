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
 */