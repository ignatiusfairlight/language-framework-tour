using Microsoft.AspNetCore.Mvc;
using BlogAPI.Services;
using BlogAPI.Models;

namespace BlogAPI.Controllers;

[ApiController]
[Route("api/posts")]
public class PostController : ControllerBase
{
    private readonly PostService _postService;

    public PostController(PostService postService)
    {
        _postService = postService;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<Post>>> Index()
    {
        var posts = await _postService.GetAll();
        return Ok(posts);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<Post>> Show(int id)
    {
        var post = await _postService.GetById(id);
        if (post == null) return NotFound();
        return Ok(post);
    }

    [HttpPost]
    public async Task<ActionResult<Post>> Store(Post post)
    {
        var created = await _postService.Create(post);
        return CreatedAtAction(nameof(Show), new { id = created.Id }, created);           
    }

    [HttpPatch("{id}")]
    public async Task<ActionResult<Post>> Update(int id, UpdatePost post)
    {
        var updated = await _postService.Update(id, post);
        if (updated == null) return NotFound();
        return Ok(updated);
    }
    
    [HttpDelete("{id}")]
    public async Task<ActionResult<Post>> Destroy(int id)
    {
        var deleted = await _postService.Delete(id);
        if (!deleted) return NotFound();
        return NoContent();           
    } 
}