# go-basicauth

[![Build Status](https://travis-ci.org/m90/go-basicauth.svg?branch=master)](https://travis-ci.org/m90/go-basicauth)
[![godoc](https://godoc.org/github.com/m90/go-basicauth?status.svg)](http://godoc.org/github.com/m90/go-basicauth)

> HTTP BasicAuth middleware

Package `basicauth` creates `func(http.Handler) http.Handler` middleware that checks for BasicAuth credentials

### Installation using go get

```sh
$ go get github.com/m90/go-basicauth
```

### Usage

Create a wrapping middleware function using `With(Credentials)`:

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("very exclusive!"))
})
middleware := basicauth.With(basicauth.Credentials{User: "me", Pass: "t0ps3c43t"})
handler = middleware(handler)
http.ListenAndServe("0.0.0.0:8080", handler)
```


### License
MIT © [Frederik Ring](http://www.frederikring.com)
