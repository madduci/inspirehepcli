module github.com/madduci/inspirehepcli

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/madduci/inspirehepcli/ihclient v1.0.1
	github.com/madduci/inspirehepcli/ihconverter v1.0.1
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace (
	github.com/madduci/inspirehepcli/ihclient => ./ihclient
	github.com/madduci/inspirehepcli/ihconverter => ./ihconverter
)
