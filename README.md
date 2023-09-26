# broken-image-reg

GOAL: An simple image registry that enables testing 'failure' scenarios

## Developing

```sh
cd <somedir>

# clone broken-image-reg and fork of go-containeregistry into same base directory
# (the current broken-image-reg go.mod uses this structure for pulling in go-containerregistry)
git clone https://github.com/dcaravel/broken-image-reg.git
git clone https://github.com/dcaravel/go-containerregistry.git

cd broken-image-reg
go run .
```

Layer contents will be stored at `$HOME/broken-reg` (by default)

Refer to [internal/env/env.go](internal/env/env.go) for list of env vars