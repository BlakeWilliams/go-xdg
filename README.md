# go-XDG

This package implements (a subset of) the XDG Base Directory Specification in
Go. It uses the user expected directories as defined by the XDG Base Directory
Specification for Linux/Unix/macOS and (soon) Windows.

For Linux/Unix/macOS the following directories are used:

- Config - `$XDG_CONFIG_HOME` or `$HOME/.config`
- Data - `$XDG_DATA_HOME` or `$HOME/.local/share`
- Cache - `$XDG_CACHE_HOME` or `$HOME/.cache`
- State - `$XDG_STATE_HOME` or `$HOME/.local/state`

## Installation

```sh
go get github.com/blakewilliams/go-xdg
```

## Usage

The API is relatively simple. Each directory has a function that returns the
path to the directory or an error (e.g. if $HOME is not set in macOS/Linux).

```go
// Get the config directory
configDir, err := ConfigHome()

// Get the data directory
dataDir, err := DataHome()

// Get the cache directory
cacheDir, err := CacheHome()

// Get the state directory
stateDir, err := StateHome()
```

This package also has support for finding files. It will return the path to the
file if it is found, otherwise it returns an `ErrNoFile` error.

```go
// Find `myapp/config.yaml` in the config directory
configFilePath, err := FindConfigFile("myapp", "config.yaml")

// Find `myapp/data.json` in the data directory
dataFilePath, err := FindDataFile("myapp", "data.json")

// Find `myapp/cache.db` in the cache directory
dataFilePath, err := FindCacheFile("myapp", "cache.db")

// Find `myapp/state.db` in the state directory
stateFilePath, err := FindStateFile("myapp", "state.db")
```

## Contributions welcome!

If you'd like to contribute, please create an issue with a proposal for your
change before opening a PR if you'd like to ensure it will be accepted.
