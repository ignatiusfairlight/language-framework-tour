from blog.models.blog import Post

def get_posts():
    result = Post.objects.all()
    return list(result.values())

def get_post(id):
    result = Post.objects.filter(id=id).values().first()
    return result

def create_post(data):
    return Post.objects.create(**data)

def update_post(id, data):
    return Post.objects.filter(id=id).update(**data)

def delete_post(id):
    return Post.objects.filter(id=id).delete()