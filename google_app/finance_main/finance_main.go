package finance_main 

import (
    "appengine"
    "appengine/datastore"
    "appengine/user"
    "fmt"
    "http"
    "time"
)

type User struct {
    Name    string
    StartDate  datastore.Time
    Account string
}

func init() {
    http.HandleFunc("/api/register", registration)
}

func httphandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, r.RawURL)
}


func registration(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    u := User{
        Name: "TestHendrik",
        StartDate: datastore.SecondsToTime(time.Seconds()), 
   }

    if g := user.Current(c); g !=nil{
        var u2 User
        u.Account = g.String()
        if err := datastore.Get(c, datastore.NewKey("user", g.String(), 0, nil), &u2) ; err == datastore.ErrNoSuchEntity {
            key, err := datastore.Put(c, datastore.NewKey("user", u.Account,0, nil), &u)
    if err != nil {
         http.Error(w, err.String(), http.StatusInternalServerError)
         return
    }
    fmt.Fprintf(w, "User %q stored %q", u.Account, key)
        return
      } else {
           fmt.Fprintf(w, "User %q  is already logged in", g.String())
           return
}
    } else {
       url, err := user.LoginURL(c, r.URL.String())
       if err !=nil{
           http.Error(w, err.String(), http.StatusInternalServerError)
           return
       }
       w.Header().Set("Location", url)
       w.WriteHeader(http.StatusFound)
       return
    }
    
    }
