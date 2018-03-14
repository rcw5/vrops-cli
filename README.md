# vRops CLI

A CLI to interact with the vRealize Operations Manager SDK.

## Installation

To install:

```
% go get -u github.com/rcw5/vrops-cli
```

## Usage

To see full command help, visit the [USAGE](USAGE.md) page.

### Authentication

All commands need the URL of, and credentials for, your vRops instance. These can be set as environment variables or using command line arguments:

```
% export VROPS_URL=https://vrops.example.org
% export VROPS_USERNAME=admin
% export VROPS_PASSWORD=SuperScretPassword
% vrops-cli getadapterkinds

+---------------------------------+--------------------------------+--------------------------------+-----------------+
|               KEY               |              NAME              |          DESCRIPTION           | ADAPTERKINDTYPE |
+---------------------------------+--------------------------------+--------------------------------+-----------------+
| vRealizeOpsMgrAPI               | vRealizeOpsMgrAPI              | vRealizeOpsMgrAPI              | OPENAPI         |
| VirtualAndPhysicalSANAdapter    | vSAN Adapter                   | vSAN Adapter                   | GENERAL         |
+---------------------------------+--------------------------------+--------------------------------+-----------------+
```

```
% vrops-cli --target https://vrops.example.org --username admin --password SuperSecretPassword get adapterkinds
...
```

### Output

All `get` commands support two output types: `table` and `json`.

`table` output shows a subset of data whereas `json` shows everything. Pipe the `json` output into `jq` for further filtering, if you like.

### Debugging

Use the `--trace` flag to enable verbose output, where all HTTP request/responses are printed to `stdout`

```
% vrops-cli --trace get adapterkinds
GET /suite-api/api/adapterkinds HTTP/1.1
Host: 192.168.0.32
Accept: application/json
Authorization: Basic BasicAuthHeaderHere=


HTTP/1.1 200 200
Transfer-Encoding: chunked
Access-Control-Allow-Origin: *
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Content-Type: application/json;charset=UTF-8
Date: Mon, 12 Mar 2018 20:23:24 GMT
Expires: 0
Pragma: no-cache
Server: Apache
Strict-Transport-Security: max-age=31536000 ; includeSubDomains
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-Request-Id: mv36CSpkInxWlIuJnbFVdXnkVed6tJju
X-Xss-Protection: 1; mode=block
...
```
