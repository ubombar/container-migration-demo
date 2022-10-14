package app

import (
	"bufio"
	"os"

	"github.com/checkpoint-restore/go-criu/v6"
	"github.com/checkpoint-restore/go-criu/v6/crit/images"
)

// App node client
type AppNodeClient struct {
	// Criu client, used for communication criu deamon
	CriuClient *criu.Criu

	// For getting input from user
	Reader *bufio.Reader
}

// Creates a new instance of the client node
func NewClient() *AppNodeClient {
	app := &AppNodeClient{
		CriuClient: criu.MakeCriu(),
		Reader:     bufio.NewReader(os.Stdin),
	}

	return app
}

// Gets input from the user
func (a *AppNodeClient) GetInput() string {
	text, err := a.Reader.ReadString('\n')

	if err != nil {
		panic("An error occured while reading user input")
	}

	return text
}

// Gets the criu version, if criu is not installed returns a non nil error
func (a *AppNodeClient) GetCriuVersion() (int, error) {
	return a.CriuClient.GetCriuVersion()
}

// Migrates the process to another process. To do this first a checkpoint
// is created. Then if the server instance is running the
func (a *AppNodeClient) MigrateContainer(pid int32) error {
	imagesFile, err := os.OpenFile("./dump", os.O_RDWR, 0755)

	if err != nil {
		return err
	}

	fd := int32(imagesFile.Fd())

	err = a.CriuClient.Dump(&images.CriuOpts{
		Pid:         &pid,
		ShellJob:    &[]bool{true}[0], // Idk why they want a *bool?
		ImagesDirFd: &fd,
	}, criu.NoNotify{})

	return err
}

// Restore the process from a file
func (a *AppNodeClient) RestoreContainer() error {
	imagesFile, err := os.OpenFile("./dump", os.O_RDONLY, 0755)

	if err != nil {
		return err
	}

	fd := int32(imagesFile.Fd())

	err = a.CriuClient.Restore(&images.CriuOpts{
		ShellJob:    &[]bool{true}[0],
		ImagesDirFd: &fd,
	}, criu.NoNotify{})

	return err
}
