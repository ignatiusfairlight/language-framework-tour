from sqlmodel import select
from sqlmodel.ext.asyncio.session import AsyncSession
from .models import Post
from .schemas import PostCreate, PostEdit

async def get_all_posts(session: AsyncSession) -> list[Post]:
    result = await session.exec(select(Post).order_by(Post.id))
    return result.all()

async def get_by_id(session: AsyncSession, id: int):
    result = await session.exec(select(Post).where(Post.id == id))
    return result.first()

async def create_post(session: AsyncSession, data: PostCreate):
    post = Post.model_validate(data)
    session.add(post)
    await session.commit()
    await session.refresh(post)
    return post

async def update_post():
    return {"message": "Fine, thank you!"}

async def delete_post(session: AsyncSession, id: int):
    result = (await session.exec(select(Post).where(Post.id == id))).first()
    if result is None:
        return False
    await session.delete(result)
    await session.commit()
    return True