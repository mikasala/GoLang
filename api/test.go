package api

// In C#: using "fmt"
import (
	"fmt"
	"net/http"
	"strings"
)

func TestAPI(w http.ResponseWriter, r *http.Request) {
	action := strings.Trim(strings.TrimPrefix(r.URL.Path, "/test"), "/")
	switch action {
	case "one":
		fmt.Println(action + "1")
		break
	case "two":
		fmt.Println(action + "2")
		break
	default:
		break
	}
}
