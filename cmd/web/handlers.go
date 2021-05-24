package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w) //use the notfound() helper
		return
	}

	//home.page must be first file int he slice

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// 	files := []string{
	// 		"./ui/html/home.page.tmpl",
	// 		"./ui/html/base.layout.tmpl",
	// 		"./ui/html/footer.partial.tmpl",
	// 	}

	// 	ts, err := template.ParseFiles(files...)
	// 	if err != nil {

	// 		app.serverError(w, err) //use serverError() helper
	// 		return

	// 	}

	// 	err = ts.Execute(w, nil)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 	}

	// }

	// func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// 	if err != nil || id < 1 {
	// 		app.notFound(w)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
	// 	//w.Write([]byte("Display a specific snippet..."))
	// }

	// func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method != http.MethodPost {

	// 		w.Header().Set("Allow", http.MethodPost)

	// 		//w.WriteHeader(405)
	// 		//w.Write([]byte("Method not alllowed"))
	// 		app.clientError(w, http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	title := " Stuff "
	// 	content := "Stuff and more stuff"
	// 	expires := "4"

	// 	id, err := app.snippets.Insert(title, content, expires)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}

	// 	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
