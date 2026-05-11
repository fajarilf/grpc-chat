import { Router } from 'express';
import type { Request, Response } from 'express';
import { roomRoutes } from './routes/room-routes';

const router = Router();

router.use('/rooms', roomRoutes);

router.get('/health', (_req: Request, res: Response) => {
  res.json({ status: 'ok', uptime: process.uptime() });
});

router.get('/hello', (req: Request, res: Response) => {
  const name = (req.query['name'] as string | undefined) ?? 'world';
  res.json({ message: `Hello, ${name}!` });
});

router.post('/echo', (req: Request, res: Response) => {
  res.json({ received: req.body });
});

export default router;
