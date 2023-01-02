# RBAC Proof of concept

This PoC is a part of the [RBAC proposal](../../proposal/2022-12-23-rbac.md). It shows how Botkube executor (in this case, `kubectl`) can use

## Instruction

1. Set up a Kubernetes cluster, e.g. with `k3d` or `colima`.
1. Deploy Botkube with the following values:

```yaml
```

1. Apply



```bash
k apply -f ./docs/investigation/rbac/assets
```

Test getting pods in the `botkube-demo` channel:

```
@Preview get po
```

```
@Preview get po -n botkube
```

Test getting services:

```
@Preview get svc
```

This won't work:
```
@Preview get svc -n botkube
```
Same with other resources:

```
@Preview get ingress
```
