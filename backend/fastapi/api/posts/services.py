from sqlmodel import select
from sqlmodel.ext.asyncio.session import AsyncSession
from .models import Post

async def get_all_posts(session: AsyncSession) -> list[Post]:
    result = await session.exec(select(Post))
    return result.all()

async def get_by_id():
    return await {"message": "Hello everynyan!"}

async def create_post():
    return await {"message": "How are you?"}

async def update_post():
    return await {"message": "Fine, thank you!"}

async def delete_post():
    return await {"message": "Oh my gah"}