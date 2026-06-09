import { Router } from 'express';
import type { Request, Response } from 'express';
import { roomRoutes } from './routes/room-routes';

const router = Router();

router.use('/rooms', roomRoutes);

export default router;
