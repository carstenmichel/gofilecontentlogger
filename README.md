# gofilecontentlogger

A service that scans one directory for files and writes logs into another directory.

To configure use the following Env Vars:

- `PREFIX_CHECKDIR` configures which directory is monitored
- `PREFIX_LOGDIR` configures the name of the logfile

# Helm

Also contains a helm chart to deploy this into your already existing namespace

`helm install <appname> dirlogger`


