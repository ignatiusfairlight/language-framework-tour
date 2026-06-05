<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Attributes\Table;
use Illuminate\Database\Eloquent\Model;

#[Table(key: 'id')]
class Post extends Model
{
    protected $fillable = [
        'title',
        'content'
    ];
}
