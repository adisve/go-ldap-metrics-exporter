# Introduction
go-ldap-metrics-exporter is a tool for exporting prometheus metrics from an LDAP server. It is written in Go and works on Linux and Windows.

# Download and installation

## Download
go-ldap-metrics-exporter can be downloaded from github:

```bash
git clone https://github.com/adisve/go-ldap-metrics-exporter.git
```

## Usage

The go-ldap-metrics-exporter executable can be run from any directory. It is recommended to place the executable in the ´/opt/go-ldap-metrics-exporter/´ directory, along with its configuration file. The configuration flie is a .json file, and can be named anything, but its path must be provided to the executable using the `-c/--config` flag.

Create a user and group for the go-ldap-metrics-exporter service:
```bash
sudo groupadd go-ldap-metrics-exporter
sudo useradd -s /sbin/nologin -r -M -d /opt/go-ldap-metrics-exporter -g go-ldap-metrics-exporter go-ldap-metrics-exporter
```

And give the user ownership of the go-ldap-metrics-exporter directory:
```bash
sudo chown -R go-ldap-metrics-exporter:go-ldap-metrics-exporter /opt/go-ldap-metrics-exporter
```

The go-ldap-metrics-exporter service can then be started using the service file provided in the repository. The service file is named `go-ldap-metrics-exporter.service`.


## Configuration

`go-ldap-exporter` expects a .json file, and can be named anything, but its path must be provided to the executable using the `-c/--config` flag.

An example configuration could look like this

```json
{
  "ldap": {
    "address": "ldap://localhost", // could also be ldaps://, will use port 389 or 636 respectively
    "username": "<service-user-username>",
    "password": "<service-user-password>",
    "baseDn": "dc=nonp,dc=dp,dc=dig,dc=nchosting,dc=dk",
  },
  "scrape": {
    "interval": 30 // interval for the server, not the exporter
  },
  "server": {
    "active": true,
    "address": "localhost",
    "port": "9496"
  },
  "log": {
    "level": "info",
    "json": false
  },
  "export": {
    "file": "/var/lib/node_exporter/textfile_collector/metrics.prom",
    "interval": 300
  }
}