import { Router } from 'express';
import type { Request, Response } from 'express';
import { roomRoutes } from './routes/room-routes';
import { userRoutes } from './routes/user-routes';

const router = Router();

router.use('/rooms', roomRoutes);
router.use('/users', userRoutes);

export default router;
