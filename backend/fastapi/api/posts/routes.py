from fastapi import APIRouter, Depends
from sqlmodel.ext.asyncio.session import AsyncSession
from core.database import get_session, SessionDep
from . import services
from .schemas import PostRead

router = APIRouter()

@router.get("/", response_model=list[PostRead])
async def show_all(session: SessionDep):
    return await services.get_all_posts(session)

@router.get("/{id}")
async def show_one(id: int):
    return {"message": "Hello everynyan!"}

@router.post("/")
async def create():
    return {"message": "How are you?"}

@router.patch("/{id}")
async def update(id: int):
    return {"message": "Fine, thank you!"}

@router.delete("/{id}")
async def delete(id: int):
    return {"message": "Oh my gah"}