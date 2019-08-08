package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	//"strings"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	cli()
}

func cli() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var input = "help"
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
			fmt.Printf("Enter: 'distro git branch'")
			scanner.Scan()
			distro := scanner.Text()
			scanner.Scan()
			git := scanner.Text()
			scanner.Scan()
			branch := scanner.Text()

			fmt.Printf("%s %s %s\n", distro, git, branch)
			// do stuff
			fmt.Println("go install && wails init (1)")
			fmt.Println("go install && bin/bash /root (2)")
			fmt.Println("go install && wails init && bin/bash /root (3)")
			scanner.Scan()
			input := scanner.Text()
			switch input {
			case "1":
				goInstallWailsInit()
			case "2":
				//goInstallBashRoot()
			case "3":
				//goInstallBoth()
			}

			// when finished return for user input
			fmt.Printf("finished doing stuff!")
			fmt.Printf("enter new command:")
			scanner.Scan()
			text := scanner.Text()
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
			fmt.Println("test-branch $distro $git $branch (1) distribution to test on, git repo and specific branch to test against")
			fmt.Println("supported-distros (7) show all currently support distributions")
			fmt.Println("tester-prune (8) delete from host all wails built docker images")
			fmt.Println("exit (0) exit the tester")
			fmt.Println("help (9) (meta)")

			scanner.Scan()
			text := scanner.Text()
			//cmd := strings.TrimSuffix(text, "\n")
			input = text

		}
	}
}

func goInstallWailsInit() {
	// check the image of selected distro exists
	// if not use Dockerfile to build it

	// run 'git-branch' container to produce the test build

	// move newly built wails binary to 'test-branch' dir

	// start the container

}

func wailsInit() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("wails", "init")
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

func checkDockerImageExist() {

}
