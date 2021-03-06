package server

import (
	"log"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/buro9/microcosm/web/controllers"
	mm "github.com/buro9/microcosm/web/middleware"
	"github.com/buro9/microcosm/web/opts"
)

// ListenAndServe will run the web server
func ListenAndServe() chan error {
	router := chi.NewRouter()

	// Pages group, handles all routes for pages and defines the appropriate
	// middleware for web pages
	router.Group(func(router chi.Router) {
		router.Use(mm.RealIP)
		router.Use(middleware.RequestID)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(mm.APIRoot)
		router.Use(mm.ForceSSL)
		router.Use(mm.Session)

		router.Get("/", controllers.HomeGet)
		router.Get("/auth0login/", controllers.Auth0LoginGet)
		router.Get("/conversations/{conversationID:[1-9][0-9]+}/", controllers.ConversationGet)
		router.Get("/microcosms/{microcosmID:[1-9][0-9]+}/", controllers.MicrocosmGet)
		router.Get("/profiles/{profileID:[1-9][0-9]+}/", controllers.ProfileGet)
		router.Get("/profiles/", controllers.ProfilesGet)
		router.Get("/today/", controllers.TodayGet)
		router.Get("/updates/", controllers.UpdatesGet)

		router.Post("/logout/", controllers.LogoutPost)

		router.NotFound(controllers.NotFound)

	})

	// Static file group, defines minimal middleware
	router.Group(func(router chi.Router) {
		// TODO: Log the static, disabled during dev
		router.Use(mm.RealIP)
		//router.Use(middleware.RequestID)
		router.Use(middleware.Logger)
		// router.Use(middleware.Recoverer)
		// router.Use(apiRoot)
		router.Use(mm.ForceSSL)

		router.Mount("/static", staticFiles())
		// TODO: clear these stubs
		ok := func(w http.ResponseWriter, req *http.Request) { w.Write([]byte("OK")) }

		router.Get("/isogram", ok)
		router.Get("/favicon.ico", ok)
		router.Get("/robots.txt", ok)

		router.NotFound(controllers.NotFoundStatic)
	})

	// This is the microcosm client and can work over http as well as https,
	// whilst we'll handle redirecting all *.apidomain.tld to https and likewise
	// for any *Site.ForceSSL to https... we cannot do it for every site as some
	// will be CNAMEd to us and we do not have the certs for their
	// customdomain.tld
	//
	// This means that we serve *everything* over both http and https and we
	// use the forceSSL middleware to use SSL where needed.
	//
	// The by-product of this long-winded explanation is that we listen for both
	// standard http and TLS connections

	// Channel for returning any error out of either of the http or https
	// listeners
	errs := make(chan error)

	go func() {
		log.Printf("Listening for HTTPS on %s ...", *opts.TLSListen)
		err := http.ListenAndServeTLS(
			*opts.TLSListen,
			*opts.CertFile,
			*opts.KeyFile,
			router,
		)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		log.Printf("Listening for HTTP on %s ...", *opts.Listen)
		err := http.ListenAndServe(
			*opts.Listen,
			router,
		)
		if err != nil {
			errs <- err
		}
	}()

	return errs
}

func staticFiles() http.Handler {
	router := chi.NewRouter()

	var css = regexp.MustCompile("\\.css$")
	var gif = regexp.MustCompile("\\.gif$")
	var js = regexp.MustCompile("\\.js$")
	var png = regexp.MustCompile("\\.png$")

	// Do nothing, but implement http.Handler
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, router *http.Request) {
			switch {
			case css.MatchString(router.RequestURI):
				w.Header().Set("Content-Type", "text/css")
			case gif.MatchString(router.RequestURI):
				w.Header().Set("Content-Type", "image/gif")
			case js.MatchString(router.RequestURI):
				w.Header().Set("Content-Type", "text/javascript")
			case png.MatchString(router.RequestURI):
				w.Header().Set("Content-Type", "image/png")
			}
			next.ServeHTTP(w, router)
		})
	})

	// Serve static files
	router.Mount("/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir(*opts.FilesPath+"/static/")),
		),
	)

	return router
}
