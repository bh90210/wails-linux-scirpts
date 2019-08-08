# wails-linux-scripts v0.1-alpha

a small go script for testing specific wails branches against supported Linux distributions

### use

`git clone https://github.com/bh90210/wails-linux-scirpts.git`

`cd wails-linux-scirpts`

`go run .`

#### available commands
```bash
* `test-branch (1)` it will promprt for distribution to test on, git repo and specific branch to test against
- go install && wails init (1)
- go install && bin/bash /root (2) (opens a new terminal window)
- go install && wails init && bin/bash /root (3) (opens a new terminal window)
```
```bash
* `supported-distros (7)` show a list of all currently support distributions
```
```bash
* `tester-prune (8)` delete from host all wails built docker images
```
```bash
* `exit (0)` exit the tester
```
```bash
* `help (9)` (meta)
```
