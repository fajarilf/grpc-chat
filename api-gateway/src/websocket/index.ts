import { WebSocketServer, WebSocket } from 'ws';
import type { Server as HTTPServer } from 'http';
import {
  rooms,
  registerSocket,
  unregisterSocket,
  broadcastToRoom,
  roomsForUser,
} from '../store/rooms.js';

type Incoming =
  | { type: 'hello'; username: string }
  | { type: 'message'; room: string; body: string };

export function attachWebSocket(server: HTTPServer): WebSocketServer {
  const wss = new WebSocketServer({ server, path: '/ws' });
  const usernames = new WeakMap<WebSocket, string>();

  const send = (ws: WebSocket, payload: unknown) => {
    if (ws.readyState === WebSocket.OPEN) ws.send(JSON.stringify(payload));
  };

  wss.on('connection', (ws) => {
    send(ws, { type: 'welcome', message: 'send {type:"hello", username:"..."}' });

    ws.on('message', (raw) => {
      let msg: Incoming;
      try {
        msg = JSON.parse(raw.toString()) as Incoming;
      } catch {
        return send(ws, { type: 'error', message: 'invalid JSON' });
      }

      if (msg.type === 'hello') {
        if (usernames.get(ws)) {
          return send(ws, { type: 'error', message: 'already identified' });
        }
        if (!msg.username || typeof msg.username !== 'string') {
          return send(ws, { type: 'error', message: 'username required' });
        }
        usernames.set(ws, msg.username);
        registerSocket(msg.username, ws);
        return send(ws, { type: 'hello_ok', username: msg.username });
      }

      const username = usernames.get(ws);
      if (!username) {
        return send(ws, { type: 'error', message: 'send hello first' });
      }

      if (msg.type === 'message') {
        if (!msg.room || typeof msg.body !== 'string') {
          return send(ws, { type: 'error', message: 'room and body required' });
        }
        const members = rooms.get(msg.room);
        if (!members || !members.has(username)) {
          return send(ws, { type: 'error', message: 'join the room first (POST /api/rooms/:roomId/join)' });
        }
        broadcastToRoom(msg.room, {
          type: 'message',
          room: msg.room,
          from: username,
          body: msg.body,
          at: Date.now(),
        });
        return;
      }

      send(ws, { type: 'error', message: 'unknown type' });
    });

    ws.on('close', () => {
      const username = usernames.get(ws);
      if (!username) return;
      const lastSocket = unregisterSocket(username, ws);
      if (!lastSocket) return; // user still has other tabs open

      // user fully disconnected → leave every room and notify others
      for (const roomId of roomsForUser(username)) {
        const members = rooms.get(roomId);
        if (!members) continue;
        members.delete(username);
        broadcastToRoom(roomId, { type: 'user_leaving', room: roomId, username });
      }
    });

    ws.on('error', (err) => console.error('[ws] error:', err));
  });

  return wss;
}
