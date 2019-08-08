package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	// I used to hate netflix, now I kind of love them!
	expect "github.com/Netflix/go-expect"
)

func main() {
	cli()
}

func cli() {
	// start a new keybard scanner
	scanner := bufio.NewScanner(os.Stdin)

	// and split input into words
	// this is for multiple word input
	// to be treated as flags of some sort
	scanner.Split(bufio.ScanWords)

	// helper viariable for the bellow switch
	// default value is help meaning this is the
	// landing case when cli starts
	var input = "help"
	// label for the 'for' loop
	// this is done so we can escpate later using 'brake loop'
loop:
	for {
		switch input {
		// when user enters a value not declared in any 'case'
		// app hangs. 'default' catches this and gracefully
		// returns some help and puts user back in the 'for' loop
		default:
			fmt.Println("command does not exist!")
			fmt.Println("if you need it type 'help'")
			scanner.Scan()
			text := scanner.Text()
			input = text

		case "test-branch", "1":
			// show a list of all supported-distros
			// because probably you don't remember it
			cmd := "cd ./supported-distros && ls"
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				fmt.Sprintf("failed to execute command: %s", cmd)
			}
			fmt.Printf("%s", out)

			fmt.Printf(" - enter: 'distro git branch'\n")
			// TODO: code cleanup
			scanner.Scan()
			distro := scanner.Text()
			scanner.Scan()
			git := scanner.Text()
			scanner.Scan()
			branch := scanner.Text()
			//fmt.Printf("%s %s %s\n", distro, git, branch)

			fmt.Println("go install && wails init (1)")
			fmt.Println("go install && bin/bash /root (2)")
			fmt.Println("go install && wails init && bin/bash /root (3)")
			scanner.Scan()
			text := scanner.Text()
			switch text {
			case "1":
				goInstallWailsInit(distro, git, branch)
			case "2":
				//goInstallBashRoot()
			case "3":
				//goInstallBoth()
			}

			// when finished return for user input
			fmt.Println("finished doing stuff!")
			fmt.Println("enter new command:")
			scanner.Scan()
			text = scanner.Text()
			input = text

		case "supported-distros", "7":
			cmd := "cd ./supported-distros && ls"
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				fmt.Sprintf("failed to execute command: %s", cmd)
			}
			fmt.Printf("%s\n", out)

			scanner.Scan()
			text := scanner.Text()
			input = text

			// exit
		case "exit", "0":
			fmt.Println("testing is over")
			fmt.Println("hopefully everything works :+1: :+1:")
			fmt.Println("see u soon!")
			break loop

		case "help", "9":
			fmt.Println("wails-linux-scripts v0.1-alpha helpfile!")
			fmt.Println("available commands:")
			fmt.Println(" * test-branch $distro $git $branch (1) distribution, git repo and specific branch to test against")
			fmt.Println(" * supported-distros (7) show all currently support distributions")
			fmt.Println(" * tester-prune (8) delete from host all wails built docker images")
			fmt.Println(" * exit (0) exit the tester")
			fmt.Println(" * help (9) (meta)")

			scanner.Scan()
			text := scanner.Text()
			//cmd := strings.TrimSuffix(text, "\n")
			input = text

		}
	}
}

func goInstallWailsInit(distro, git, branch string) {
	// check the image of selected distro exists
	// if not use Dockerfile to build it
	feedback := checkDockerImageExist(distro)
	fmt.Println(feedback)

	// build 'git-branch' container to produce the test build
	//cmd := "docker build -t wails-test-latest --build-arg GIT=https://github.com/wailsapp/wails.git --build-arg BRANCH=linux-db --no-cache ./git-branch"
	cmd := "docker build -t wails-test-latest --build-arg GIT=https://github.com/wailsapp/wails.git --build-arg BRANCH=linux-db ./git-branch"
	out := exec.Command("bash", "-c", cmd)
	out.Stdin = os.Stdin
	stdout, _ := out.StdoutPipe()
	b := bufio.NewScanner(stdout)
	err := out.Start()
	if err != nil {
		log.Println(err)
	}
	for b.Scan() {
		//print the input
		fmt.Println(b.Text())
	}

	// collect binary from inside the container to subdir '/test-branch'
	cmd2 := "docker run --rm --entrypoint '/bin/sh' -v $(pwd)/test-branch:/binary wails-test-latest -c ' cp /go/bin/wails /binary && cp -r /wails /binary/source'"
	out2 := exec.Command("bash", "-c", cmd2)
	out2.Stdin = os.Stdin
	stdout2, _ := out2.StdoutPipe()
	b = bufio.NewScanner(stdout2)
	err = out2.Start()
	if err != nil {
		log.Println(err)
	}
	for b.Scan() {
		//print the input
		fmt.Println(b.Text())
	}

	// TODO: remove wails-test-latest docker image

	// build selected distro test container and populate it with newly built 'wails' and './wails/wails.json'
	//docker run -it --rm --name wails-debian9-test  --entrypoint "/bin/bash" wails-debian9
	cmd3 := "cd test-branch && docker build -t wails-test-" + distro + " --build-arg DISTRO=" + distro + " --no-cache ."
	//cmd3 := "cd test-branch && docker build -t wails-test-" + distro + " --build-arg DISTRO=" + distro + " ."
	out3 := exec.Command("bash", "-c", cmd3)
	out3.Stdin = os.Stdin
	stdout3, _ := out3.StdoutPipe()
	b = bufio.NewScanner(stdout3)
	err = out3.Start()
	if err != nil {
		log.Println(err)
	}
	for b.Scan() {
		//print the input
		fmt.Println(b.Text())
	}

	//cmd4 := "docker run --rm --entrypoint '/bin/sh' -v $(pwd)/test-branch:/binary wails-test-latest -c 'cp /binary/wails /go/bin && cp -r /binary/.wails /'"
	//out4, _ := exec.Command("bash", "-c", cmd4).Output()
	//fmt.Println(out4)

	// run 'wails init'

	//exec.Command("bash", "-c", cmd5)
	//fmt.Println(out5)
	wailsInit(distro)
}

func checkDockerImageExist(distro string) string {
	cmd := "docker images | grep -c wails-" + distro
	out, _ := exec.Command("bash", "-c", cmd).Output()

	if string(out) == "0" {
		cmd := "docker build -t wails-" + distro + " /supported-distros/" + distro
		out, _ := exec.Command("bash", "-c", cmd).Output()
		fmt.Printf("%s\n", out)
		return " >>> docker image was missing, but not anymore.."
	}

	return " >>> docker image already exists :+1:"
}

func wailsInit(distro string) {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	distroFull := "wails-test-" + distro
	cmd := exec.Command("docker", "run", "-it", "--rm", distroFull)
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	go func() {
		c.ExpectEOF()
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
	c.SendLine("wails setup")
	time.Sleep(time.Second)
	c.SendLine("test")
	time.Sleep(time.Second)
	c.SendLine("test@test.test")

	c.SendLine("wails init")
	time.Sleep(time.Second * 2)
	c.SendLine("test")
	time.Sleep(time.Second)
	c.SendLine("test")
	time.Sleep(time.Second)
	c.SendLine("test")
	time.Sleep(time.Second)
	c.SendLine("3")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
