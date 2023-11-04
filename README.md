# Construct JSON without backing it up with Go structures
[![Go Reference](https://pkg.go.dev/badge/github.com/ridge/tj.svg)](https://pkg.go.dev/github.com/ridge/tj)

Use `tj` to produce well-formed schemaless JSON:

    json.Marshal(tj.O{
        "hello": tj.A{
            "world",
            42,
            tj.O{
                "go": nil,
            },
        },
    })

`tj/json2go` tool produces `tj` trees from existing JSON files.

## Why whould I need that?

This package comes in handy for interacting with large amounts of external
systems all demanding different JSON schemas. Creating structures, tagging them
and filling them is overkill if the structures are used just once.

## Why do I need this package to do so?

`tj.O` and `tj.A` are short and unobtrusive, this makes the resulting code
readable.

## Is it always a good idea to use this package?

If you find yourself using `tj` to describe JSONs with the same schema more than
once then it's time to switch to Go struct-backed JSON generation.

## Legal

Copyright Tectonic Labs Ltd.

Licensed under [Apache 2.0](LICENSE) license.

Authors:
- [Mikhail Gusarov](https://github.com/dottedmag)
- [Alexey Feldgendler](https://github.com/feldgendler)
