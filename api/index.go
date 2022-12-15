package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
	<head>
	<html>
	Nothing to see here !!!		
	</html>
	</head>
	`

// Index - write the index
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
