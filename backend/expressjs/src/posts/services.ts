import db from '../common/database';

const getAll = async () => {
  return await db.any('SELECT * FROM posts');
};

const getById = async (id: number) => {
  return await db.one('SELECT * FROM posts WHERE id = $1', [id]);
};

const createPost = () => {
  return { "message": "How are you?" };
};

const editPost = () => {
  return { "message": "Fine, thank you!" };
};

const deletePost = () => {
  return { "message": "Oh my gah!" };
};

export default { getAll, getById, createPost, editPost, deletePost };