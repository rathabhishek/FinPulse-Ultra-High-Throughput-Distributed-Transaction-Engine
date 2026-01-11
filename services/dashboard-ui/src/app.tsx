import React, { useEffect, useState } from 'react';

interface Transaction {
  transaction_id: string;
  amount: number;
  type: string;
}

const Dashboard: React.FC = () => {
  const [alerts, setAlerts] = useState<Transaction[]>([]);

  // Reflects your experience building interfaces for asset management [cite: 13]
  return (
    <div style={{ padding: '20px', fontFamily: 'Arial' }}>
      <h1>FinPulse Real-Time Monitor</h1>
      <div style={{ border: '1px solid red', padding: '10px' }}>
        <h3>🚨 High-Value Alerts (Live)</h3>
        {alerts.length === 0 ? <p>Scanning stream...</p> : 
          alerts.map(tx => (
            <div key={tx.transaction_id}>
              FLAGGED: {tx.type} - ${tx.amount.toLocaleString()}
            </div>
          ))
        }
      </div>
    </div>
  );
};

export default Dashboard;