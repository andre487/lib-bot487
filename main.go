package main

import (
	"lib-bot487/util"
	"log/slog"
)

func main() {
	util.SetupLogger()
	slog.Info("Starting slog")
}
