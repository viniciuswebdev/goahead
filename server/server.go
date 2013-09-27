package server 

import (
    "net"
    "net/http"
    "net/http/fcgi"
    "log"
    "github.com/viniciuswebdev/goahead/database"
)

type Server struct {
    Port string
    Host string 
    FastCgi bool 
}

var _db *database.Database
var _table *database.TableConf

func (server *Server) initialize(db *database.Database, table *database.TableConf) {
    _db = db 
    _table = table
    if server.Port == "" {
        server.Port = "9000"
    }

    if server.Host == "" {
        server.Host = "localhost"
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    hash := r.URL.Path[1:]
    log.Printf("Searching url with hash '%s' \n", hash)

    url, error := _db.FindShortenedUrlByHash(hash, _table)
    if error != nil {
        log.Printf("%s \n", error.Error())
        http.NotFound(w, r)
        return
    }
    http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    handler(w, r)   
}

func (s *Server) TurnOn(db *database.Database, table *database.TableConf) {
    s.initialize(db, table)
    if s.FastCgi {
        log.Printf("Starting Goahead on %s:%s using fastcgi...\n", s.Host, s.Port)
        s.turnOnFastCGI()
        return 
    }
    log.Printf("Starting Goahead on %s:%s ...\n", s.Host, s.Port)
    s.turnOnSimple()
}

func (s *Server) turnOnSimple() {
    http.HandleFunc("/", handler)
    var err = http.ListenAndServe(":"+s.Port, nil)
    if err != nil {
        panic(err.Error())
    }
}

func (s *Server) turnOnFastCGI() {
    listener, err := net.Listen("tcp", s.Host+":"+s.Port)
    if err != nil {
        panic(err.Error())
    }
    srv := new(Server)
    fcgi.Serve(listener, srv)
}