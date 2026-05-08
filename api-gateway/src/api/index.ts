import { Router } from 'express';
import type { Request, Response } from 'express';

const router = Router();

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
