## Some notes

### Start a http service
```go
http.ListenAndServe(":8000", yamlHandler)
// or 
http.ListenAndServe(":8000", mapHandler)
mapHandler := urlshort.MapHandler(pathsToUrls, mux)
```

1. `mapHandler` here needs to be a `http.handler` which is an interface with `ServeHTTP(ResponseWriter, *Request)`implemented.

2. `MapHandler` function is defined as such, it actually returns a `http.HandlerFunc`, and it's a **type** not a function!!

```go
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    // some stuff here
}
```

3. `http.HandlerFunc` is a function type, it has a `ServeHTTP` so it's a handler. As long as the type of the `func(ResponseWriter, *Request)`. In other words, it simply add ServeHTTP methods to the function. Yes, function can have methods in this case.
```go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
