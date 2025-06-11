export function formatTokenAmount(rawAmount: string | number, decimals = 18): string {
  const amount = typeof rawAmount === 'string' ? BigInt(rawAmount) : BigInt(rawAmount);
  const divisor = BigInt(10 ** decimals);
  const whole = amount / divisor;
  const fraction = amount % divisor;
  // Show up to 4 decimal places for readability
  const fractionStr = (fraction * BigInt(10000) / divisor).toString().padStart(4, '0');
  return `${whole}.${fractionStr}`;
} 