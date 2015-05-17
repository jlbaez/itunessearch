# Itunes search
#About
A package to make itunes searchs in Go.
#Installation
Install the package to your local machine:
```
go get github.com/jlbaez/itunessearch
```
#Usage
Just call:
```go
itunessearch.Search(query string, mediatype string, limit int)
```
It will return an array of structs of type SearchResult.
So far the structs have a limited amount of information but please feel free to add more as you see fit.
