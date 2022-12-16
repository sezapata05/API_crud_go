package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct{}

type BooksParams struct {
	Offset int `schema:"offset"`
	Limit  int `schema:"limit"`
}

type PostBook struct {
	Title string `schema:"title"`
}

var resp = make(map[string]string)

var (
	books   = []string{"Book 1", "Book 2", "Book 3"}
	decoder = schema.NewDecoder()
)

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	// resp := make(map[string]string)

	params := &BooksParams{}

	err := decoder.Decode(params, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp["status"] = http.StatusText(http.StatusInternalServerError)
		resp["message"] = "Parameters error"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	// limitParam := r.URL.Query().Get("limit")

	// if limitParam == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	resp["status"] = http.StatusText(http.StatusBadRequest)
	// 	resp["message"] = "Limit is mandatory"
	// 	jsonResp, err := json.Marshal(resp)
	// 	if err != nil {
	// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// 	}
	// 	w.Write(jsonResp)
	// 	return
	// }

	// limit, err := strconv.Atoi(limitParam)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp["status"] = http.StatusText(http.StatusInternalServerError)
	// 	resp["message"] = fmt.Sprintf("%s is not a valid parameter", limitParam)
	// 	jsonResp, err := json.Marshal(resp)
	// 	if err != nil {
	// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// 	}
	// 	w.Write(jsonResp)
	// 	return
	// }

	if params.Offset > len(books) || params.Offset < 0 {
		resp["status"] = http.StatusText(http.StatusBadRequest)
		resp["message"] = "OffSet parameter is bigger that books len"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	if params.Limit < 0 || params.Limit > len(books) {
		resp["status"] = http.StatusText(http.StatusBadRequest)
		resp["message"] = "Limit parameter is bigger that books len"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}

	booksJson, err := json.Marshal(books[from:to])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(booksJson)

	// json.NewEncoder(w).Encode(books)
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request) {
	pathParamrs := mux.Vars(r)

	idParam := pathParamrs["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (id-1) < 0 || (id-1) > len(books)-1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books[id-1])

}

func (a *API) postBook(w http.ResponseWriter, r *http.Request) {
	book := &PostBook{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	books = append(books, book.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(books)
}
