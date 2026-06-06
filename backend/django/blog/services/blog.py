from blog.models.blog import Post

def get_posts():
    result = Post.objects.all()
    return list(result.values())

def get_post(id):
    result = Post.objects.filter(id=id).values().first()
    return result

def create_post():
    post = Post(title="")
    post.save()

# def update_post():

# def delete_post():