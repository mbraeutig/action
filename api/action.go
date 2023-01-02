package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const URL = "https://www.actionforhappiness.de/wp-content/uploads/"

func Action(w http.ResponseWriter, r *http.Request) {
	_, month, day := time.Now().Date()
	// fmt.Println(year, month, day)      // For example 2009 November 10
	// fmt.Println(year, int(month), day) // For example 2009 11 10

	// https://www.actionforhappiness.de/wp-content/uploads/2022/01/AFH_ACTION_2022_02-DE.jpg

	var b strings.Builder
	b.WriteString(URL)
	b.WriteString("2022")
	b.WriteString("/")
	var strMonth = strconv.Itoa(int(month))
	if int(month) < 10 {
		b.WriteString("0")
	}
	b.WriteString(strMonth)
	b.WriteString("/AFH_ACTION_")
	b.WriteString("2022")
	b.WriteString("_")
	var strDay = strconv.Itoa(day)
	if int(day) < 10 {
		b.WriteString("0")
	}
	b.WriteString(strDay)
	b.WriteString("-DE.jpg")
	//fmt.Print(b.String())

	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(b.String())

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
		http.Error(w, fmt.Sprintf("ioutil.ReadAll -> %v", err), http.StatusInternalServerError)
	}

	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	res.Body.Close()
	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, bytes.NewReader(data)); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/jpg")
}
