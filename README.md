GoAhead
=================

GoAhead is a simple service that redirects routes (in development)
### Dependencies 
golang >= 1.1   
code.google.com/p/gcfg


### Usage
Copy the etc/goahead.ini.sample to etc/goahead.ini, edit it to use your database and persons informations.  
1. Compile gohead with `go build`.  
2. Run it with `goahead` 

#### FastCGI (optional)
At first enable fastcgi for true in goahead.ini. So, write the rule at /etc/nginx/sites-enabled/redir-go.dev.conf:

```
server {
    listen 80;
    server_name redir-go.dev ;
    index index.html ;

    location ~ / {
        include /etc/nginx/fastcgi_params;
        fastcgi_pass 127.0.0.1:9000;
    }
}

```
Run goahead normally! :) 

### Maintainers

Claudson Oliveira - [github.com/cloudson](http://github.com/cloudson)  
Vinicius Oliveira - [github.com/viniciuswebdev](http://github.com/viniciuswebdev)