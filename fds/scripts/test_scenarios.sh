#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color
BLUE='\033[0;34m'

# Function to print section headers
print_header() {
    echo -e "\n${BLUE}=== $1 ===${NC}\n"
}

# Function to check if a command exists
check_command() {
    if ! command -v $1 &> /dev/null; then
        echo -e "${RED}Error: $1 is not installed${NC}"
        exit 1
    fi
}

# Check required commands
check_command cast
check_command jq

# Load environment variables
if [ -f .env ]; then
    source .env
else
    echo -e "${RED}Error: .env file not found${NC}"
    exit 1
fi

# Function to run large transfer scenario
run_large_transfer() {
    print_header "Scenario 1: Large Transfer"
    echo "Creating large transfer..."
    cast send --private-key $PRIVATE_KEY --rpc-url $RPC_URL \
        $TOKENX_ADDRESS \
        "transfer(address,uint256)" \
        $TEST_ADDRESS "1000000000000000000000000"
}

# Function to run multiple transfers scenario
run_multiple_transfers() {
    print_header "Scenario 2: Multiple Transfers in Short Time"
    for i in {1..5}; do
        echo "Transfer $i..."
        cast send --private-key $PRIVATE_KEY --rpc-url $RPC_URL \
            $TOKENX_ADDRESS \
            "transfer(address,uint256)" \
            $TEST_ADDRESS "100000000000000000000000"
        sleep 1
    done
}

# Function to run multiple incoming transfers scenario
run_incoming_transfers() {
    print_header "Scenario 3: Multiple Incoming Transfers"
    for i in {1..3}; do
        echo "Incoming transfer $i..."
        cast send --private-key $PRIVATE_KEY --rpc-url $RPC_URL \
            $TOKENX_ADDRESS \
            "transfer(address,uint256)" \
            $RECEIVER_ADDRESS "400000000000000000000"
        sleep 20
        print_header "================"
    done
}

# Function to run suspicious address scenario
run_suspicious_address() {
    print_header "Scenario 4: Suspicious Address Interaction"
    echo "Transfer to suspicious address..."
    cast send --private-key $PRIVATE_KEY --rpc-url $RPC_URL \
        $TOKENX_ADDRESS \
        "transfer(address,uint256)" \
        $SUSPICIOUS_ADDRESS "100000000000000000000000"
}

# Function to run all scenarios
run_all_scenarios() {
    run_large_transfer
    run_multiple_transfers
    run_incoming_transfers
    run_suspicious_address
    echo -e "\n${GREEN}All scenarios completed${NC}"
}

# Check if contract address is set
if [ -z "$TOKENX_ADDRESS" ]; then
    echo -e "${RED}Error: TOKENX_ADDRESS not set in .env${NC}"
    exit 1
fi

# Main script
case "$1" in
    "large")
        run_large_transfer
        ;;
    "multiple")
        run_multiple_transfers
        ;;
    "incoming")
        run_incoming_transfers
        ;;
    "suspicious")
        run_suspicious_address
        ;;
    "all")
        run_all_scenarios
        ;;
    *)
        echo "Usage: $0 {large|multiple|incoming|suspicious|all}"
        echo "  large: Run large transfer scenario"
        echo "  multiple: Run multiple transfers scenario"
        echo "  incoming: Run multiple incoming transfers scenario"
        echo "  suspicious: Run suspicious address scenario"
        echo "  all: Run all scenarios"
        exit 1
        ;;
esac 