using Microsoft.EntityFrameworkCore;
using BlogAPI.Models;

namespace BlogAPI.Services;

public class PostService
{
    private readonly PostContext _context;

    public PostService(PostContext context)
    {
        _context = context;
    }

    public async Task<List<Post>> GetAll()
    {
        return await _context.Posts.ToListAsync();
    }

    public async Task<Post?> GetById(int id)
    {
        return await _context.Posts.FindAsync(id);
    }

    public async Task<Post> Create(Post post)
    {
        _context.Posts.Add(post);
        await _context.SaveChangesAsync();
        return post;
    }

    public async Task<Post?> Update(int id, UpdatePost input)
    {
        var post = await _context.Posts.FindAsync(id);
        if (post == null) return null;

        if (input.Title != null ) post.Title = input.Title;
        if (input.Content != null) post.Content = input.Content;

        await _context.SaveChangesAsync();
        return post;
    }

    public async Task<bool> Delete(int id)
    {
        var post = await _context.Posts.FindAsync(id);
        if (post == null) return false;

        _context.Posts.Remove(post);
        await _context.SaveChangesAsync();
        return true;
    }
}