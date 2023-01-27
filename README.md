# gdbs
Application to switch tools of gdb. This application supports legacy gdb, peda, gef, pwndbg.

# installaion
Run following commands to install application.
```
git clone https://github.com/slothryo/gdbs.git
cd gdbs
chmod +x install.sh
./install.sh
```
This application also manages installation of all tools so, you needn't install whatever install tools.
However, tools are installed by this application won't conflict tools are already installed in your system.

# Usage
Run following command to show usage.
```
gdbs
```
```
Usage:
  command [option]

Options:
  show			show current tool
  gdb			set tool to legacy gdb
  peda			set tool to gdb-peda
  gef			set tool to gef
  pwndbg		set tool to pwndbg
  clean			clean application
```

# Uninstallation
```
gdbs clean
```
This tool manages all in directory you git cloned and ~/.gdbs. Above command removes ~/.gdbs so, when you want to uninstall this tool, run above command and remove directory you git cloned.



