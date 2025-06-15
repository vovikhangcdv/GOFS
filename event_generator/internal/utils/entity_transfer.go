package utils

// EntityType represents the type of entity in the system
type EntityType uint8

const (
	UNKNOWN EntityType = iota
	INDIV
	PRIV_ENT
	LLC_ONE
	LLC_MULTI
	JSC
	PARTNER
	COOP
	COOP_UNI
	COOP_GRP
	HOUSE_BIZ
	FOR_INDIV
	FOR_ORG
	FOR_ECO
	EXCHANGE_PORTAL
)

// TransactionType represents the type of transaction in the system
type TransactionType uint8

const (
	UNKNOWN_TX TransactionType = iota
	PAYMENT
	GIVE
	LOAN
	INVESTMENT
	EXCHANGE
)

// entityTransferPolicy is a lookup table for entity type transfer permissions
// The table is based on the deployment log matrix where:
// - First index is the "from" entity type
// - Second index is the "to" entity type
// - Value is true if transfer is allowed, false otherwise
var entityTransferPolicy = [][]bool{
	// UNKNOWN
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	// INDIV
	{false, true, true, true, true, true, true, true, false, false, true, false, false, false, true},
	// PRIV_ENT
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// LLC_ONE
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// LLC_MULTI
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// JSC
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// PARTNER
	{false, true, true, true, true, true, true, true, false, false, false, false, false, false, false},
	// COOP
	{false, true, false, false, false, false, false, true, true, true, false, false, false, false, false},
	// COOP_UNI
	{false, true, false, false, false, false, false, true, true, true, false, false, false, false, false},
	// COOP_GRP
	{false, true, false, false, false, false, false, true, true, true, false, false, false, false, false},
	// HOUSE_BIZ
	{false, true, true, false, false, false, false, false, false, false, true, false, false, false, false},
	// FOR_INDIV
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// FOR_ORG
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// FOR_ECO
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	// EXCHANGE_PORTAL
	{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
}

// txTypeAllowedFromTo maps transaction types to their allowed from/to entity types
var txTypeAllowedFromTo = map[TransactionType]struct {
	fromTypes []EntityType
	toTypes   []EntityType
}{
	UNKNOWN_TX: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, HOUSE_BIZ, UNKNOWN, EXCHANGE_PORTAL},
		toTypes:   []EntityType{INDIV, PRIV_ENT, HOUSE_BIZ, UNKNOWN, EXCHANGE_PORTAL},
	},
	PAYMENT: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI, COOP_GRP, HOUSE_BIZ, FOR_INDIV, FOR_ORG, FOR_ECO, UNKNOWN},
		toTypes:   []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI, COOP_GRP, HOUSE_BIZ, FOR_INDIV, FOR_ORG, FOR_ECO, UNKNOWN},
	},
	GIVE: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI, COOP_GRP, HOUSE_BIZ, FOR_INDIV, FOR_ORG, FOR_ECO, EXCHANGE_PORTAL},
		toTypes:   []EntityType{INDIV, HOUSE_BIZ, FOR_INDIV},
	},
	LOAN: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC},
		toTypes:   []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, HOUSE_BIZ},
	},
	INVESTMENT: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, FOR_INDIV, FOR_ORG, FOR_ECO},
		toTypes:   []EntityType{PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI},
	},
	EXCHANGE: {
		fromTypes: []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI, COOP_GRP, HOUSE_BIZ, FOR_INDIV, FOR_ORG, FOR_ECO, EXCHANGE_PORTAL},
		toTypes:   []EntityType{INDIV, PRIV_ENT, LLC_ONE, LLC_MULTI, JSC, PARTNER, COOP, COOP_UNI, COOP_GRP, HOUSE_BIZ, FOR_INDIV, FOR_ORG, FOR_ECO, EXCHANGE_PORTAL},
	},
}

// IsTransferAllowed checks if a transfer is allowed between two entity types
func IsTransferAllowed(fromType, toType EntityType) bool {
	if int(fromType) >= len(entityTransferPolicy) || int(toType) >= len(entityTransferPolicy[0]) {
		return false
	}
	return entityTransferPolicy[fromType][toType]
}

// GetEntityTypeName returns the string representation of an entity type
func GetEntityTypeName(entityType EntityType) string {
	switch entityType {
	case UNKNOWN:
		return "UNKNOWN"
	case INDIV:
		return "INDIVIDUAL"
	case PRIV_ENT:
		return "PRIVATE_ENTERPRISE"
	case LLC_ONE:
		return "LLC_ONE_MEMBER"
	case LLC_MULTI:
		return "LLC_MULTI_MEMBER"
	case JSC:
		return "JOINT_STOCK_COMPANY"
	case PARTNER:
		return "PARTNERSHIP"
	case COOP:
		return "COOPERATIVE"
	case COOP_UNI:
		return "UNION_OF_COOPERATIVES"
	case COOP_GRP:
		return "COOPERATIVE_GROUP"
	case HOUSE_BIZ:
		return "HOUSEHOLD_BUSINESS"
	case FOR_INDIV:
		return "FOREIGN_INDIVIDUAL_INVESTOR"
	case FOR_ORG:
		return "FOREIGN_ORGANIZATION_INVESTOR"
	case FOR_ECO:
		return "FOREIGN_INVESTED_ECONOMIC_ORGANIZATION"
	case EXCHANGE_PORTAL:
		return "EXCHANGE_PORTAL"
	default:
		return "UNKNOWN"
	}
}

// GetTransactionTypeName returns the string representation of a transaction type
func GetTransactionTypeName(txType TransactionType) string {
	switch txType {
	case UNKNOWN_TX:
		return "UNKNOWN"
	case PAYMENT:
		return "PAYMENT"
	case GIVE:
		return "GIVE"
	case LOAN:
		return "LOAN"
	case INVESTMENT:
		return "INVESTMENT"
	case EXCHANGE:
		return "EXCHANGE"
	default:
		return "UNKNOWN"
	}
}

// GetAllowedTransactionTypes returns a list of transaction types that can be used for transfer between two entity types
func GetAllowedTransactionTypes(fromType, toType EntityType) []TransactionType {
	var allowedTypes []TransactionType

	// Check each transaction type
	for txType, policy := range txTypeAllowedFromTo {
		// Check if fromType is allowed
		fromAllowed := false
		for _, allowedFrom := range policy.fromTypes {
			if fromType == allowedFrom {
				fromAllowed = true
				break
			}
		}

		// Check if toType is allowed
		toAllowed := false
		for _, allowedTo := range policy.toTypes {
			if toType == allowedTo {
				toAllowed = true
				break
			}
		}

		// If both from and to types are allowed, add this transaction type
		if fromAllowed && toAllowed {
			allowedTypes = append(allowedTypes, txType)
		}
	}

	return allowedTypes
}
