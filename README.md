teapot
======

Package teapot implements HTCPCP-TEA.

### documentation:
* https://godoc.org/github.com/davsk/teapot

### reference:
* https://www.ietf.org/rfc/rfc2324.txt
* http://www.rfc-editor.org/rfc/rfc7168.txt
* http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
* https://en.wikipedia.org/wiki/Hyper_Text_Coffee_Pot_Control_Protocol#Protocol

### usage:
```GO
import (
  "net/http"
  "github.com/davsk.net/teapot"
)

init (
	http.HandleFunc("/teacup", teapot.Handler)
)
```
