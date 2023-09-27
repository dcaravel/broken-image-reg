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

## Using / API

### Error
Produces HTTP errors based on tag provided

&nbsp; | &nbsp;
--- | ---
Path | `/broken/error`
Tag  | `400` (bad request)<br>`403` (forbidden)<br>`404` (not found)<br> `405` (method not allowed)<br>`500` (internal server error)

#### Example
```sh
$ docker pull $REG_HOST/broken/error:404
Error response from daemon: manifest for REG_HOST/broken/error:404 not found: manifest unknown: Looking for something?

$ docker pull $REG_HOST/broken/error:403
Error response from daemon: unknown: Fakereg: Access to this resource is forbidden
```

### Timeout
Does not respond to the request until after the duration set by the tag has elapsed, at which point a 408 / request timeout error is returned.

Duration adheres to the syntax that can be parsed by [time.ParseDuration()](https://pkg.go.dev/time#ParseDuration)

&nbsp; | &nbsp;
--- | ---
Path | `/broken/timeout`
Tag  | `dur-XhYmZs`<br>ex: `dur-5s` (5 secs)

#### Example
```sh
# Docker daemon will perform both a HEAD and GET request regardless if is error with first request
# which is why the timeout takes 'double' whats specified.
❯ time docker pull $REG_HOST/broken/timeout:dur-2s
Error response from daemon: unknown: Timeout duration has elapsed!

real   0m4.380s
user   0m0.011s
sys    0m0.022sj

$ time docker pull $REG_HOST/broken/timeout:dur-5s
Error response from daemon: unknown: Timeout duration has elapsed!

real   0m10.393s
user   0m0.024s
sys    0m0.017s
```

### Random
Returns a random image with `X` layers of `Y` size in bytes.

`Z` defines the seed for randomness, assumes random seed `0` if not specified

Once the index / layers are generated, the blobs are cached so that pulls work

&nbsp; | &nbsp;
--- | ---
Path | `/broken/random`
Tag  | `layers-X_size-Y[_seed-Z]`<br>ex:<br>`layers-2_size-1024`<br>`layers-2_size-1024_seed-1`

#### Example
```sh
$ docker pull $REG_HOST/broken/random:layers-3_size-1024layers-3_size-1024: Pulling from broken/random
48b7e84b8a8d: Pull complete 
88e1c106d2b0: Pull complete 
c876421f51fb: Pull complete 
Digest: sha256:463a2d39a0585afdff7903f32a995b7e294b02d9cd9ca1e0a3454d2ab6346c73
Status: Downloaded newer image for REG_HOST/broken/random:layers-3_size-1024
REG_HOST/broken/random:layers-3_size-1024

$ docker pull $REG_HOST/broken/random:layers-3_size-1024layers-3_size-1024: Pulling from broken/random
Digest: sha256:463a2d39a0585afdff7903f32a995b7e294b02d9cd9ca1e0a3454d2ab6346c73
Status: Image is up to date for REG_HOST/broken/random:layers-3_size-1024
REG_HOST/broken/random:layers-3_size-1024
```


### Prepared
Responds with a manifest / manifest list that was previously uploaded via docker push

Intent is to be able to test handling different versions/variations of ManifestList's (oci vs docker), manifests, broken layers (invalid tars), etc. 

&nbsp; | &nbsp;
--- | ---
Path | `/broken/prepared`
Tag  | `<name>`

#### Example

```sh
$ docker push $REG_HOST/broken/prepared:roxctl
The push refers to repository [REG_HOST/broken/prepared]
8005c993c85c: Pushed 
79530facdb42: Pushed 
roxctl: digest: sha256:aa825e071722ee01173055f833c6e71c2babb8b0d25cba4c81eba3f90b09ea2e size: 739

$ docker pull $REG_HOST/broken/prepared:roxctl           
roxctl: Pulling from broken/prepared
0e6d0b7d441e: Pull complete 
2bff04353949: Pull complete 
Digest: sha256:aa825e071722ee01173055f833c6e71c2babb8b0d25cba4c81eba3f90b09ea2e
Status: Downloaded newer image for REG_HOST/broken/prepared:roxctl
REG_HOST/broken/prepared:roxctl
```

### Proxy
Proxies requests to another registry and caches locally

Will default to `quay.io` if no hostname found in the image path

Executed when none of the above APIs match the request

&nbsp; | &nbsp;
--- | ---
Path | `<any>`
Tag  | `<any>`

```sh
$ docker pull $REG_HOST/dcaravel/brokenreg:latest
latest: Pulling from dcaravel/brokenreg
1828023a8b9e: Pull complete 
Digest: sha256:dcb25b1ddc052c82ba825ee775a77e4a5947c5d7fd3e619f036a8d6cccccd1af
Status: Downloaded newer image for REG_HOST/dcaravel/brokenreg:latest
REG_HOST/dcaravel/brokenreg:latest

$ docker pull quay.io/dcaravel/brokenreg:latest
latest: Pulling from dcaravel/brokenreg
Digest: sha256:dcb25b1ddc052c82ba825ee775a77e4a5947c5d7fd3e619f036a8d6cccccd1af
Status: Downloaded newer image for quay.io/dcaravel/brokenreg:latest
quay.io/dcaravel/brokenreg:latest
```

```sh
$ docker pull nginx:1.24.0
1.24.0: Pulling from library/nginx
7dbc1adf280e: Pull complete 
a7184f3665ed: Pull complete 
f144d5d97503: Pull complete 
9097eea98b48: Pull complete 
356d4b647b64: Pull complete 
608e661a622a: Pull complete 
Digest: sha256:73341830a31bf12a44c846b6b323dd8a4fab7668e72c16e9124913ff097c9536
Status: Downloaded newer image for nginx:1.24.0
docker.io/library/nginx:1.24.0

$ docker pull $REG_HOST/docker.io/library/nginx:1.24.0
1.24.0: Pulling from nginx
Digest: sha256:5be2b646dfda41632549b19795721e3e676903c7d94567838fb1aa0e39ae1bfc
Status: Downloaded newer image for REG_HOST/docker.io/library/nginx:1.24.0
REG_HOST/docker.io/library/nginx:1.24.0

$ docker images
REPOSITORY                         TAG                  IMAGE ID       CREATED       SIZE
REG_HOST/docker.io/library/nginx   1.24.0               22c2ef579d56   7 days ago    142MB
nginx                              1.24.0               22c2ef579d56   7 days ago    142MB
```