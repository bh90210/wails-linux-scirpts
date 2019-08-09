# τέστερ v0.2-alpha

a small go script for testing specific wails branches against supported Linux distributions

### use

`git clone https://github.com/bh90210/wails-linux-scirpts.git`

`cd wails-linux-scirpts`

`go run .`

#### available commands

```bash
. go run .
├── * `branch` it will promprt for distribution, git repo and specific branch to test against
│   ├── `go install && wails init`
│   ├── `go install && bin/bash /root`
│   └── `go install && wails init && bin/bash /root`
├── * `all` test (go install & wails init) of given git & branch against all supported distros (cpu intense!)
├── * `prune` delete wails built docker images from host (excluding distro builds)
├── * `pruneall` delete wails built docker images from host (including distro builds)
├── * `exit`
└── * `help` (meta)
```
