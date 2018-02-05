# namespace-deleter

## Overview

When a [Kubernetes][k8s] namespace is no longer needed it can be deleted. This
application will delete the namespace it is configured to delete after a signal
is given.

This is used by [Sonobuoy Scanner][scanner] to clean up the `heptio-sonobuoy`
namespace after the [Sonobuoy][sonobuoy] run is complete.

## Configuration

This application only needs one environment variable to work and two to work
well.

`NAMESPACES` defines which namespaces to delete. Namespaces should be a comma
delimited string: `"example-ns-1,example-ns-2"`

`READ_RESULTS_DIR` tells the application to read this directory for a file named
`done`. This file is used as the signal used to indicate the namespace is ready
for deletion.

## Example

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: deleter
  namespace: example-ns-1,example-ns-2
spec:
  containers:
  - env:
    - name: NAMESPACES
      value: "example-ns-1,example-ns-2"
    image: gcr.io/heptio-images/namespace-deleter:v0.0.2
    imagePullPolicy: Always
    name: deleter
  serviceAccountName: example-serviceaccount
```

This pod will delete the namespaces `example-ns-1` and `example-ns-2` when it sees a file named `done` in
`/tmp/results`.

[k8s]: https://kubernetes.io/
[scanner]: https://scanner.heptio.com/
[sonobuoy]: https://github.com/heptio/sonobuoy/
