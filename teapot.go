// teapot.go

// Package teapot HTCPCP-TEA implementation.
//
// by David Skinner.
//
// reference:
//  https://www.ietf.org/rfc/rfc2324.txt
//  http://www.rfc-editor.org/rfc/rfc7168.txt
//  http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
package teapot

import (
	//	"bitbucket.org/ww/goautoneg"
	"fmt"
	"net/http"
)

// Handler reports status error for inappropriate HTCPCP Request.
func Handler(rw http.ResponseWriter, req *http.Request) {
	// Implement  HTTP/1.1 §14.4 Support.
	//	myLang := goautoneg.Negotiate(req.Header.Get("Accept-Language"), Langs)

	// Implement HTCPCP/1.0 §2.3.2
	if req.Header.Get("Content-type") != "message/teapot" {
		rw.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(rw, "<h1>%s</h1><div>%s</div><e>I am short and stout.<span lang=zh-Hans>短而粗壮</span></e>", req.URL.Host, http.StatusText(http.StatusTeapot))
	}
}

/*
the Accept-Additions header field defined in the base HTCPCP specification is updated to allow the following options:
      addition-type   = ( "*"
                        | milk-type
                        | syrup-type
                        | sweetener-type
                        | spice-type
                        | alcohol-type
                        | sugar-type
                        ) *( ";" parameter )
      sugar-type      = ( "Sugar" | "Xylitol" | "Stevia" )
*/

// BUG(skinner.david@gmail.com) Excessive use of the Sugar addition may cause the BREW request to exceed the segment size allowed by the transport layer, causing fragmentation and a delay in brewing.

var TeaBags = []string{
	"/assam",
	"/ceylon-tea",
	"/darjeeling",
	"/earl-grey",
	"/lapsang-souchong",
	"/peppermint",
	"/yunnan",
}

var Langs = map[string]string{
	"af":  "Ek is kort en geset.",
	"sq":  "Unë jam i shkurtër dhe i shëndoshë.",
	"ar":  "أنا قصير وقوي البنية.",
	"hy":  "Ես կարճ եւ Գեր.",
	"az":  "Mən qısa və qalın edirəm.",
	"eu":  "Naiz labur eta Stout.",
	"be":  "Я коратка і тоўсты.",
	"bn":  "আমি ছোট এবং স্থুলকায় না.",
	"bs":  "Ja sam kratko i stout.",
	"bg":  "Аз съм кратък и як.",
	"ca":  "Jo sóc baix i robust.",
	"ceb": "Ako mao mubo ug magahi.",
	"zh":  "我是短而粗壮。",
	"hr":  "Ja sam kratko i krupna.",
	"cs":  "Já jsem krátký a tlustý.",
	"da":  "Jeg er kort og stout.",
	"nl":  "Ik ben kort en stout.",
	"en":  "I am short and stout.",
}

// TODO Update the stakeholder requirements to resolve requirements that cannot be realized or are impractical to achieve.
