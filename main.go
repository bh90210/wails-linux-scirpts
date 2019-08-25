package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	// I used to hate netflix, now I kind of love them!
	expect "github.com/Netflix/go-expect"
	"github.com/abiosoft/ishell"
	_ "github.com/fatih/color"
)

func main() {
	shell()
}

func shell() {
	// start a new shell
	shell := ishell.New()

	// display info.
	shell.Println("τέστερ v0.2-alpha")
	shell.Println("available commands:")
	shell.Println(" * branch (1) | distribution, git repo and specific branch to test against")
	shell.Println(" * all (2) | test (go install & wails init) of given git & branch against all supported distros (cpu intense!) NOTWORKINGYET")
	//LABEL wails="remove"
	shell.Println(" * prune (3) | delete wails built docker images from host (excluding distro builds)")
	// LABEL wails="removeall"
	shell.Println(" * prune-all (4) | delete wails built docker images from host (including distro builds) -- NOTWORKINGYET")
	shell.Println("exit | exit τέστερ")
	shell.Println("help | (meta)")

	shell.AddCmd(&ishell.Cmd{
		Name:    "branch",
		Aliases: []string{"1"},
		Help:    "a) select distro, b) enter GIT and BRANCH to check against",
		Func: func(c *ishell.Context) {
			// cd in supported-distros directory and print in MultiChoice
			cmd := "cd ./supported-distros && ls"
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				fmt.Sprintf("failed to execute command: %s", cmd)
			}
			//fmt.Printf("%s\n", out)
			stringConvert := string(out)
			distros := strings.Fields(stringConvert)
			cliMultiChoiceList := []string{}

			for _, distro := range distros {
				cliMultiChoiceList = append(cliMultiChoiceList, distro)
			}

			choice := c.MultiChoice(cliMultiChoiceList, "Choose a distro")
			var distro string
			for {
				// for each new distro we support, and a dockerfile exist
				// we should make an entry here as well
				switch choice {
				case 0:
					distro = "alpine310"
				case 1:
					distro = "archlinux"
				case 2:
					distro = "centos7"
				case 3:
					distro = "debian9"
				case 4:
					distro = "elementary5"
				case 5:
					distro = "fedora30"
				case 6:
					distro = "mint19"
				case 7:
					distro = "opensusetumbleweed"
				case 8:
					distro = "parrot47"
				case 9:
					distro = "ubuntu1804"
				case 10:
					distro = "voidlinux"
				case 11:
					distro = "voidlinux-musl"
				}
				break
			}

			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Println("Enter Git and Branch")
			// prompt for input
			c.Print("Git repo: ")
			git := c.ReadLine()
			c.Print("Branch: ")
			branch := c.ReadLine()

			choice = c.MultiChoice([]string{
				"go install && wails init",
				"go install && bin/bash /root -- NOTWORKINGYET",
				"go install && wails setup && bin/bash /root -- NOTWORKINGYET",
			}, "Choose a build/test option")
			var option string
			for {
				switch choice {
				case 0:
					option = "wailsinit"
				case 1:
					option = "binbash"
				case 2:
					option = "wailsinitbinbash"
				}
				break
			}

			goInstallWails(distro, git, branch, option)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "prune",
		Aliases: []string{"3"},
		Help:    "delete wails built docker images from host (not the distro builds)",
		Func: func(c *ishell.Context) {
			c.Println("deleting images..")
			dockerImagesPrune()
		},
	})

	// when started with "exit" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}

	// loop:
	// 	for {
	// 		switch input {
	// 		// when user enters a value not declared in any 'case'
	// 		// app hangs. 'default' catches this and gracefully
	// 		// returns some help and puts user back in the 'for' loop
	// 		default:
	// 			fmt.Println("command does not exist!")
	// 			fmt.Println("if you need it type 'help'")
	// 			scanner.Scan()
	// 			text := scanner.Text()
	// 			input = text
	//
	// 		case "test-branch", "1":
	// 			// show a list of all supported-distros
	// 			// because probably you don't remember it
	// 			cmd := "cd ./supported-distros && ls"
	// 			out, err := exec.Command("bash", "-c", cmd).Output()
	// 			if err != nil {
	// 				fmt.Sprintf("failed to execute command: %s", cmd)
	// 			}
	// 			fmt.Printf("%s", out)
	//
	// 			fmt.Printf(" - enter: 'distro git branch'\n")
	// 			// TODO: code cleanup
	// 			scanner.Scan()
	// 			distro := scanner.Text()
	// 			scanner.Scan()
	// 			git := scanner.Text()
	// 			scanner.Scan()
	// 			branch := scanner.Text()
	// 			//fmt.Printf("%s %s %s\n", distro, git, branch)
	//
	// 			fmt.Println("go install && wails init (1)")
	// 			// TODO: fix the NOTWORKINGYET
	// 			fmt.Println("go install && bin/bash /root (2) NOTWORKINGYET")
	// 			fmt.Println("go install && wails init && bin/bash /root (3) NOTWORKINGYET")
	// 			scanner.Scan()
	// 			text := scanner.Text()
	// 			switch text {
	// 			case "1":
	// 				goInstallWailsInit(distro, git, branch)
	// 			case "2":
	// 				//goInstallBashRoot()
	// 			case "3":
	// 				//goInstallBoth()
	// 			}
	//
	// 			// when finished return for user input
	// 			fmt.Println("finished doing stuff!")
	// 			fmt.Println("enter new command:")
	// 			scanner.Scan()
	// 			text = scanner.Text()
	// 			input = text
	//
	// 		case "supported-distros", "7":
	// 			cmd := "cd ./supported-distros && ls"
	// 			out, err := exec.Command("bash", "-c", cmd).Output()
	// 			if err != nil {
	// 				fmt.Sprintf("failed to execute command: %s", cmd)
	// 			}
	// 			fmt.Printf("%s\n", out)
	//
	// 			scanner.Scan()
	// 			text := scanner.Text()
	// 			input = text
	//
	// 			// exit
	// 		case "exit", "0":
	// 			fmt.Println("testing is over")
	// 			fmt.Println("hopefully everything works :+1: :+1:")
	// 			fmt.Println("see u soon!")
	// 			break loop
	//
	// 		case "help", "9":
	// 			fmt.Println("wails-linux-scripts v0.1-alpha helpfile!")
	// 			fmt.Println("available commands:")
	// 			fmt.Println(" * test-branch $distro $git $branch (1) distribution, git repo and specific branch to test against")
	// 			fmt.Println(" * test-all (2) test (go install & wails init) of given git & branch against all supported distros (cpu intense!) NOTWORKINGYET")
	// 			fmt.Println(" * supported-distros (7) show all currently support distributions")
	// 			fmt.Println(" * tester-prune (8) delete from host all wails built docker images")
	// 			fmt.Println(" * exit (0) exit the tester")
	// 			fmt.Println(" * help (9) (meta)")
	//
	// 			scanner.Scan()
	// 			text := scanner.Text()
	// 			//cmd := strings.TrimSuffix(text, "\n")
	// 			input = text
	//
	// 		}
	// 	}
	//
}

func goInstallWails(distro, git, branch string, option string) {
	// check the image of selected distro exists
	// if not use Dockerfile to build it
	feedback := checkDockerImageExist(distro)
	log.Println(feedback)

	// build 'git-branch' container to produce the test build

	// TODO: feat add --no-cache flag as option
	cmd := "docker build -t wails-test-latest --build-arg GIT=" + git + " --build-arg BRANCH=" + branch + " --no-cache ./git-branch"
	//cmd := "docker build -t wails-test-latest --build-arg GIT=" + git + " --build-arg BRANCH=" + branch + " ./git-branch"
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

	// TODO: remove wails-test-latest docker images

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
	string := string(out[:])
	fmt.Printf("checking if %s is on host\n", distro)
	if string == "0\n" {
		fmt.Printf("image for %s doesn't exist, building now, this will take a while\n", distro)
		cmd := "cd supported-distros/" + distro + " && docker build -t wails-" + distro + " ."
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

	time.Sleep(time.Second * 4)
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

func dockerImagesPrune() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	label := "label=wails=remove"
	cmd := exec.Command("docker", "image", "prune", "-a", "--filter", label)
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
	time.Sleep(time.Second)
	c.SendLine("y")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
