from sqlmodel import select
from sqlmodel.ext.asyncio.session import AsyncSession
from .models import Post

async def get_all_posts(session: AsyncSession) -> list[Post]:
    result = await session.exec(select(Post))
    return result.all()