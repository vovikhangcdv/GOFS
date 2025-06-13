export function formatTokenAmount(rawAmount: string | number | null | undefined, decimals = 18): string {
  try {
    if (rawAmount === null || rawAmount === undefined) {
      return '0.0000';
    }
    const amount = typeof rawAmount === 'string' ? BigInt(rawAmount) : BigInt(rawAmount);
    const divisor = BigInt(10 ** decimals);
    const whole = amount / divisor;
    const fraction = amount % divisor;
    // Show up to 4 decimal places for readability
    const fractionStr = (fraction * BigInt(10000) / divisor).toString().padStart(4, '0');
    return `${whole}.${fractionStr}`;
  } catch (error) {
    console.error('Error formatting token amount:', error);
    return '0.0000';
  }
} 