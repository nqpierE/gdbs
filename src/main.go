package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/slothryo/gdbs/install"

	"github.com/slothryo/gdbs/setup"
)

func include(list []string, target string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			return true
		}
	}
	return false
}

type Gdbmod struct {
	Name            string
	InstallDir      string
	RootDir         string
	GdbmodsDirsList []string
}

func Install(mod Gdbmod) bool {
	if mod.Name == "gdb" {
		return true
	}
	if !include(mod.GdbmodsDirsList, mod.Name) {
		fmt.Println("koko")
		var err_flag bool = false
		switch mod.Name {
		case "peda":
			if !install.Installpeda(mod.InstallDir) {
				err_flag = !err_flag
			}
		case "gef":
			if !install.Installgef(mod.InstallDir) {
				err_flag = !err_flag
			}
		case "pwndbg":
			if !install.Installgef(mod.InstallDir) {
				err_flag = !err_flag
			}
		}
		if err_flag {
			fmt.Fprintln(os.Stderr, "failed to install "+mod.Name)
			return false
		} else {
			fmt.Println("installed " + mod.Name + " successfully")
			return true
		}
	} else {
		return true
	}
}

func Setup(mod Gdbmod) {
	var err_flag bool = false
	switch mod.Name {
	case "gdb":
		if !setup.Setgdb() {
			err_flag = !err_flag
		}
	case "peda":
		fmt.Println(err_flag)
		if !setup.Setpeda(mod.InstallDir) {
			err_flag = !err_flag
		}
	case "gef":
		if !setup.Setgef(mod.InstallDir) {
			err_flag = !err_flag
		}
	case "pwndbg":
		if !setup.Setpwndbg(mod.InstallDir) {
			err_flag = !err_flag
		}
	}
	if err_flag {
		fmt.Fprintln(os.Stderr, "[x] failed to change tool to "+mod.Name)
		return
	} else {
		fmt.Println("[*] changed tool to " + mod.Name + " successfully")
		err := os.Truncate(mod.RootDir+"/state.txt", 0)
		if err != nil {
			os.Create(mod.RootDir + "/state.txt")
		}
		os.WriteFile(mod.RootDir+"/state.txt", []byte(mod.Name), 0644)
		return
	}
}

func main() {
	var argslen int = len(os.Args) - 1
	if argslen == 0 {
		fmt.Println("Usage:")
		fmt.Println("  command [option]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  show			show current tool")
		fmt.Println("  gdb			set tool to legacy gdb")
		fmt.Println("  peda			set tool to gdb-peda")
		fmt.Println("  gef			set tool to gef")
		fmt.Println("  pwndbg		set tool to pwndbg")
		fmt.Println("  clean			clean application")
		return
	}
	if argslen == 1 {

		root_dir, _ := os.Getwd()
		root_dirs, _ := os.ReadDir(root_dir)
		var root_dirs_list []string = []string{}
		for i := 0; i < len(root_dirs); i++ {
			root_dirs_list = append(root_dirs_list, root_dirs[i].Name())
		}
		if include(root_dirs_list, "gdbmods") {
			os.Mkdir(root_dir+"/gdbmods/", 0755)
		}
		fmt.Println(root_dir)
		var gdbmods_dir string = root_dir + "/gdbmods/"
		gdbmods_dirs, _ := os.ReadDir(root_dir + "/gdbmods/")
		var gdbmods_dirs_list []string = []string{}
		for i := 0; i < len(gdbmods_dirs); i++ {
			gdbmods_dirs_list = append(gdbmods_dirs_list, gdbmods_dirs[i].Name())
		}
		var mod_list []string = []string{"gdb", "peda", "gef", "pwndbg"}
		var state_file_path string = root_dir + "/state.txt"
		for i := 0; i < len(mod_list); i++ {
			if os.Args[1] == mod_list[i] {
				var eachmod Gdbmod
				eachmod.Name = mod_list[i]
				eachmod.InstallDir = gdbmods_dir + eachmod.Name
				fmt.Println(eachmod.InstallDir)
				eachmod.RootDir = root_dir
				eachmod.GdbmodsDirsList = gdbmods_dirs_list
				if !Install(eachmod) {
					return
				} else {
					Setup(eachmod)
				}
			}
		}
		switch os.Args[1] {
		//legacy gdb session
		case "show":
			state, err := os.ReadFile(state_file_path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "[x] the file to manage state of tools is not found")
				return
			} else {
				if include(mod_list[:], string(state)) {
					for i := 0; i < len(mod_list); i++ {
						if string(state) == mod_list[i] {
							fmt.Println("[" + color.RedString("*") + "] " + color.RedString(mod_list[i]))
						} else {
							fmt.Println("[-] " + mod_list[i])
						}
					}
					return
				} else {
					os.Truncate(state_file_path, 0)
					os.WriteFile(state_file_path, []byte("gdb"), 0644)
					return
				}
			}
		//cleaning session
		case "clean":
			err := os.RemoveAll(gdbmods_dir)
			if err != nil {
				fmt.Fprintln(os.Stderr, "[x] ")
				fmt.Fprintln(os.Stderr, err)
				fmt.Fprintln(os.Stderr, "[x] failed to clean up")
				return
			} else {
				if !(setup.Setgdb()) {
					fmt.Fprintln(os.Stderr, "[x] failed to change tool to legacy gdb")
					return
				} else {
					fmt.Println("[*] changed tool to legacy gdb successfully")
					err := os.Truncate(state_file_path, 0)
					if err != nil {
						os.Create(state_file_path + "/state.txt")
					} else {
						os.WriteFile(state_file_path+"/state.txt", []byte("gdb"), 0644)
					}
					fmt.Println("[*] uninstalled all tools successfully")
					return
				}
			}
		}
		//help session
	} else {
		return
	}
}
