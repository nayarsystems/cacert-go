# Up-to-date root CA bundle for Golang

This package provides a simple way to embed an up to date root CA bundle in your Go application

The Github Action workflow pulls the latest CA bundle from the `cacert` package in the nixpkgs repository every day at `0 0 * * *`, and if there is any change it will create a commit and a tag, compatible with Golang's module versioning system