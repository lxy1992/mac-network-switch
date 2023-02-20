package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "toggleAirport",
		Short: "Auto Switch Ethernet And WiFi",
		Run:   ToggleAirport,
	}
	rootCmd.AddCommand(cmd)
}

//nolint:forbidigo // have to
func ToggleAirport(cmd *cobra.Command, _ []string) {
	switchCommand, err := cmd.Flags().GetString("switch")
	if err != nil {
		fail(err.Error())
	}
	// 放置脚本并修改权限
	cmdLine := "sudo cp config/toggleAirport.sh /Library/Scripts/toggleAirport.sh && " +
		"chmod 755 /Library/Scripts/toggleAirport.sh"

	err = runCommand(cmdLine)
	if err != nil {
		fail(err.Error())
	}

	// 放置配置文件
	cmdLine = "sudo cp config/com.mine.toggleairport.plist /Library/LaunchAgents/com.mine.toggleairport.plist " +
		"&& sudo chmod 600 /Library/LaunchAgents/com.mine.toggleairport.plist && " +
		"sudo chown root /Library/LaunchAgents/com.mine.toggleairport.plist && " +
		"sudo chgrp wheel /Library/LaunchAgents/com.mine.toggleairport.plist"

	err = runCommand(cmdLine)
	if err != nil {
		fail(err.Error())
	}

	// 启动或关闭服务
	res := "service error"
	switch switchCommand {
	case "start":
		cmdLine = "sudo launchctl bootstrap system /Library/LaunchAgents/com.mine.toggleairport.plist"
		res = "service start successfully"
	case "stop":
		cmdLine = "sudo launchctl bootout system /Library/LaunchAgents/com.mine.toggleairport.plist"
		res = "service stop successfully"
	default:
		fail("Switch enums in ['start', 'stop']")
	}
	err = runCommand(cmdLine)
	if err != nil {
		fail(err.Error())
	}
	fmt.Println(res)
}

//nolint:forbidigo // have to
func runCommand(cmd string) error {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func fail(msg string) {
	if msg != "" {
		_, err := fmt.Fprintln(os.Stderr, msg)
		if err != nil {
			return
		}
	}
	os.Exit(1)
}
