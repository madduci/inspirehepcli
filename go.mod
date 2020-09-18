module github.com/madduci/inspirehepcli

go 1.15

require (
	github.com/madduci/inspirehepcli/ihclient v1.0.0
	github.com/madduci/inspirehepcli/ihconverter v1.0.0
)

replace (
	github.com/madduci/inspirehepcli/ihclient => ./ihclient
	github.com/madduci/inspirehepcli/ihconverter => ./ihconverter
)
