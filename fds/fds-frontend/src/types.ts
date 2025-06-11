export interface AddressTotals {
  address: string;
  in: number;
  out: number;
}

export interface SuspiciousTransfer {
  id: number;
  from_address: string;
  to_address: string;
  amount: string;
  txHash: string;
  blockNumber: number;
  timestamp: string;
  reason: string;
  severity: string;
  details: string;
  isBlacklisted: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface BlacklistedAddress {
  id: number;
  address: string;
  txHash: string;
  blockNumber: number;
  reason: string;
  severity: string;
  details: string;
  createdAt: string;
  updatedAt: string;
}

export interface RelatedAddresses {
  address: string;
  related: string[];
}

export interface Rule {
  id: number;
  name: string;
  status: string;
  violations: number;
}

export interface RuleViolation {
  id: number;
  rule_id: number;
  tx_hash: string;
  block_number: number;
  created_at: string;
}

export interface SuspiciousAddress {
  id: number;
  address: string;
  reason: string;
  created_at: string;
}

export interface WhitelistAddress {
  id: number;
  address: string;
  reason: string;
  created_at: string;
} 