package web

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/shurcooL/httpfs/filter"
	handlers "github.com/tmthrgd/httphandlers"
	"go.tmthrgd.dev/spork/internal/assets"
)

func MountAssets(r chi.Router) {
	fs := http.FileServer(filter.Skip(assets.FileSystem, excludeAssets))

	r.With(handlers.SetHeaderWrap("Cache-Control", "public, max-age=1209600" /* 14 days*/)).
		Get("/favicon.ico", fs.ServeHTTP)

	r.Get("/robots.txt", handlers.ServeString("robots.txt", time.Now(), "User-agent: *\nDisallow: /").ServeHTTP)

	r.With(handlers.SetHeaderWrap("Cache-Control", "public, no-cache")).
		Mount("/assets", http.StripPrefix("/assets", fs))
}

func excludeAssets(path string, info os.FileInfo) bool {
	return info.IsDir() || strings.HasPrefix(info.Name(), ".")
}

func assetPath(name string) string {
	return path.Join("/assets", name)
}
