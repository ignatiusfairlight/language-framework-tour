from rest_framework.decorators import api_view
from rest_framework.response import Response
from blog.serializers.blog import PostSerializer
from blog.services.blog import get_posts, get_post, create_post, delete_post

@api_view(['GET', 'POST'])
def posts(request):
    if request.method == 'GET':
        result = get_posts()
        return Response(result)
    elif request.method == 'POST':
        serializer = PostSerializer(data=request.data)
        if serializer.is_valid():
            create_post(serializer.validated_data)
            return Response()
        return Response(serializer.errors, status=400)

@api_view(['GET', 'PATCH', 'DELETE'])
def post_details(request, id):
    if request.method == 'GET':
        result = get_post(id)
        return Response(result)
    elif request.method == 'PATCH':
        return Response({'message': 'I like doggos too'})
    elif request.method == 'DELETE':
        delete_post(id)
        return Response()