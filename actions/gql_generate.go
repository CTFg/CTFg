package main

//go:generate bash -c "pushd ../frontend; yarn graphql-codegen --config config/1_introspect.yml || true; popd; gqlgenc"
