[![CircleCI](https://circleci.com/gh/serainville/inquisitor/tree/master.svg?style=svg)](https://circleci.com/gh/serainville/inquisitor/tree/master) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/009c4a31594543fca3d36e2927420a35)](https://www.codacy.com/app/serainville/inquisitor?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=serainville/inquisitor&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/serainville/inquisitor)](https://goreportcard.com/report/github.com/serainville/inquisitor)

# Inquisitor
Inquisitor is a quick, light-weight system monitoring tool with plugin support. It has three
modes of operation: client, server, standalone. 

# Backend Storage
## Time series data
Time series data is stored in InfluxDB. There are plans to expand the backend storage by supporting other databases.

# Plugins
## Builtin
The following are plugins available by default.
* CPU
* Memory

## Third-Party
Third-party support can be added by creating exeutable files in a plugins directory. Inquisitor will scan the plugins directory during startup. The plugins will be executed during metric collection.

Plugins are expected to have three values, delimited by a comma: Name, group, value. So long as those are outputted, Inquisitor will store the data in the backend storage.

