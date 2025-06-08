package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"token-monitor/config"
	"token-monitor/contracts/restrict"
	"token-monitor/services"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Ethereum client
	rpcClient, err := rpc.Dial(cfg.Monitor.EthereumWSURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RPC: %v", err)
	}
	client := ethclient.NewClient(rpcClient)

	// Create owner key for contract interactions
	privateKey, err := crypto.HexToECDSA(os.Getenv("BLACKLIST_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Create analyzer service
	analyzer := services.NewAnalyzer(
		db,
		cfg.Monitor.LargeAmountThreshold,
		10, // shortTimeBlocks
		cfg.Monitor.SuspiciousAddresses,
		time.Second*5, // analysis interval
	)

	// Create monitor service
	monitor, err := services.NewMonitor(db, cfg.Monitor, analyzer)
	if err != nil {
		log.Fatalf("Failed to create monitor: %v", err)
	}

	// Create blacklist monitor
	restrictContract, err := restrict.NewRestrict(common.HexToAddress(os.Getenv("RESTRICT_CONTRACT_ADDRESS")), client)
	if err != nil {
		log.Fatalf("Failed to create Restrict contract instance: %v", err)
	}

	blacklistMonitor := services.NewBlacklistMonitor(
		db,
		restrictContract,
		common.HexToAddress(os.Getenv("RESTRICT_CONTRACT_ADDRESS")),
		auth,
		time.Second*5, // Check for new suspicious addresses every 10s
	)

	// Create mempool monitor
	mempoolMonitor := services.NewMempoolMonitor(
		client,
		rpcClient,
		db,
		common.HexToAddress(cfg.Monitor.ContractAddress),
		analyzer,
		time.Second,
	)

	// Create web server

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start analyzer
	analyzer.Start(ctx)

	// Start monitoring
	if err := monitor.Start(ctx); err != nil {
		log.Fatalf("Failed to start monitoring: %v", err)
	}

	// Start blacklist monitor
	blacklistMonitor.Start(ctx)

	// Start mempool monitor
	mempoolMonitor.Start(ctx)

	// Start web server in a goroutine

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	monitor.Stop()
	analyzer.Stop()
	blacklistMonitor.Stop()
	mempoolMonitor.Stop()
}
