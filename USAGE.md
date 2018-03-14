# Usage

```
Wouldn't it be amazing if you didn't need to worry about the vROps REST API!?

Usage:
  vrops-cli [command]

Available Commands:
  create      Change stuff in vROps
  get         Retrieve stuff from vROps
  help        Help about any command

Flags:
  -h, --help              help for vrops-cli
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username

Use "vrops-cli [command] --help" for more information about a command.
```

## `create`

```
Change stuff in vROps

Usage:
  vrops-cli create [command]

Available Commands:
  stats       Create stats for a resource

Flags:
  -h, --help   help for create

Global Flags:
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username

Use "vrops-cli create [command] --help" for more information about a command.
```

### `create stats`

```
Create stats for a resource

Usage:
  vrops-cli create stats [resource] [flags]

Flags:
  -h, --help               help for stats
      --statsjson string   JSON-encoded stats to be uploaded

Global Flags:
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username
```

Example `--statsjson`:

```
[{
  "statKey": "cpu|usage",
  "timestamps": [
    1520883900000,1520883960000,1520884020000,1520884080000,1520884140000,1520884200000
  ],
  "data": [
    10,20,30,40,50,60
  ],
  "others": [],
  "otherAttributes": {}
}]
```
(Note: timestamps are in milliseconds)


## `get`

```
Retrieve stuff from vROps

Usage:
  vrops-cli get [command]

Available Commands:
  adapterkinds  Get all adapterkinds
  resourcekinds List resourceskinds for a given adapterkind
  resources     List resources for a given adapterkind

Flags:
  -h, --help            help for get
  -o, --output string   Output format (table or json) (default "table")

Global Flags:
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username

Use "vrops-cli get [command] --help" for more information about a command.
```

### `get adapterkinds`

```
Get all adapterkinds

Usage:
  vrops-cli get adapterkinds [flags]

Flags:
  -h, --help   help for adapterkinds

Global Flags:
  -o, --output string     Output format (table or json) (default "table")
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username
```

```
% vrops-cli get adapterkinds
+---------------------------------+--------------------------------+--------------------------------+-----------------+
|               KEY               |              NAME              |          DESCRIPTION           | ADAPTERKINDTYPE |
+---------------------------------+--------------------------------+--------------------------------+-----------------+
| Container                       | Container                      |                                | GENERAL         |
| EP Ops Adapter                  | EP Ops Adapter                 | EP Ops Adapter                 | GENERAL         |
| Http Post                       | Http Post                      |                                | OPENAPI         |
| OPENAPI                         | OPENAPI                        | OPENAPI                        | OPENAPI         |
| VMWARE                          | vCenter Adapter                | Provides the connection        | GENERAL         |
|                                 |                                | information and credentials    |                 |
|                                 |                                | required to monitor your       |                 |
|                                 |                                | vCenter Server instances       |                 |
| PythonRemediationVcenterAdapter | vCenter Python Actions Adapter | Provides actions for vCenter   | GENERAL         |
|                                 |                                | objects using Python scripts   |                 |
| vRBCAdapter                     | vRBC Adapter                   |                                | GENERAL         |
| VCACAdapter                     | vRealize Automation Adapter    |                                | GENERAL         |
| LogInsightAdapter               | vRealize Log Insight Adapter   |                                | GENERAL         |
| vCenter Operations Adapter      | vRealize Operations Adapter    | vRealize Operations Manager    | GENERAL         |
|                                 |                                | Adapter                        |                 |
| vRealizeOpsMgrAPI               | vRealizeOpsMgrAPI              | vRealizeOpsMgrAPI              | OPENAPI         |
| VirtualAndPhysicalSANAdapter    | vSAN Adapter                   | vSAN Adapter                   | GENERAL         |
+---------------------------------+--------------------------------+--------------------------------+-----------------+
```

### `get resourcekinds`

```
List resourceskinds for a given adapterkind

Usage:
  vrops-cli get resourcekinds [adapterkind] [flags]

Flags:
  -h, --help   help for resourcekinds

Global Flags:
  -o, --output string     Output format (table or json) (default "table")
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username
```

```
% vrops-cli get resourcekinds VCACAdapter
+--------------------------+
|           NAME           |
+--------------------------+
| BLUEPRINT                |
| BUSINESSGROUP            |
| DEPLOYMENT               |
| FABRICGROUP              |
| NETWORKPROFILE           |
| RESERVATION              |
| RESERVATIONPOLICY        |
| STORAGERESERVATIONPOLICY |
| TENANT                   |
| ENTITYSTATUS             |
| VCACAdapter Instance     |
+--------------------------+
```

### `get resources`

```
List resources for a given adapterkind

Usage:
  vrops-cli get resources [adapterkind] [flags]

Flags:
  -h, --help   help for resources

Global Flags:
  -o, --output string     Output format (table or json) (default "table")
  -p, --password string   vROps password
  -t, --target string     url to vROps instance
      --trace             enable request tracing
  -u, --username string   vROps username
```

```
% vrops-cli get resources "OPENAPI"
+--------------------------------+--------------------------------------+-------------+--------------------------+-------------+--------+
|              NAME              |              IDENTIFIER              | ADAPTERKIND |       RESOURCEKIND       | DESCRIPTION | HEALTH |
+--------------------------------+--------------------------------------+-------------+--------------------------+-------------+--------+
| OPENAPI (vRealize Operations   | c9b812b7-0764-4bde-b6a1-613c24c09207 | OPENAPI     | OPENAPI Adapter Instance |             | GREEN  |
| Manager Collector-vRealize     |                                      |             |                          |             |        |
| Cluster Node)                  |                                      |             |                          |             |        |
+--------------------------------+--------------------------------------+-------------+--------------------------+-------------+--------+
```
