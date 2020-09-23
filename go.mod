module github.com/madduci/inspirehepcli

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/madduci/inspirehepcli/ihclient v1.0.0
	github.com/madduci/inspirehepcli/ihconverter v1.0.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace (
	github.com/madduci/inspirehepcli/ihclient => ./ihclient
	github.com/madduci/inspirehepcli/ihconverter => ./ihconverter
)
