# JSON Formatter

This project is meant to be used in the pipeline with tools like `kubectl`.

The goal of this projects is simple:

1. Format valid JSON strings received from `kubectl` or any other tool.
1. Skip non-valid json strings or strings which doesn't contain JSON.
1. Join multi-line log messages which aren't meant to be split by formatter, but were split by `php-fpm` or any other tool adding limits for a line length.

## Build

Project can be easily built with this command:

```bash
go build
```

## Run

When projects is built you can add this command to the pipeline:

```bash
kubectl -n ... logs ... | <path-to-project>/json-formatter
```

Add alias or shortcut if needed.
