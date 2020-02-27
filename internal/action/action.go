package action

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

// Initialize is used to create a new machine
func Initialize(c *cli.Context) error {
	machine := c.String("machine")

	multipass := fmt.Sprintf("%s", c.Context.Value("multipass"))

	// create the machine
	// TODO move the yaml file into the binary as stdin
	cmd := exec.Command(multipass, "launch", "--name", machine, "--cloud-init", "./internal/init.yaml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Bootstrap(c *cli.Context, e CommandLineExecutor) error {
	machine := c.String("machine")
	php := c.String("php-version")
	database := c.String("database")

	args := []string{"multipass", "exec", machine, "--", "sudo", "bash", "/opt/nitro/bootstrap.sh", php, database}

	return e.Exec(e.Path(), args, os.Environ())
}

// Update will perform system updates on a given machine
func Update(c *cli.Context) error {
	machine := c.String("machine")
	multipass := fmt.Sprintf("%s", c.Context.Value("multipass"))

	cmd := exec.Command(multipass, "exec", machine, "--", "sudo", "bash", "/opt/nitro/update.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func AddHost(c *cli.Context, e CommandLineExecutor) error {
	machine := c.String("machine")
	host := c.Args().First()
	php := c.String("php-version")

	if host == "" {
		return errors.New("missing param host")
	}

	if php == "" {
		fmt.Println("missing php-version")
		php = "7.4"
	}

	args := []string{"multipass", "exec", machine, "--", "sudo", "bash", "/opt/nitro/nginx/add-site.sh", host, php}

	return e.Exec(e.Path(), args, os.Environ())
}

// SSH will login a user to a specific machine
func SSH(m string, e CommandLineExecutor) error {
	return e.Exec(e.Path(), []string{"multipass", "shell", m}, os.Environ())
}

func Delete(c *cli.Context) error {
	machine := c.String("machine")

	multipass := fmt.Sprintf("%s", c.Context.Value("multipass"))

	cmd := exec.Command(multipass, "delete", machine)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Stop(c *cli.Context) error {
	machine := c.String("machine")

	multipass := fmt.Sprintf("%s", c.Context.Value("multipass"))

	cmd := exec.Command(multipass, "stop", machine)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func DatabasePassword(c *cli.Context, e CommandLineExecutor) error {
	return e.Exec(e.Path(), []string{"multipass", "exec", c.String("machine"), "--", "cat", "/home/ubuntu/.db_password"}, os.Environ())
}
