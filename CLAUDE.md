# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Omni is an all-in-one media manager written in Go (module `github.com/gregmulvaney/omni`). The project is at a very early stage: only the CLI skeleton and config scaffolding exist.

## Commands

- `make run` — run the app (`go run ./cmd/omni`)
- `make dev` — run with `--log.dev --log.format console`
- `make build` — build to `./bin/omni` (gitignored)
- `go run ./cmd/omni config generate [path] [-f]` — write a default config file (default path `./config/omni.yaml`; `-f` overwrites)
- `go test ./...` / `go test ./path/to/pkg -run TestName` — run all tests / a single test

There are no tests or lint config yet.

## Architecture

- `cmd/omni/main.go` is a thin entrypoint: it creates a context cancelled on SIGINT/SIGTERM and calls `cli.Execute(ctx)`, exiting 1 on error. Keep logic out of `main`.
- `internal/cli` holds the cobra command tree. `rootCmd()` builds an `app` struct (a per-invocation `*viper.Viper` plus the shared config `*pflag.FlagSet`) and passes it to subcommand constructors such as `configCmd(a)`. New subcommands should follow this constructor pattern and be registered in `rootCmd()`.
- `internal/config` owns configuration: the `Config` struct (`mapstructure` tags; `server`, `db`, `log` sections), `RegisterFlags()` (the persistent flags `--log.level`, `--log.format`, `--log.dev`), and `WriteDefault()`. Flag names are dotted so they map onto nested config keys.

## Work in progress

- `config.Load()` is an empty stub, so the config file and flags are not yet read into `Config`. `--log.*` flags are accepted but have no effect.
- The root command has no `Run`, so `make run` only prints help.
- `config generate` only writes the `log` section; `server` and `db` have no flag defaults yet.
