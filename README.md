# kbkb controller for Chaos Engineering
[![Go Report Card](https://goreportcard.com/badge/github.com/omakeno/kbkb-controller)](https://goreportcard.com/report/github.com/omakeno/kbkb-controller)

I know you're tempted to delete 4 or more pods of the same color stuck together. If your pods have annotation of "kbkb.k8s.omakenoyouna.net/color: red/blue/green/yellow/purple", this operator does just that.

## Install

```bash
kubectl apply -f https://raw.githubusercontent.com/omakeno/kbkb-controller/master/deploy/deploy.yaml
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
* If Kbkb object exists in a workspace, this controller start to delete pods.
