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

    public function store(string $title, string $content)
    {
        $post = Post::create([
            'title' => $title,
            'content' => $content
        ]);

        return "Post created!";
    }

    public function update()
    {

    }

    public function destroy()
    {

    }
}
