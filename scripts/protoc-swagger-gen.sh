#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen

cd proto

proto_dirs=$(find ./xfury -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yml $query_file
  fi
done

cd ..

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./docs/swagger/config.json -o ./docs/swagger/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen

# set -eo pipefail

# mkdir -p ./docs/client
# proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
# for dir in $proto_dirs; do

#   # generate swagger files (filter query files)
#   query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
#   if [[ ! -z "$query_file" ]]; then
#     buf protoc \
#       -I "proto" \
#       -I "third_party/proto" \
#       "$query_file" \
#       --swagger_out=./docs/client \
#       --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
#   fi
# done
