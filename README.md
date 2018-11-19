# got
got is simple tmux tool

## Support OS
- MacOS
- Linux

## Features
- new session
- attach session
- kill session
- swtich session

## Required
tmux

## Instllation/Update
```
$ go get -u github.com/skanehira/got
$ got
```

## Release
- [MacOS](https://github.com/skanehira/got/releases/download/v1.0.3/MacOS.zip)
- [Linux](https://github.com/skanehira/got/releases/download/v1.0.3/Linux.zip)

## Homebrew (or Linuxbrew)
```
$ brew tap skanehira/got
$ brew install got
```

When installing on x86-64 Linux, do as follows.

```
$ brew tap z80oolong/tmux
$ brew install z80oolong/tmux/got-bin
```

Also, when compiling and installing "got" from the source code, do as follows.

```
$ brew tap z80oolong/tmux
$ brew install z80oolong/tmux/got-src
$ brew link --force z80oolong/tmux/got-src
```

## Usage
| Operation | Key                            |
|-----------|--------------------------------|
| quit      | <kbd>Tab</kbd>  + <kbd>d</kbd> |
| cancel    | <kbd>Ctrl</kbd> + <kbd>c</kbd> |
| select    | <kbd>Enter</kbd>               |

## Screenshots
![s1.png](https://github.com/skanehira/got/blob/master/images/s1.png)
