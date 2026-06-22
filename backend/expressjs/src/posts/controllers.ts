import { NextFunction, Request, Response, Router } from "express";
import postService from './services';

const router = Router();

router.get('/posts', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.getAll();
  res.json(result);
});

router.get('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.getById();
  res.json(result);
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