import 'dotenv/config';
import express from 'express';
import cors from 'cors';
import http from 'http';
import path from 'path';
import SwaggerParser from '@apidevtools/swagger-parser';
import apiRouter from './api/index';
import { attachWebSocket } from './websocket/index';
import { apiReference } from '@scalar/express-api-reference';

async function bootstrap() {
  const app = express();
  app.use(cors());
  app.use(express.json());
  app.use('/api', apiRouter);

  // Bundle the split OpenAPI YAML (main-docs.yaml + its external $refs) into a
  // single resolved document. Scalar does not reliably resolve external
  // path-item $refs on its own, so we resolve them server-side and hand it the
  // finished spec via `content`.
  const openapiSpec = await SwaggerParser.bundle(
    path.join(__dirname, 'api/docs/main-docs.yaml')
  );

  app.use('/docs', apiReference({
    content: openapiSpec,
    layout: 'classic',
  }));

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
}

bootstrap().catch((err) => {
  console.error('[api-gateway] failed to start:', err);
  process.exit(1);
});
