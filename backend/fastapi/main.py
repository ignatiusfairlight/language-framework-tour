from fastapi import FastAPI
from contextlib import asynccontextmanager
from sqlmodel import SQLModel

from api import router as api_router
from core.config import settings
from core.database import engine

@asynccontextmanager
async def lifespan(app: FastAPI):
    async with engine.begin() as conn:
        await conn.run_sync(SQLModel.metadata.create_all)
    yield

app = FastAPI(lifespan=lifespan)

app.include_router(api_router, prefix=settings.api_prefix)