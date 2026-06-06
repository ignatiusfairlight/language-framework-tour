import json 
from django.http import JsonResponse

class ApiResponseMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        response = self.get_response(request)

        if request.path.startswith('/api'):
            result = json.loads(response.content)
            wrapped = {
                'success' : True if response.status_code == 200 else False,
                'message' : 'Success' if response.status_code == 200 else 'Fail',
                'data' : result
            }
            return JsonResponse(wrapped)

        return response