A basic CLI to interact with the vRealize Operations Manager SDK

```
NAME:
   vrops-cli - CLI to interact with VMware vRealize Operations Manager

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     get      Retrieve data from vROps
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username value, -u value  vRops Username [$VROPS_USERNAME]
   --password value, -p value  vRops Password [$VROPS_PASSWORD]
   --url value                 vRops URL [$VROPS_URL]
   --help, -h                  show help
   --version, -v               print the version
```

```
NAME:
   vrops-cli get - Retrieve data from vROps

USAGE:
   vrops-cli get command [command options] [arguments...]

COMMANDS:
     adapterkinds   get all adapterkinds
     resourcekinds  get all resourcekinds for an adapter

OPTIONS:
   --help, -h  show help
```
