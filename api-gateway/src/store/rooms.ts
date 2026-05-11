import { WebSocket } from 'ws';

// roomId -> set of usernames currently in the room
export const rooms = new Map<string, Set<string>>();

// username -> set of active websocket connections for that user
export const conns = new Map<string, Set<WebSocket>>();

export function registerSocket(username: string, ws: WebSocket): void {
  let set = conns.get(username);
  if (!set) {
    set = new Set();
    conns.set(username, set);
  }
  set.add(ws);
}

// returns true if this was the user's last open socket
export function unregisterSocket(username: string, ws: WebSocket): boolean {
  const set = conns.get(username);
  if (!set) return false;
  set.delete(ws);
  if (set.size === 0) {
    conns.delete(username);
    return true;
  }
  return false;
}

export function sendToUser(username: string, payload: unknown): void {
  const sockets = conns.get(username);
  if (!sockets) return;
  const data = JSON.stringify(payload);
  for (const ws of sockets) {
    if (ws.readyState === WebSocket.OPEN) ws.send(data);
  }
}

export function broadcastToRoom(roomId: string, payload: unknown, except?: string): void {
  const members = rooms.get(roomId);
  if (!members) return;
  for (const username of members) {
    if (username === except) continue;
    sendToUser(username, payload);
  }
}

export function roomsForUser(username: string): string[] {
  const out: string[] = [];
  for (const [roomId, members] of rooms) {
    if (members.has(username)) out.push(roomId);
  }
  return out;
}
