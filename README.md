# Go-Social-Feed-API
=================
RESTful API to manage social feeds written in Go and uses MongoDB as storage


# Fetching Dependencies
=======================
Get the dependent packages, we need to setup the API:

go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux

toml :  Parse the configuration file (MongoDB server & credentials)
mux : Request router and dispatcher for matching incoming requests to their respective handler
mgo : MongoDB driver

The code creates a controller for each endpoint, then expose an HTTP server on port 3000.

# API structure
===========
dao/
        feeds_dao.go
config/
        config.go
feed.go
config.toml


#API
======
1. getNewFeeds
    Parameters Expected :
        user_id
        page    // for pagination of 10 records
    This api gives the update of the people user follows. It first fetches the list of people (followTo) whom user follows
    and then get the feeds for those users and respond using http.ResponseWriter in json format.

2. search
    Parameters Expected :
        keyword
        page    // for pagination of 10 records
    This api search various data like, posts, People (first_name or last_name of users), places (found in address field of users).
    and then respond using http.ResponseWriter in json format.
