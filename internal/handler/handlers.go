package handler

import (
	"fmt"
	"net/http"
	"roadmaps/projects/personal-blog/internal/model"
	"text/template"
	"time"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // URL.path is /edit/pageName. However, we re-slice this path to drop the /edit/ and get only page name/title.
	page, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	RenderTemplate(w, "view", page)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := model.LoadPage(title)
	if err != nil {
		page = &model.Article{
			ID:        "",
			Title:     title,
			Content:   nil,
			CreatedAt: "",
			UpdatedAt: "",
		}
	}
	RenderTemplate(w, "edit", page)
}

func RenderTemplate(w http.ResponseWriter, templateName string, page interface{}) {
	tpl, _ := template.ParseFiles("internal/tmpl/templates/" + templateName + ".html")
	err := tpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	if title == "" {
		title = r.FormValue("title")
		body := r.FormValue("body")
		articleID := time.Now().Format("200601021504105")
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := time.Now().Format("2006-01-02 15:04:05")

		article := &model.Article{
			ID:        articleID,
			Title:     title,
			Content:   []byte(body),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		err := article.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
		return
	}

	body := r.FormValue("body")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	existingArticle, err := model.LoadPage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	article := &model.Article{
		ID:        existingArticle.ID,
		Title:     title,
		Content:   []byte(body),
		CreatedAt: existingArticle.CreatedAt,
		UpdatedAt: updatedAt,
	}
	err = article.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	return
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Generate the list of all articles
	articles, err := model.GenerateListOfAllArticles()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating list of articles: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the template
	tpl, err := template.ParseFiles("internal/tmpl/templates/home.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading template: %v", err), http.StatusInternalServerError)
		return
	}

	// Execute the template with the articles slice
	err = tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Generate the list of all articles
	articles, err := model.GenerateListOfAllArticles()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating list of articles: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the template
	tpl, err := template.ParseFiles("internal/tmpl/templates/admin.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading template: %v", err), http.StatusInternalServerError)
		return
	}

	// Execute the template with the articles slice
	err = tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	page := model.Article{}
	RenderTemplate(w, "new", page)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization required"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
