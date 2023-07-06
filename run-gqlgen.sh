#!/bin/bash

echo 'Running generate schema'
go get github.com/99designs/gqlgen && go get github.com/99designs/gqlgen/codegen && go get github.com/99designs/gqlgen/internal/imports && go get github.com/99designs/gqlgen/codegen/config

go run github.com/99designs/gqlgen
echo 'Graphql generate'