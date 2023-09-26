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

## Using

```sh
# with registry running as described above
$ HOST=<host>

$ docker pull $HOST/broken/error:404
Error response from daemon: manifest for host/broken/error:404 not found: manifest unknown: Looking for something?

$ docker pull $HOST/broken/error:403
Error response from daemon: unknown: Fakereg: Access to this resource is forbidden

# Docker daemon will perform both a HEAD and GET request regardless if is error with first request
# which is why the timeout takes 'double' whats specified.
❯ time docker pull $HOST/broken/timeout:dur-2s
Error response from daemon: unknown: Timeout duration has elapsed!

real   0m4.380s
user   0m0.011s
sys    0m0.022sj

$ time docker pull $HOST/broken/timeout:dur-5s
Error response from daemon: unknown: Timeout duration has elapsed!

real   0m10.393s
user   0m0.024s
sys    0m0.017s
```