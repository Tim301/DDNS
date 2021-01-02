Installation:
  go get github.com/mattn/go-sqlite3
  go get -u github.com/gorilla/mux
  go build

Database:
  Db store key value as follow:
    location string, accesskey string, updatekey string, ip string
  Where location is a nickname of your system,
	accesskey is a password for allowing redirection,
	updatekey is a password for allowing redirection,
	ip is the url of redirection.
  Use DB Browser (SQlite) to edit database.

How to use it?
  go to http://127.0.0.1/access/{location}/{accesskey} to be redirected.
  go to http://127.0.0.1/update/{location}/{updatekey} to update redirection ip.
  accesskey and updatekey must be different to prevent access user to update redirection ip.
  


