# inspirehepcli

A CLI Client to look for publications in [InspireHEP](https://inspirehep.net) programmatically and convert them in Markdown.

### How it works

- Retrieves information from InspireHEP by Literature ID
- Extracts basic information from the JSON Payload (authors' list, publication name, journal name)
- Converts the basic information to a Markdown file saved on disk

### How to use it

To retrieve publication information, simply execute the command:

```inspirehepcli [LiteratureId]```

and a file named `[LiteratureId].md` will be created on disk, in the same folder where the CLI client is stored.

### Build from Source

It requires Go 1.15 and can be built with the command:

`go build` 

on your platform, or you can use the environment variables **GOOS** and **GOARCH** to perform a cross-compilation to the desired target platform of your choice.

### Library Used

- [Resty v2](https://github.com/go-resty/resty)
