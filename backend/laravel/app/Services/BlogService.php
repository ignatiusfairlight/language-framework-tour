<?php

namespace App\Services;

use App\Models\Post;

class BlogService
{
    public function index()
    {
        return Post::all();
    }

    public function show(int $id)
    {
        return Post::where('id', $id)->first();
    }

    public function store(array $input)
    {
        return Post::create([
            'title' => $input['title'],
            'content' => $input['content']
        ]);
    }

    public function update(array $input, int $id)
    {
        $post = Post::find($id);

        if ($post == null) {
            return null;
        } else {
            $post->update($input);
            return $post;
        }

        
    }

    public function destroy(int $id)
    {
        $post = Post::find($id);

        if ($post == null) {
            return null;
        } else {
            return $post->delete();
        }
    }
}
