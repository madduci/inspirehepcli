# inspirehepcli

A CLI Client to look for publications in [InspireHEP](https://inspirehep.net) programmatically and convert them in Markdown or HTML.

### How it works

- Retrieves information from InspireHEP by Literature ID or ArXiv ID
- Extracts basic information from the JSON Payload (authors' list, publication name, journal name)
- Converts the basic information to a Markdown/HTML file saved on disk

### How to use it

To retrieve publication information, simply execute the command:

```inspirehepcli [<flags>] <id>```

where <id> is the Publication ID to look for in InspireHEP, so a file named `<id>.html` will be created on disk, in the same folder where the CLI client is being executed.

The following <flags> are allowed:

* -o, --output="html"  The desired output type.
* -a, --arxiv          Uses the Arxiv ID for the search (default).
* -l, --literature     Uses the Literature ID for the search.
*     --help           Show context-sensitive help (also try --help-long and --help-man).
*     --version        Show application version.

### Build from Source

It requires Go 1.15 and can be built with the command:

`go build` 

on your platform, or you can use the environment variables **GOOS** and **GOARCH** to perform a cross-compilation to the desired target platform of your choice.

### Library Used

- [Resty v2](https://github.com/go-resty/resty)
- [Kingpin v2](gopkg.in/alecthomas/kingpin.v2)
