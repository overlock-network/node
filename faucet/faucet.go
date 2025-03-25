package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cosmossdk.io/math"
	"github.com/blang/semver/v4"
	"github.com/ignite/cli/v28/ignite/pkg/chaincmd"
	chaincmdrunner "github.com/ignite/cli/v28/ignite/pkg/chaincmd/runner"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosfaucet"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosver"
)

func main() {
	ctx := context.Background()

	// Load environment variables
	nodeCmd := os.Getenv("NODE_CMD")
	chainID := os.Getenv("CHAIN_ID")
	keyringDir := os.Getenv("KEYRING_DIR")
	account := os.Getenv("ACCOUNT")

	if nodeCmd == "" || chainID == "" || keyringDir == "" || account == "" {
		log.Fatal("Missing required environment variables: NODE_CMD, CHAIN_ID, KEYRING_DIR, ACCOUNT")
	}

	cmdchain := chaincmd.New(
		nodeCmd,
		chaincmd.WithChainID(chainID),
		chaincmd.WithKeyringBackend("test"),
		chaincmd.WithHome(keyringDir),
	)

	ccr, err := chaincmdrunner.New(ctx, cmdchain)
	if err != nil {
		log.Fatalf("Failed to create chain command runner: %v", err)
	}

	var opt []cosmosfaucet.Option
	opt = append(opt, cosmosfaucet.Account(account, "", ""))

	amount, ok := math.NewIntFromString("1000000")
	if !ok {
		log.Fatalf("Failed to parse amount: %s", "1000000")
	}

	limit, ok := math.NewIntFromString("1000000000")
	if !ok {
		log.Fatalf("Failed to parse limit: %s", "1000000000")
	}

	coinOption := cosmosfaucet.Coin(amount, limit, "utoken")
	versionOption := cosmosfaucet.Version(cosmosver.Version{
		Version:  "0.50.1",
		Semantic: semver.MustParse("0.50.1"),
	})
	opt = append(opt, coinOption, versionOption)

	f, err := cosmosfaucet.New(ctx, ccr, opt...)
	if err != nil {
		log.Fatalf("Failed to initialize faucet: %v", err)
	}

	// Set up HTTP server and route
	http.Handle("/", f)
	port := ":4500"
	log.Printf("Starting faucet server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
