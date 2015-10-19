package main

// This is my version of the DB persistance with Go article here: http://www.alexedwards.net/blog/organising-database-access
// Note that in the models, I had to use the sql.Null* types, because of the
// flakiness of the underlying data. There is probably a better way to do this.
import (
  "github.com/kryptykfysh/data_access/models"
  "fmt"
  "log"
  "net/http"
)

type Env struct {
  db models.Datastore
}

func main() {
  db, err := models.NewDB("postgres://username:password@host/database?sslmode=disable")
  if err != nil {
    log.Panic(err)
  }

  env := &Env{db}

  http.HandleFunc("/accounts", env.accountsIndex)
  http.ListenAndServe(":3000", nil)
}

func (env *Env) accountsIndex(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }
  accs, err := env.db.AllAccounts()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  for _, acc := range accs {
    fmt.Fprintf(w, "%s, %s, %s", acc.CompanyName, acc.ABN, acc.ACN)
  }
}
