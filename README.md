# geneos-REST-server
REST server written in go to service Geneos components

This currently serves up connection files for the ActiveConsole and netprobe configurations to self-announcing probes. This was originally coded to work against a kubernetes cluster which generates the console connection file(s) and spawns the SAN probes. 

## SAN-probe 

/netprobe/:netprobe

Uses the template file assets/netprobe/probe.xml.tmpl to generate the configuration.

## Active Console

/activeconsole/<connection_file>

Simply serves up a connection file from assets/activeconsole. Right now this is generated outside this program but could be incorporated. 


