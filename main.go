package main

import (
	"bytes"
	"encoding/hex"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/lmittmann/tint"
	buyer "github.com/selesy/x402-buyer"
)

func main() {
	const (
		dinNetURL = "https://blast-mainnet.gateway.defura.dev.infhq.net/rpc"
		rpcReq    = `{"jsonrpc": "2.0", "method": "eth_getBlockByNumber", "params": ["latest", false], "id": 1}`
	)

	log := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		// Level: slog.LevelDebug,
	}))

	privHex, ok := os.LookupEnv("X402_BUYER_PRIVATE_KEY")
	if !ok {
		log.Error("X402_BUYER_PRIVATE_KEY not set")
		os.Exit(1)
	}

	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		log.Error("Failed to decode hex private key value", tint.Err(err))
		os.Exit(1)
	}

	priv := secp256k1.PrivKeyFromBytes(privBytes)

	client, err := buyer.ClientForKey(priv.ToECDSA(), buyer.WithLogger(log))
	if err != nil {
		log.Error("Failed to create buyer client", tint.Err(err))
		os.Exit(1)
	}

	client.Timeout = 30 * time.Second

	log.Info("DIN Request:", slog.String("body", rpcReq))

	resp, err := client.Post(dinNetURL, "application/json", bytes.NewReader([]byte(rpcReq)))
	if err != nil {
		log.Error("Failed to send DIN request", tint.Err(err))
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read JSON-RPC response", tint.Err(err))
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		log.Error("Unexpected HTTP status code", slog.Int("code", resp.StatusCode))
		os.Exit(1)
	}

	log.Info("DIN response", slog.String("body", string(body)), slog.Int("code", resp.StatusCode))

	for k, vs := range resp.Header {
		for _, v := range vs {
			log.Debug("HTTP response header", slog.String("key", k), slog.String("value", v))
		}
	}
}
