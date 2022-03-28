package main

import (
	"fmt"
	"github.com/mfpierre/go-mcrypt"
)

//func index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello, gsjddd!dddd")
//}
//
//func f1(w http.ResponseWriter, r *http.Request)  {
//	str := "服务端第一次测试"
//	w.Write([]byte(str))
//}

func main() {

	key := []byte("19781160")
	plaintext := []byte("14505000")
	iv := make([]byte, 8)

	// using CAST-256 in ECB mode
	encrypted, _ := mcrypt.Encrypt(key, iv, plaintext, "des", "ecb")
	decrypted, _ := mcrypt.Decrypt(key, iv, encrypted, "des", "ecb")
	fmt.Println(string(encrypted))
	fmt.Println(string(decrypted))

	//http.HandleFunc("/test/gsj" , f1)
	//http.ListenAndServe("0.0.0.0: 121141", nil)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", index)
	//
	//server := &http.Server{
	//	Handler: mux,
	//	Addr:    ":8080",
	//}

	//log.Fatal(server.ListenAndServe())
}
