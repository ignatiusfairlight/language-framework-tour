import db from '../common/database';

const getAll = async () => {
  return await db.any('SELECT * FROM posts');
};

const getById = () => {
  return { "message": "Hello everynyan!" };
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