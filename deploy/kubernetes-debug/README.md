# Debugging deployments

The deployments in this directory are intended to be used for debugging. 

## Usage

Deploy the node plugin:
```bash
kubectl apply -f csi-nodeplugin-sshfs-debug.yaml
```

When the pod is started it waits until a debugger connects to it before it will do anything.
It will wait for debugging connections on NodePort `31040` and the CSI Socket interface on NodePort `31010`.

Please refer your IDE's documentation for information about connecting.

- IntelliJ & Goland: https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/