package httpadapter

import (
    "fmt"
    "net/http"
)

// HealthHandler responde com "ok".
func HealthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ok")
}
