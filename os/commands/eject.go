package commands

import "os/exec"

func EjectDevice(device string) error {
	// execute the bash command to eject the device
	cmd := exec.Command("eject", device)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func InsertDevice(device string) error {
	// execute the bash command to insert the device
	cmd := exec.Command("eject", "-t", device)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
