import { NextFunction, Request, Response, Router } from "express";
import postService from './services';

const router = Router();

router.get('/posts', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.getAll();
  res.json(result);
});

router.get('/posts/:id', async (req: Request, res: Response, next: NextFunction) => {
  try {
    const result = await postService.getById(Number(req.params.id));
    res.json(result);
  } catch(error) {
    next(error)
  }
});

router.post('/posts', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.createPost();
  res.json(result);
});

router.patch('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.editPost();
  res.json(result);
});

router.delete('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.deletePost();
  res.json(result);
});

export default router;