data-streams
============

A backend service for storing arbitrary data streams into Cassandra. Install instructions below.

installing go
-------------

    $ rm -rf /usr/local/go
    $ sudo tar -C /usr/local -xzf ~/Downloads/go1.1.2.linux-amd64.tar.gz
    $ mkdir -p ~/go
    $ export GOROOT=/usr/local/go
    $ export GOPATH=~/go

ide setup
---------

* installing code completion: `$ go get -u github.com/nsf/gocode`
* using eclipse plugin `goclipse`; install site: http://goclipse.googlecode.com/svn/trunk/goclipse-update-site/ 
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

installing web framework: martini
---------------------------------

    $ go get github.com/codegangsta/martini
