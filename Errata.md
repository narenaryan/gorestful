# Errata for _gorestful_

Corrections for the book [_Building RESTful Web Services with Go_](). The pages listed are for the released book itself, not for any preprints or other forms of the articles.

### First Printing, December 2017

## Page 162, Installing the PostgrSQL database.
#### sudo sh -c 'echo “deb ht<span>tp://apt.postgresql.org/pub/repos/apt/ \`lsb_release -cs\`-pgdg main” >> /etc/apt/sources.list.d/pgdg.list'

For `Ubuntu 18.04LTS` open the file `/etc/apt/sources.list.d/pgdg.list` as [root]() and if the first and
 last characters are double quotes, remove them to overcome error:

#### E: Type ‘“deb’ is not known on line 1 in source list /etc/apt/sources.list.d/pgdg.list

produced when doing:

#### sudo apt-get update

## Page 213, Go Kit, a package for building microservices.
#### `go get github.com/go-kit/kit`
You may get a response for this package saying :no Go files in $GOPATH/src/github.com/go-kit/kit
If you do, then try:
#### `go get github.com/go-kit/kit/...`
This gets all packages.

## Page 256, Installing a Kong database and Kong.
Earlier in the book, postgres was installed and may still be running as a service that will interfere with this chapter.

On Ubuntu, check what services are listening to ports with:

`sudo lsof -i -P -n | grep LISTEN`

If you see postgres in the list, stop it with:

`systemctl stop postgresql`

You may also need to stop `nginx`

The latest version of Kong (2.0.2 as of March 2020) does not work with what is detailed in the book, so we need to specify the specific version as used in the book and make a few other changes.
The details of what version of kong was used in the book is hidden in the details at the top of page 264, that is kong version 0.11.0
Use the following to get the **kong-database** to run:

`docker run -d --name kong-database -p 5432:5432 -e "POSTGRES_USER=kong" -e "POSTGRES_DB=kong" -e "POSTGRES_PASSWORD=kong" postgres:9.4`

And the line to do the migrations becomes:

`docker run --rm --link kong-database:kong-database -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" -e "KONG_PG_USER=kong" -e "KONG_PG_PASSWORD=kong" kong:0.11.0 kong migrations up`

## Page 258
The docker command sequence then needs to be:

`docker run -p 3000:3000 --name go-server -dit gobuild`

-

`docker run -d --name kong --link kong-database:kong-database --link go-server:go-server -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" -e "KONG_PG_PASSWORD=kong" -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" -e "KONG_PROXY_ERROR_LOG=/dev/stderr" -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" -p 8000:8000 -p 8443:8443 -p 127.0.0.1:8001:8001 -p 127.0.0.1:8444:8444 kong:0.11.0`

## Page 259
to see the output from `curl -X GET http://localhost:8001/apis/myapi` formatted similar to the book, install `jq` with (on Ubuntu):

`sudo apt-get install jq`

and then enter:

`curl -X GET http://localhost:8001/apis/myapi | jq`

## Page 261
From command line, use:

`curl -X GET http://localhost:8000/api/v1/healthcheck -H 'Host:server1'`

to get the time from the go server application running in its docker container.

## Page 266
If you experience an error when trying to do the `Kong Create API Key for Consumer`, then for a quick work around add a `Body` to the POST, thus:

{

    "key":"secret"
}

Then from the command line, use:

`curl -i -X GET http://localhost:8000/api/v1/healthcheck -H 'Host:server1' -H 'apikey:secret'`

## Page 277 to 279
`curl` commands to exercise the simpleAuth program:

`curl -c tmp.txt -v -X POST http://localhost:8000/login -d "username=admin&password=password"`

`curl -b tmp.txt -i -X GET http://localhost:8000/healthcheck`

-