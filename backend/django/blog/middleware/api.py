import json 
from django.http import JsonResponse

class ApiResponseMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    # Refactor this
    def __call__(self, request):
        response = self.get_response(request)

        if request.path.startswith('/api'):
            content_type = response.get('Content-Type', '')
            if 'application/json' not in content_type:
                return response

            if request.method == 'GET':
                result = json.loads(response.content)
                wrapped = {
                    'success' : True if response.status_code < 400 else False,
                    'message' : 'Success' if response.status_code < 400 else 'Fail',
                    'data' : result
                }
            else:
                wrapped = {
                    'success' : True if response.status_code < 400 else False,
                    'message' : 'Success' if response.status_code < 400 else 'Fail',
                }
            return JsonResponse(wrapped)

        return response