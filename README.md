# wails-linux-scripts v0.1-alpha

a small go script for testing wails against supported Linux distributions

### use

`git clone https://github.com/bh90210/wails-linux-scirpts.git`

`cd wails-linux-scirpts`

`go run .`

`./script-name.sh`

### available commands
1. test-branch $distro $git $branch (1) distribution to test on, git repo and specific branch to test against
2. supported-distros (7) show all currently support distributions
3. tester-prune (8) delete from host all wails built docker images

4. exit (0) exit the tester
5. help (9) (meta)
