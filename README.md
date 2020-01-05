[![Build Status](https://travis-ci.com/ioncloud64/freemegb.svg?branch=lang%2Fgo)](https://travis-ci.com/ioncloud64/freemegb)
![GitHub issues](https://img.shields.io/github/issues-raw/ioncloud64/freemegb)
![License](https://img.shields.io/github/license/ioncloud64/freemegb)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/ioncloud64/freemegb)
![GitHub last commit](https://img.shields.io/github/last-commit/ioncloud64/freemegb)
# FreeMe!GB
A GameBoy simulator written in Go. GUI bindings are written for GTK+. This program is designed to simulate the real hardware that was equipped in the system itself.

## Features
*Note: not all features may be implemented yet*
* Debugging Utility
  - CPU Registers
  - OPCODE descriptions
  - Breakpoint debugging
* Emulation Core
  - ROM Detection
    + ROM Name
    + ROM Type
    + ROM Size
    + Compatibility Check
* Controller Support (Not yet planned)
* Shader effects (Not yet planned)

## Build Requirements
* Windows
  - Install MSYS2 mingw64
  - Modify bash profile to include GoPath, Project Directory, and mingw64 bin folder
  - [Install dependencies](https://github.com/gotk3/gotk3/wiki)
* Linux
  - [Install dependencies](https://github.com/gotk3/gotk3/wiki)
* Mac OS X
  - Not planned

# Project Inspiration
I have always interested in how emulators work, so I started to have this idea of creating my own emulator variant of the original GameBoy. This simulator is designed to be both educational and eventually complete. Please give me time as this is a personal project that may be dormant from time to time, depending on how life is at home. Feel free to give suggestions and tips at admin@ioncloud64.com.

I look forward to completing FreeMe!GB and discovering Node.js's ability to perform.

Please check out the [wiki](https://github.com/ioncloud64/freemegb/wiki) for more information.
