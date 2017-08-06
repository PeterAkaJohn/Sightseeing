# Sightseeing-Application
Use multiple API to deliver an application to be used on the web or on smartphone

# Usage:

### MUST-DO:

* console:
  ```
  git clone
  cd into folder
  ```

### FRONT-END:

* In console:
  ```
  npm install
  webpack-dev-server
  ```

### DB-SETUP:

* have vagrant installed

* setup:
  ```
  cd into db-development
  vagrant up
  vagrant ssh
  sudo - su postgres psql
  \i sightseeing.sql
  ```

### BACK-END:

* have go installed

* either:
  ```
  cd into src
  go run main.go
  ```
  or
  ```
  cd into src
  go install
  ```

# API:

* Google Maps / google places / foursquare / yelp

# Language, framework, database used:

* Go (vanilla or buffalo found here http://gobuffalo.io/), gorillatoolkit

* Angular(1 or 2) or React

* mongodb or postgresql (use ORM like xorm found here http://xorm.io/)

* Ionic for a hybrid app to be used on the phone

# DONE:

* improved db tables and types

* bcrypt

* https server, http requests will be redirected to https

* Leaving routes like they were and only changed AddFavorite and Logout to implement the jwt middleware (will come back to it if I need to secure more routes to only be visibile by those authenticated)

* layouts and components refactored in presentational and container components, found a way to pass props to children in react-router v4

* react-router instruction: 1- Router as the outermost level, 2- Route that will get replaced when the path attribute match
* implement google map

# IMPLEMENTING:

* Start with front-end (React now, later angular 2 on a different branch) -- In Progress (realized how to use react-router v4.0, need to refactor code to reuse components, api calls in componentdidMount but don't know which component layout or need to refactor code so that App component implements everything, read redux documentation)

* style everything

* implement redux using react-redux redux-thunk redux-form (see example projects for reference, search immutability)
