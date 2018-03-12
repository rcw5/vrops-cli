# Usage

## `create`

```
NAME:
   vrops-cli create - Create objects in vRops

USAGE:
   vrops-cli create command [command options] [arguments...]

COMMANDS:
     stats  Create new stats in vRops

OPTIONS:
   --help, -h  show help
```

### `create stats`

```
NAME:
   vrops-cli create stats - Create new stats in vRops

USAGE:
   vrops-cli create stats [command options] [arguments...]

DESCRIPTION:
   Upload one or more stats for a specific resource

OPTIONS:
   --stats-json value
```

Example `--stats-json`:

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
NAME:
   vrops-cli get - Retrieve data from vROps

USAGE:
   vrops-cli get command [command options] [arguments...]

COMMANDS:
     adapterkinds   get all adapterkinds
     resourcekinds  get all resourcekinds for an adapterkind
     resources      get all resources for an adapterkind

OPTIONS:
   --help, -h  show help
```

### `get adapterkinds`

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
NAME:
   vrops-cli get resourcekinds - get all resourcekinds for a given adapterkind

USAGE:
   vrops-cli get resourcekinds [command options] [arguments...]

DESCRIPTION:
   Returns all resourcekinds configured in vRops
   An resourcekind is the class of entities that represent objects or information sources

OPTIONS:
   --output value, -o value  Set an output format: json or table (default: "table")
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
NAME:
   vrops-cli get resources - get all resources for an adapterkind

USAGE:
   vrops-cli get resources [command options] [arguments...]

DESCRIPTION:
   Returns all resources for a specific adapterkind. Use the returned identifier to create stats.

OPTIONS:
   --output value, -o value  Set an output format: json or table (default: "table")
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
