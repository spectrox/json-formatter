# JSON Formatter

This project is meant to be used in the pipeline with tools like `kubectl`.

The goals of this projects are simple:

1. Pretty print valid JSON strings received from `kubectl` or any other tool.
1. Skip non-valid JSON strings or strings which doesn't contain JSON at all.
1. Join multi-line JSON messages which were split by `php-fpm` or any other tool adding limits for a line length.

## Build

Project can be easily built with this command:

```bash
go build
```

## Run

After `go build`, you can use this tool in the pipeline:

```bash
kubectl -n ... logs ... | <path-to-project>/json-formatter
```

Add alias or shortcut if needed.
