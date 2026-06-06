from rest_framework.decorators import api_view
from rest_framework.response import Response
from blog.services.blog import get_posts, get_post

@api_view(['GET', 'POST'])
def posts(request):
    if request.method == 'GET':
        result = get_posts()
        return Response(result)
    elif request.method == 'POST':
        return Response({'message': 'How are you? Fine, thank you.'})

@api_view(['GET', 'PATCH', 'DELETE'])
def post_details(request, id):
    if request.method == 'GET':
        result = get_post(id)
        return Response(result)
    elif request.method == 'PATCH':
        return Response({'message': 'I like doggos too'})
    elif request.method == 'DELETE':
        return Response({'message': 'But I like rabbits more.'})