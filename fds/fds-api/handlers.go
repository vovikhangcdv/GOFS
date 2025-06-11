package main

import (
	"net/http"
	"os"

	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fds-api/contracts"

	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"encoding/json"
	"time"
)

var (
	rpcURL           string
	contractAddress  string
	privateKeyHex    string
	evndTokenAddress string
)

func init() {
	godotenv.Load()
	rpcURL = os.Getenv("MAINNET_RPC_URL")
	contractAddress = os.Getenv("RESTRICT_CONTRACT_ADDRESS")
	privateKeyHex = os.Getenv("BLACKLIST_PRIVATE_KEY")
	evndTokenAddress = os.Getenv("EVND_TOKEN_ADDRESS")
}

func getAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, _ := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	chainID, _ := client.NetworkID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	return auth, nil
}

func CallBlacklistContract(addresses []string) (string, error) {
	client, _ := ethclient.Dial(rpcURL)
	defer client.Close()
	auth, _ := getAuth(client)
	contract, _ := contracts.NewContracts(common.HexToAddress(contractAddress), client)

	var addrList []common.Address
	for _, addr := range addresses {
		addrList = append(addrList, common.HexToAddress(addr))
	}

	tx, err := contract.Blacklist(auth, addrList)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func CallUnblacklistContract(addresses []string) (string, error) {
	client, _ := ethclient.Dial(rpcURL)
	defer client.Close()
	auth, _ := getAuth(client)
	contract, _ := contracts.NewContracts(common.HexToAddress(contractAddress), client)

	var addrList []common.Address
	for _, addr := range addresses {
		addrList = append(addrList, common.HexToAddress(addr))
	}

	tx, err := contract.Unblacklist(auth, addrList)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// getAddressTotals returns the total amount in and out for a given address
func getAddressTotals(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	var inTotal, outTotal float64

	// Query total amount received (in)
	if err := db.Unscoped().Model(&Transaction{}).Where("to_address = ?", address).Select("COALESCE(SUM(CAST(value AS FLOAT)), 0)").Scan(&inTotal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Query total amount sent (out)
	if err := db.Unscoped().Model(&Transaction{}).Where("from_address = ?", address).Select("COALESCE(SUM(CAST(value AS FLOAT)), 0)").Scan(&outTotal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"in":      inTotal,
		"out":     outTotal,
	})
}

// getSuspiciousTransactions returns a list of suspicious transactions
func getSuspiciousTransactions(c *gin.Context) {
	var transactions []SuspiciousTransfer
	if err := db.Order("created_at DESC").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// getBlacklist returns a list of blacklisted addresses and their reasons
func getBlacklist(c *gin.Context) {
	var addresses []BlacklistedAddress
	if err := db.Order("created_at DESC").Find(&addresses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

// getRelatedAddresses returns addresses that have sent or received tokens from/to the given address
func getRelatedAddresses(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	var relatedAddresses []string

	// Query addresses that sent tokens to the given address
	var senders []string
	if err := db.Model(&Transaction{}).Where("to_address = ?", address).Distinct().Pluck("from_address", &senders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Query addresses that received tokens from the given address
	var receivers []string
	if err := db.Model(&Transaction{}).Where("from_address = ?", address).Distinct().Pluck("to_address", &receivers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Combine unique addresses
	relatedAddresses = append(relatedAddresses, senders...)
	relatedAddresses = append(relatedAddresses, receivers...)

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"related": relatedAddresses,
	})
}

// getRelatedTransactionsOfSuspicious returns all transactions involving the same addresses as the suspicious tx
func getRelatedTransactionsOfSuspicious(c *gin.Context) {
	txHash := c.Query("txHash")
	if txHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "txHash parameter is required"})
		return
	}

	var suspiciousTx SuspiciousTransfer
	if err := db.Where("tx_hash = ?", txHash).First(&suspiciousTx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Suspicious transaction not found"})
		return
	}

	var relatedTxs []Transaction
	if err := db.Where(
		"from_address = ? OR to_address = ? OR from_address = ? OR to_address = ?",
		suspiciousTx.From, suspiciousTx.From, suspiciousTx.To, suspiciousTx.To,
	).Order("block_number DESC").Find(&relatedTxs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relatedTxs)
}

// getTransactionsByAddress returns all transactions where the address is sender or receiver
func getTransactionsByAddress(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	var txs []Transaction
	if err := db.Unscoped().Where("from_address = ? OR to_address = ?", address, address).
		Order("block_number DESC").
		Find(&txs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txs)
}

type BlacklistRequest struct {
	Addresses []string `json:"addresses"`
}

func blacklistAddresses(c *gin.Context) {
	var req BlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "addresses required"})
		return
	}

	// Call contract
	txHash, err := CallBlacklistContract(req.Addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "blacklisted", "addresses": req.Addresses, "txHash": txHash})
}

func unblacklistAddresses(c *gin.Context) {
	var req BlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Addresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "addresses required"})
		return
	}

	// Call contract
	txHash, err := CallUnblacklistContract(req.Addresses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "unblacklisted", "addresses": req.Addresses, "txHash": txHash})
}

func deleteBlacklistAddress(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address required"})
		return
	}

	if err := db.Where("address = ?", req.Address).Delete(&BlacklistedAddress{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted", "address": req.Address})
}

// getETHBalance returns the ETH balance for a given address
func getETHBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum node"})
		return
	}
	defer client.Close()

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ETH balance"})
		return
	}

	// Convert balance from wei to ETH (1 ETH = 10^18 wei)
	ethBalance := float64(balance.Int64()) / 1e18

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": ethBalance,
		"unit":    "ETH",
	})
}

// getERC20Balance returns the token balance for a given address and token contract
func getERC20Balance(client *ethclient.Client, tokenAddress string, userAddress string) (float64, error) {
	// ERC20 balanceOf function signature
	balanceOfSignature := []byte("balanceOf(address)")
	balanceOfHash := crypto.Keccak256(balanceOfSignature)[:4]

	// Pack the address parameter
	addr := common.HexToAddress(userAddress)
	paddedAddress := common.LeftPadBytes(addr.Bytes(), 32)

	// Combine the function selector and the padded address
	data := append(balanceOfHash, paddedAddress...)

	// Create the call message
	tokenAddr := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		To:   &tokenAddr,
		Data: data,
	}

	// Make the call
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return 0, err
	}

	// Convert the result to a big.Int
	balance := new(big.Int).SetBytes(result)

	// Convert to float64 (assuming 18 decimals like most ERC20 tokens)
	balanceFloat := new(big.Float).SetInt(balance)
	balanceFloat.Quo(balanceFloat, big.NewFloat(1e18))

	balanceFloat64, _ := balanceFloat.Float64()
	return balanceFloat64, nil
}

// getEVNDBalance returns the eVND token balance for a given address
func getEVNDBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	if evndTokenAddress == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "EVND token address not configured"})
		return
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum node"})
		return
	}
	defer client.Close()

	// Get the balance using ERC20 balanceOf
	balance, err := getERC20Balance(client, evndTokenAddress, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get eVND balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": balance,
		"unit":    "eVND",
	})
}

// getRules returns all compliance rules
func getRules(c *gin.Context) {
	var rules []Rule
	if err := db.Order("id").Find(&rules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}

// getRuleViolations returns all rules with their violation counts in the last 24h
func getRuleViolations(c *gin.Context) {
	var rules []Rule
	if err := db.Order("id").Find(&rules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// For each rule, count violations in the last 24h
	var result []gin.H
	cutoff := time.Now().Add(-24 * time.Hour)
	for _, rule := range rules {
		var count int64
		db.Model(&RuleViolation{}).Where("rule_id = ? AND created_at >= ?", rule.ID, cutoff).Count(&count)
		result = append(result, gin.H{
			"id":         rule.ID,
			"name":       rule.Name,
			"status":     rule.Status,
			"violations": count,
		})
	}
	c.JSON(http.StatusOK, result)
}

// Add address to suspicious_addresses
func addSuspiciousAddress(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
		Reason  string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address required"})
		return
	}
	addr := SuspiciousAddress{
		Address: req.Address,
		Reason:  req.Reason,
	}
	if err := db.Create(&addr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addr)
}

// Remove address from suspicious_addresses
func removeSuspiciousAddress(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address required"})
		return
	}
	if err := db.Where("address = ?", req.Address).Delete(&SuspiciousAddress{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted", "address": req.Address})
}

// Add address to whitelist_addresses
func addWhitelistAddress(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
		Reason  string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address required"})
		return
	}
	addr := WhitelistAddress{
		Address: req.Address,
		Reason:  req.Reason,
	}
	if err := db.Create(&addr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addr)
}

// Remove address from whitelist_addresses
func removeWhitelistAddress(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address required"})
		return
	}
	if err := db.Where("address = ?", req.Address).Delete(&WhitelistAddress{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted", "address": req.Address})
}

// Get all suspicious addresses
func getSuspiciousAddresses(c *gin.Context) {
	var addresses []SuspiciousAddress
	if err := db.Order("created_at DESC").Find(&addresses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

// Get all whitelist addresses
func getWhitelistAddresses(c *gin.Context) {
	var addresses []WhitelistAddress
	if err := db.Order("created_at DESC").Find(&addresses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

// UpdateRule updates a rule's parameters
func updateRule(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Parameters  string `json:"parameters"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate parameters JSON
	var params map[string]interface{}
	if err := json.Unmarshal([]byte(req.Parameters), &params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters JSON"})
		return
	}

	// Update rule
	result := db.Model(&Rule{}).Where("name = ?", req.Name).Updates(map[string]interface{}{
		"description": req.Description,
		"status":      req.Status,
		"parameters":  req.Parameters,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "rule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated", "rule": req.Name})
}
