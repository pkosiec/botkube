# RBAC Proof of concept

1. Deploy Botkube
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
