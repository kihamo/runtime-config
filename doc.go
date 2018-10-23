package config

//go:generate /bin/bash -c "goimports -w `find . -type f -name '*.go' -not -path './vendor/*'`"
//go:generate /bin/bash -c "cd config && enumer -type=VariableType -trimprefix=VariableType -output=variable_enumer.go -transform=snake"
//go:generate mockgen -destination=./client/mocks/mock_client.go -package=mocks github.com/kihamo/runtime-config/client Store
//go:generate mockgen -destination=./config/mocks/mock_config.go -package=mocks github.com/kihamo/runtime-config/config Value,Variable,Version
//go:generate mockgen -destination=./internal/store/mocks/mock_store.go -package=mocks github.com/kihamo/runtime-config/internal/store Store
//go:generate mockgen -destination=./manager/mocks/mock_manager.go -package=mocks github.com/kihamo/runtime-config/manager Store
