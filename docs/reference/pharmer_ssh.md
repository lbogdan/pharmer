## pharmer ssh

SSH into a Kubernetes cluster instance

### Synopsis


SSH into a cluster instance.

```
pharmer ssh [flags]
```

### Examples

```
appctl cluster ssh -c cluster-name node-name
```

### Options

```
  -c, --cluster string   Cluster name
  -h, --help             help for ssh
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [pharmer](pharmer.md)	 - Pharmer by Appscode - Manages farms
