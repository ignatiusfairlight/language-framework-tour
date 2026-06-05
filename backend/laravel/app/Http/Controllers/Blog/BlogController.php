<?php

namespace App\Http\Controllers\Blog;

use App\Http\Controllers\APIController;
use App\Services\BlogService;
use Illuminate\Http\Request;

class BlogController extends APIController
{
    protected BlogService $blogService;

    public function __construct()
    {
        $this->blogService = new BlogService;
    }

    public function index()
    {
        return $this->success($this->blogService->index());
    }

    public function show(int $id)
    {
        return $this->success($this->blogService->show($id));
    }

    // Create one blog
    public function store()
    {

    }
    
    // Edit blog
    public function update()
    {

    }

    // Delete blog
    public function destroy()
    {

    }
}
