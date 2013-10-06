GoAhead
=================

GoAhead is a simple service that redirects routes (in development)
### Dependencies 
golang >= 1.1   


### Usage
Copy the etc/goahead.ini.sample to etc/goahead.ini, edit it to use your database and persons informations.     
Get dependencies with `go get` at the goahead folder.  
Compile gohead with `go build`.

#### Do you not have the redirect table? (optional)
Use our scaffold command running `goahead -build`. 

#### Do you need the FastCGI support? (optional)
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

#### Go ahead!!!
Run it with `goahead -run`

#### Questions? 
`goahead -help` 


### Maintainers

Claudson Oliveira - [github.com/cloudson](http://github.com/cloudson)  
Vinicius Oliveira - [github.com/viniciuswebdev](http://github.com/viniciuswebdev)
