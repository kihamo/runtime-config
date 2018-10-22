package config

//go:generate /bin/bash -c "goimports -w `find . -type f -name '*.go' -not -path './vendor/*'`"
//go:generate /bin/bash -c "cd config && enumer -type=VariableType -trimprefix=VariableType -output=variable_enumer.go -transform=snake"
