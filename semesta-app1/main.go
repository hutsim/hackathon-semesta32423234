package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, `<body style="background-color:#000000">`)
		fmt.Fprint(w, `<center><h1 style="color:#FFFFFF">Selamat datang di Semesta System Administrator</center></h1>`)
		fmt.Fprint(w, `<center><img src="https://maukuliah.id/assets/img/semesta/logo-semesta-light.png" alt="Gambar"></center>`)
		fmt.Fprint(w, `</body>`)
	} else if r.URL.Path == "/aboutus" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Fprintf(w, "<center><h1>Gagal memuat file .env, silahkan cek kembali</h1></center>")
			fmt.Println("Gagal memuat file .env:", err)
			return
		}
		targetURL := os.Getenv("APP2_URL")
		if targetURL == "" {
			fmt.Fprintf(w, "<center><h1>URL Web App ke 2 tidak ditemukan, silahkan cek file .env</h1></center>")
			fmt.Fprintf(w, `<center><img src="https://edlink.id/assets/img/404.gif" alt="err"></center>`)
			fmt.Println("URL tujuan tidak ditentukan")
			return
		}
		resp, err := http.Get(targetURL)
		if err != nil {
			fmt.Fprintf(w, "<center><h1>Halaman yang dicari tidak ditemukan, Silahkan cek kembali url Web App ke 2</h1></center>")
			fmt.Fprint(w, "<center>Gagal memuat konten dari ", targetURL, ": ", err.Error(), "</center>")
			fmt.Fprintf(w, `<center><img src="https://edlink.id/assets/img/404.gif" alt="err"></center>`)
			fmt.Println("Gagal memuat konten:", err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			http.Error(w, "Gagal membaca respons", http.StatusInternalServerError)
			fmt.Println("Gagal membaca respons:", err)
			return
		}
		fmt.Fprintf(w, "%s", body)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1><center>Halaman yang dicari tidak ditemukan</center></h1>")
	}
}

func main() {
	server := &http.Server{
		Addr:              ":3000",
		Handler:           http.HandlerFunc(handlerFunc),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println(run())
}
func run() string {
	return "Setup Travis CI for Golang SEMESTA Hackathon"
}
