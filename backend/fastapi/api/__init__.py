from fastapi import APIRouter

from .posts.routes import router as posts_router

router = APIRouter()
router.include_router(posts_router, tags=["posts"], prefix="/posts")