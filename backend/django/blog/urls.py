from django.urls import path
from .views.blog import posts, post_details

urlpatterns = [
    path('posts/', posts),
    path('posts/<int:id>/', post_details),
]