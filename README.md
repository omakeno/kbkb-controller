# kbkb operator for Chaos Engineering
[![Go Report Card](https://goreportcard.com/badge/github.com/omakeno/kbkb-operator)](https://goreportcard.com/report/github.com/omakeno/kbkb-operator)

I know you're tempted to delete 4 or more pods of the same color stuck together. If your pods have annotation of "k8s.omakenoyouna.net/kubeColor: red/blue/green/yellow/purple", this operator does just that.

## Install

```bash
kubectl apply -f https://raw.githubusercontent.com/omakeno/kbkb-operator/master/deploy/deploy.yaml
```

## Usage

* Create kbkb object.

```yaml
apiVersion: k8s.omakenoyouna.net/v1beta1
kind: Kbkb
metadata:
  name: kbkb-four
spec:
  kokeshi: 4
```

* You can change the number of pods stuck to delete by editing "kokeshi".
* If Kbkb object exists in a workspace, this operator start to delete pods.
