package web

import (
	"github.com/campadrenalin/spiro_identicons/art"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/render/", render)
}

func render(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers["Content-Type"] = []string{"image/png"}

	readInt := func(pname string) (int, bool) {
		i, err := strconv.Atoi(r.FormValue(pname))
		if err != nil {
			return 0, false
		} else {
			return i, true
		}
	}

	seed := r.URL.Path[len("/render/"):]
	ar := art.NewRequest(seed)

	if i, ok := readInt("size"); ok {
		ar.SetSize(i)
	}

	ar.RenderPNG(w)
}
