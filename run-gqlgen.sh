#!/bin/bash

echo 'Running generate schema'
go get github.com/99designs/gqlgen@v0.17.24 && go get github.com/99designs/gqlgen/codegen@v0.17.24 && go get github.com/99designs/gqlgen/internal/imports@v0.17.24 && go get github.com/99designs/gqlgen/codegen/config@v0.17.24

go run github.com/99designs/gqlgen
echo 'Graphql generate'