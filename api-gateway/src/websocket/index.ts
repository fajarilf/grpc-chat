import { WebSocketServer, WebSocket } from 'ws';
import type { Server as HTTPServer } from 'http';

export function attachWebSocket(server: HTTPServer): WebSocketServer {
  const wss = new WebSocketServer({ server, path: '/ws' });

  wss.on('connection', (ws: WebSocket) => {
    console.log('[ws] client connected');

    ws.send(JSON.stringify({ type: 'welcome', message: 'connected to api-gateway' }));

    ws.on('message', (raw) => {
      const text = raw.toString();
      console.log('[ws] received:', text);

      for (const client of wss.clients) {
        if (client.readyState === WebSocket.OPEN) {
          client.send(JSON.stringify({ type: 'broadcast', data: text }));
        }
      }
    });

    ws.on('close', () => {
      console.log('[ws] client disconnected');
    });

    ws.on('error', (err) => {
      console.error('[ws] error:', err);
    });
  });

  return wss;
}
