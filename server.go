package main

import (
	"fmt"
	"github.com/acsellers/site13/assets"
	"github.com/acsellers/site13/models"
	"net/http"
	"regexp"
)

var PageId = regexp.MustCompile("/page/(.*)")
var CatId = regexp.MustCompile("/cat/(.*)")

func main() {
	SetupAssets()
	StaticPages()
	DynamicPages()

	http.ListenAndServe(os.Getenv("PORT"), nil)
}

func SetupAssets() {
	NewCss("normalize", assets.NORMALIZE_CSS)
	NewCss("bundle", assets.BUNDLE_CSS)
	NewJS("bundle", assets.ZEPTO_JS, assets.BUNDLE_JS)
	CloseAssets()
}

func DynamicPages() {
	DynamicPage("/page/", BlogPage)
	DynamicPage("/cat/", CategoryPage)
	DynamicPage("/", BlogIndex)
}

func StaticPages() {
	StaticPage("/about", ABOUT)
}

func BlogIndex(w http.ResponseWriter, req *http.Request) {
	RenderBlogIndex(w, req)
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
	if page, found := models.FindPage(getPageId(req)); found {
		RenderBlogPage(page, w)
	} else {
		RenderBlogIndex(w, req)
	}
}

func CategoryPage(w http.ResponseWriter, req *http.Request) {
	if category := getCategory(req); category != "" {
		RenderCategory(category, w)
	} else {
		fmt.Println("Couldn't find category")
		NotFoundPage(w, req)
	}

}

func getPageId(req *http.Request) string {
	if PageId.MatchString(req.URL.Path) {
		return PageId.FindStringSubmatch(req.URL.Path)[1]
	}
	return ""
}

func getCategory(req *http.Request) string {
	if CatId.MatchString(req.URL.Path) {
		return CatId.FindStringSubmatch(req.URL.Path)[1]
	}
	return ""
}

func NotFoundPage(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
}
