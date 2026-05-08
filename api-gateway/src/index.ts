import 'dotenv/config';
import express from 'express';
import cors from 'cors';
import http from 'http';
import apiRouter from './api/index.js';
import { attachWebSocket } from './websocket/index.js';

const app = express();
app.use(cors());
app.use(express.json());
app.use('/api', apiRouter);

app.get('/', (_req, res) => {
  res.json({ name: 'api-gateway', endpoints: ['/api/health', '/api/hello', '/api/echo', 'ws://.../ws'] });
});

const server = http.createServer(app);
attachWebSocket(server);

const port = Number(process.env['PORT'] ?? 3000);
server.listen(port, () => {
  console.log(`[api-gateway] http listening on http://localhost:${port}`);
  console.log(`[api-gateway] ws listening on ws://localhost:${port}/ws`);
});
