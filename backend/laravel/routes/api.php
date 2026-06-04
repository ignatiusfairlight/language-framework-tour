<?php

use App\Http\Controllers\Blog\BlogController;
use Illuminate\Support\Facades\Route;

Route::get('/posts', [BlogController::class, '']);
