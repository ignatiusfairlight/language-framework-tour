from rest_framework.decorators import api_view
from rest_framework.response import Response

@api_view(['GET', 'POST'])
def posts(request):
    if request.method == 'GET':
        return Response({'message': 'Hello evernyan!'})
    elif request.method == 'POST':
        return Response({'message': 'How are you? Fine, thank you.'})

@api_view(['GET', 'PATCH', 'DELETE'])
def post_details(request, id):
    if request.method == 'GET':
        return Response({'message': 'I like cats'})
    elif request.method == 'PATCH':
        return Response({'message': 'I like doggos too'})
    elif request.method == 'DELETE':
        return Response({'message': 'But I like rabbits more.'})