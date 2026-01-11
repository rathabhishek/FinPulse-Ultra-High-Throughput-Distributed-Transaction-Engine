const express = require('express');
const http = require('http');
const WebSocket = require('ws');

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

// Logic to broadcast Flink alerts to the React UI
wss.on('connection', (ws) => {
    console.log('Dashboard client connected');
});

// Endpoint for Flink to POST anomalies to
app.post('/api/alerts', express.json(), (req, res) => {
    wss.clients.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(JSON.stringify(req.body));
        }
    });
    res.status(200).send('Alert Broadcasted');
});

server.listen(3001, () => console.log('Relay Server running on port 3001'));