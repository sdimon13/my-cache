module github.com/sdimon13/cache/store/bigcache/v1

go 1.22

require (
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/sdimon13/cache/lib/v4 v4.1.5
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.8.1
)

replace github.com/sdimon13/cache/lib/v4 => ../../lib/

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20230315142452-642cacee5cc0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
