data-streams
============

A backend service for storing arbitrary data streams into Cassandra. Install instructions below.

installing go
-------------

    $ sudo add-apt-repository ppa:gophers/go
    $ sudo apt-get update
    $ sudo apt-get install golang-stable

setup go
--------

    $ mkdir -p ~/go
    $ export GOPATH=~/go

ide setup
---------

* installing code completion: `$ go get -u github.com/nsf/gocode`
* using eclipse plugin `goclipse`, install site for: http://goclipse.googlecode.com/svn/trunk/goclipse-update-site/ 
* configure goclipse using values from `go env`

installing cassandra
--------------------

data stax enterprise build

    $ echo "deb http://username:password@debian.datastax.com/enterprise stable main" | \
      sudo tee -a /etc/apt/sources.list.d/datastax.sources.list
    $ curl -L https://debian.datastax.com/debian/repo_key | sudo apt-key add -
    $ sudo apt-get update
    $ sudo apt-get install dse-full opscenter ## Installs both DataStax Enterprise and OpsCenter.
    
install cassandra go driver
---------------------------

    $ go get github.com/pomack/thrift4go/lib/go/src/thrift
    $ go get github.com/araddon/cass

