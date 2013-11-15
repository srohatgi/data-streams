data-streams
============

A backend service for storing arbitrary data streams into Cassandra. The 

installing go
-------------

* `sudo add-apt-repository ppa:gophers/go`
* `sudo apt-get update`
* `sudo apt-get install golang-stable`

setup go
--------

* `mkdir -p ~/go`
* `export GOPATH=~/go`

installing gocode
-----------------

* `go get -u github.com/nsf/gocode`

installing goclipse
-------------------

* `go env` provides all environment variables

installing cassandra
--------------------

data stax enterprise build

    $ echo "deb http://username:password@debian.datastax.com/enterprise stable main" | \
      sudo tee -a /etc/apt/sources.list.d/datastax.sources.list
    $ curl -L https://debian.datastax.com/debian/repo_key | sudo apt-key add -
    $ sudo apt-get update
    $ sudo apt-get install dse-full opscenter ## Installs both DataStax Enterprise and OpsCenter.
    
install go driver
-----------------

