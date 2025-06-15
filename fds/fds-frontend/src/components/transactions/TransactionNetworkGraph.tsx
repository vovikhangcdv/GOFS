import React, { useCallback, useEffect, useState } from 'react';
import {
  ReactFlow,
  Node,
  Edge,
  addEdge,
  Connection,
  useNodesState,
  useEdgesState,
  Controls,
  Background,
  NodeTypes,
  Handle,
  Position,
  MarkerType,
} from 'reactflow';
import 'reactflow/dist/style.css';
import { Box, Card, CardContent, Typography, styled } from '@mui/material';

// Custom node component for addresses
const AddressNode = ({ data }: { data: any }) => {
  const isTarget = data.isTarget;
  const isSource = data.isSource;
  const isDestination = data.isDestination;
  
  // Determine background and text color for better contrast
  const getNodeStyle = () => {
    if (isTarget) {
      return {
        background: 'rgba(59, 130, 246, 0.95)',
        border: '2px solid #3b82f6',
        textColor: 'white',
      };
    } else if (isSource) {
      return {
        background: 'rgba(168, 85, 247, 0.95)',
        border: '2px solid #a855f7',
        textColor: 'white',
      };
    } else if (isDestination) {
      return {
        background: 'rgba(34, 197, 94, 0.95)',
        border: '2px solid #22c55e',
        textColor: 'white',
      };
    } else {
      return {
        background: 'rgba(30, 41, 59, 0.95)',
        border: '2px solid #64748b',
        textColor: 'white',
      };
    }
  };

  const nodeStyle = getNodeStyle();
  
  return (
    <Card
      sx={{
        minWidth: 280,
        maxWidth: 320,
        background: nodeStyle.background,
        backdropFilter: 'blur(10px)',
        border: nodeStyle.border,
        borderRadius: 2,
        boxShadow: '0 8px 32px rgba(0, 0, 0, 0.4)',
      }}
    >
      <Handle type="target" position={Position.Top} />
      <CardContent sx={{ p: 2, '&:last-child': { pb: 2 } }}>
        <Typography variant="subtitle2" sx={{ 
          color: nodeStyle.textColor, 
          fontWeight: 700,
          fontSize: '0.85rem',
          mb: 1,
          textAlign: 'center',
        }}>
          {data.label}
        </Typography>
        
        <Typography variant="body2" sx={{ 
          color: nodeStyle.textColor, 
          fontFamily: 'monospace',
          fontSize: '0.75rem',
          wordBreak: 'break-all',
          backgroundColor: 'rgba(0, 0, 0, 0.2)',
          padding: '4px 8px',
          borderRadius: 1,
          mb: 1,
        }}>
          {data.address}
        </Typography>
        
        {data.balance && (
          <Typography variant="caption" sx={{ 
            color: nodeStyle.textColor, 
            display: 'block',
            fontSize: '0.7rem',
            fontWeight: 600,
            textAlign: 'center',
            backgroundColor: 'rgba(0, 0, 0, 0.1)',
            padding: '2px 4px',
            borderRadius: 1,
          }}>
            {data.balance}
          </Typography>
        )}
      </CardContent>
      <Handle type="source" position={Position.Bottom} />
    </Card>
  );
};

const nodeTypes: NodeTypes = {
  addressNode: AddressNode,
};

interface Transaction {
  id: string;
  txHash: string;
  from_address: string;
  to_address: string;
  amount: string;
  blockNumber: number;
  timestamp?: string;
}

interface TransactionNetworkGraphProps {
  targetAddress: string;
  transactions: Transaction[];
  ethBalance?: { balance: number; unit: string };
  evndBalance?: { balance: number; unit: string };
}

const StyledReactFlow = styled(ReactFlow)(({ theme }) => ({
  background: 'transparent',
  '& .react-flow__attribution': {
    display: 'none',
  },
  '& .react-flow__controls': {
    background: 'rgba(30, 41, 59, 0.9)',
    border: '1px solid rgba(100, 116, 139, 0.3)',
    borderRadius: 8,
    '& button': {
      background: 'transparent',
      color: theme.palette.text.primary,
      border: 'none',
      '&:hover': {
        background: 'rgba(100, 116, 139, 0.2)',
      },
    },
  },
  '& .react-flow__edge': {
    '& .react-flow__edge-path': {
      stroke: '#64748b',
      strokeWidth: 2,
    },
    '& .react-flow__edge-text': {
      fill: '#ffffff',
      fontSize: '11px',
      fontWeight: 700,
      textShadow: '0 0 4px rgba(0, 0, 0, 0.8), 0 0 8px rgba(0, 0, 0, 0.6)',
    },
    '& .react-flow__edge-textbg': {
      fill: 'rgba(30, 41, 59, 0.9)',
      stroke: 'rgba(100, 116, 139, 0.5)',
      strokeWidth: 1,
      rx: 4,
      ry: 4,
    },
  },
}));

export const TransactionNetworkGraph: React.FC<TransactionNetworkGraphProps> = ({
  targetAddress,
  transactions,
  ethBalance,
  evndBalance,
}) => {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  const formatAmount = (amount: string | number): string => {
    const num = Number(amount) / 1e18;
    return num.toFixed(4);
  };

  const formatBalance = (balance?: { balance: number; unit: string }): string => {
    if (!balance) return '';
    return `${Number(balance.balance).toFixed(4)} ${balance.unit}`;
  };

  useEffect(() => {
    if (!transactions.length) return;

    // Get unique addresses
    const addressSet = new Set<string>();
    addressSet.add(targetAddress);
    
    transactions.forEach(tx => {
      if (tx.from_address) addressSet.add(tx.from_address);
      if (tx.to_address) addressSet.add(tx.to_address);
    });

    const addresses = Array.from(addressSet);
    
    // Create nodes for each unique address
    const newNodes: Node[] = addresses.map((address, index) => {
      const isTarget = address === targetAddress;
      const isSource = transactions.some(tx => tx.from_address === address && tx.to_address === targetAddress);
      const isDestination = transactions.some(tx => tx.from_address === targetAddress && tx.to_address === address);
      
      // Calculate position in a network layout with proper spacing
      const centerX = 500;
      const centerY = 300;
      const radius = 350; // Increased radius for better spacing
      
      let x, y;
      
      if (isTarget) {
        // Place target address in the center
        x = centerX;
        y = centerY;
      } else {
        // Place other addresses in a circle around the center with proper spacing
        const nonTargetAddresses = addresses.filter(addr => addr !== targetAddress);
        const addressIndex = nonTargetAddresses.indexOf(address);
        const totalNonTarget = nonTargetAddresses.length;
        
        if (totalNonTarget === 1) {
          // Single node - place it to the right with good spacing
          x = centerX + radius;
          y = centerY;
        } else if (totalNonTarget === 2) {
          // Two nodes - place them on left and right
          x = centerX + (addressIndex === 0 ? -radius : radius);
          y = centerY;
        } else {
          // Multiple nodes - distribute them evenly in a circle with proper spacing
          const angle = (addressIndex * 2 * Math.PI) / totalNonTarget;
          x = centerX + radius * Math.cos(angle);
          y = centerY + radius * Math.sin(angle);
        }
      }

      const nodeBalance = isTarget && ethBalance && evndBalance 
        ? `ETH: ${formatBalance(ethBalance)}, eVND: ${formatBalance(evndBalance)}`
        : undefined;

      return {
        id: address,
        type: 'addressNode',
        position: { x, y },
        data: {
          label: isTarget ? 'Target Address' : isSource ? 'Source' : isDestination ? 'Destination' : 'Related',
          address,
          balance: nodeBalance,
          isTarget,
          isSource: isSource && !isTarget,
          isDestination: isDestination && !isTarget,
        },
      };
    });

    // Create edges for transactions
    const newEdges: Edge[] = transactions.map((tx, index) => {
      const isOutgoing = tx.from_address === targetAddress;
      const isIncoming = tx.to_address === targetAddress;
      
      return {
        id: `${tx.from_address}-${tx.to_address}-${index}`,
        source: tx.from_address,
        target: tx.to_address,
        type: 'straight',
        animated: true,
        label: `${formatAmount(tx.amount)} eVND`,
        style: {
          stroke: isOutgoing ? '#f59e0b' : isIncoming ? '#10b981' : '#6b7280',
          strokeWidth: 4,
        },
        markerEnd: {
          type: MarkerType.ArrowClosed,
          color: isOutgoing ? '#f59e0b' : isIncoming ? '#10b981' : '#6b7280',
          width: 20,
          height: 20,
        },
        labelStyle: {
          fill: '#ffffff',
          fontWeight: 700,
          fontSize: '12px',
          textShadow: '0 0 4px rgba(0, 0, 0, 0.8)',
        },
        labelBgStyle: {
          fill: 'rgba(30, 41, 59, 0.95)',
          stroke: 'rgba(100, 116, 139, 0.6)',
          strokeWidth: 1,
          rx: 6,
          ry: 6,
          fillOpacity: 0.95,
        },
        data: {
          txHash: tx.txHash,
          amount: tx.amount,
          blockNumber: tx.blockNumber,
          timestamp: tx.timestamp,
        },
      };
    });

    setNodes(newNodes);
    setEdges(newEdges);
  }, [targetAddress, transactions, ethBalance, evndBalance, setNodes, setEdges]);

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge(params, eds)),
    [setEdges]
  );

  if (!transactions.length) {
    return (
      <Box sx={{ 
        height: 400, 
        display: 'flex', 
        alignItems: 'center', 
        justifyContent: 'center',
        background: 'rgba(30, 41, 59, 0.3)',
        borderRadius: 2,
        border: '1px solid rgba(114, 127, 77, 0.2)',
      }}>
        <Typography variant="body2" color="text.secondary">
          No transactions to visualize
        </Typography>
      </Box>
    );
  }

  return (
    <Box sx={{ 
      height: 700, 
      width: '100%',
      background: 'rgba(30, 41, 59, 0.3)',
      borderRadius: 2,
      border: '1px solid rgba(100, 116, 139, 0.2)',
      overflow: 'hidden',
    }}>
      <StyledReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        nodeTypes={nodeTypes}
        fitView
        fitViewOptions={{ padding: 0.2 }}
      >
        <Controls />
        <Background 
          gap={20} 
          size={1} 
          color="rgba(100, 116, 139, 0.2)"
        />
      </StyledReactFlow>
    </Box>
  );
}; 