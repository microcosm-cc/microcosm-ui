package funcs

import (
	"html/template"
	"strings"

	"github.com/buro9/funcs/general"
	"github.com/buro9/funcs/inspect"
	"github.com/buro9/funcs/math"
	"github.com/buro9/funcs/safe"
	"github.com/buro9/funcs/sanitize"
	"github.com/buro9/funcs/transform"
)

// Map is a `template.FuncMap` containing all of the funcs within the child
// packages as well as some from the core Go library
var Map template.FuncMap

func init() {
	Map = template.FuncMap{
		// General utilities
		"default": general.Default,

		// Transformation and generation
		"dict":        transform.Dictionary,
		"lower":       strings.ToLower,
		"naturalTime": transform.NaturalTime,
		"numcomma":    transform.NumComma,
		"ord2suffix":  transform.OrdToSuffix,
		"rfcTime":     transform.RFCTime,
		"title":       strings.Title,
		"trunc":       transform.Trunc,
		"upper":       strings.ToUpper,

		// Math functions
		"subtract": math.Subtract,

		// Trusting content
		"safeCSS":      safe.CSS,
		"safeHTML":     safe.HTML,
		"safeHTMLAttr": safe.HTMLAttr,
		"safeJS":       safe.JS,
		"safeURL":      safe.URL,

		// Untrusted content
		"stripTags":         sanitize.StripTags,
		"stripTagsSentence": sanitize.StripTagsSentence,
		"sanitize":          sanitize.Sanitize,

		// Inspect your data
		"contains": strings.Contains,
		"hasField": inspect.HasField,
		"in":       inspect.In,
		"isNil":    inspect.IsNil,
		"isSet":    inspect.IsSet,
	}
}
