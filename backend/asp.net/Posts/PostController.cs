using Microsoft.AspNetCore.Mvc;
using BlogAPI.Services;

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
    public ActionResult<string> Index()
    {
        var posts = "Hello world!";
        return Ok(posts);
    }

    [HttpGet("{id}")]
    public ActionResult<string> Show(int id)
    {
        var posts = "I will drink more water.";
        return Ok(posts);
    }

    [HttpPost]
    public ActionResult<string> Store()
    {
        var posts = "Always move forward!";
        return Ok(posts);           
    }

    [HttpPatch("{id}")]
    public ActionResult<string> Update()
    {
        var posts = "I like cheese!";
        return Ok(posts);           
    }

    [HttpDelete("{id}")]
    public ActionResult<string> Destroy()
    {
        var posts = "Yes King!";
        return Ok(posts);           
    } 
}