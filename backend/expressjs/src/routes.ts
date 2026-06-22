import { Router } from "express";
import postsController from './posts/controllers';

const api = Router()
  .use(postsController);

export default Router().use('/api', api);