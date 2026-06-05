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

    public function update()
    {

    }

    public function destroy()
    {

    }
}
