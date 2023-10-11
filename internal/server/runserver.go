package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	asciiart "ascii-art-web/internal/ascii-art"
)

type Err struct {
	Text_err string
	Code_err int
}

type Output struct {
	Ascii_art, Text string
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		asciiArt := r.FormValue("res")

		if asciiArt == "" {
			IndexHandler(w, r)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
		w.Header().Set("Content-Type", "text/plain")

		w.Write([]byte(asciiArt))
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

}

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
	res := &Err{Text_err: http.StatusText(code), Code_err: code}
	err = tmpl.Execute(w, *res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	fs := r.FormValue("fs")

	isOK := false

	checkFS := []string{"standard", "shadow", "thinkertoy"}

	for _, val := range checkFS {
		if val == fs {
			isOK = true
			break
		}
	}

	if !isOK {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	if err := asciiart.AsciiChar(text); err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	} else {
		ascii_art, err := asciiart.Build(text, fs)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		out := Output{Ascii_art: ascii_art, Text: text}
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, out)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/ascii-art", AsciiArtHandler)
	mux.HandleFunc("/download", DownloadHandler)

	log.Println("Listening on: http://localhost:4000/")
	http.ListenAndServe(":4000", mux)
}
