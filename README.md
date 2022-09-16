# Christmas Countdown RPC

<!-- [![Build](https://github.com/christmas-countdown/rpc/actions/workflows/compile.yml/badge.svg)](https://github.com/christmas-countdown/rpc/actions/workflows/compile.yml) -->
[![Build](https://img.shields.io/github/workflow/status/christmas-countdown/rpc/Build?style=flat-square)](https://github.com/christmas-countdown/rpc/actions/workflows/compile.yml)

A super simple program to add a Christmas Countdown to your Discord status.

## Installation

All you need is a single (very small) executable file which can be run directly in your Downloads folder.

<details>
	<summary>📉 Resource usage</summary>

Here it is using just 0.7MB (0.0007GB):
![](https://static.eartharoid.me/k/22/06/30153654.png)
</details>

### Windows

1. Download [`cc-rpc-win-amd64.exe`](https://github.com/christmas-countdown/rpc/releases/download/continuous/cc-rpc-win-amd64.exe).
   > ⚠️ Your browser will probably give you a warning when you try to download it and Windows will try to stop you from running it. In both cases you can click "Keep" or "Run anyway". It's not a virus, as you can see in [main.go](https://github.com/christmas-countdown/rpc/blob/main/main.go).
2. Optionally, [make it run at startup](https://support.microsoft.com/en-us/windows/add-an-app-to-run-automatically-at-startup-in-windows-10-150da165-dcd9-7230-517b-cf3c295d89dd). It has an extremely minimal performance impact.

#### How to make it run in the background (recommended)

Normally, running the `.exe` file will open a terminal window.
To make it hidden (run in the background) create a file, `ChristmasCountdownRPC.vbs`, **in the same folder** that the `.exe` is in, and paste the following:
```vb
Dim WShell
Set WShell = CreateObject("WScript.Shell")
WShell.Run "cc-rpc-win-amd64.exe", 0
Set WShell = Nothing
```

> **Note**
> 
> If you have renamed the `.exe` file, replace `cc-rpc-win-amd64` with the new name in your `.vbs` file.
> 
> **Example:**
> 
> ```diff
> - WShell.Run "cc-rpc-win-amd64.exe", 0
> + WShell.Run "christmas.exe", 0
> ```

⚠️ If you want the program to run automatically at startup, create a shortcut in the startup folder **for the `.vbs` file** instead of the `.exe` file, as this will start it in the background. 

### Other

1. Download the build for your system from the [releases](https://github.com/christmas-countdown/rpc/releases) page.
2. Optionally, make it run at startup ([example guide for Ubuntu](https://www.howtogeek.com/686952/how-to-manage-startup-programs-on-ubuntu-linux))

## Screenshots

![Small profile](https://static.eartharoid.me/k/22/06/30134055.png)

![Large profile, cropped](https://static.eartharoid.me/k/22/06/30011959.png)

## Troubleshooting

### The buttons don't work

The buttons work for other people, but Discord doesn't let you click your own buttons.

### I can't run it

If double-clicking the file on Windows does not work, try opening Command Prompt and then type the file location (e.g. `Downloads/cc-rpc-win-amd64.exe`) followed by the enter key.

On Linux, you may need to right-click on the file, go to Properties, Permissions, and then allow execution as a program.

### It crashes

Make sure you have the Discord desktop app installed and it is running. Discord must be started first.

## Support

<https://lnk.earth/discord>