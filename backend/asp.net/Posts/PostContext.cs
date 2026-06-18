using Microsoft.EntityFrameworkCore;

namespace BlogAPI.Models;

public class PostContext : DbContext
{
    public PostContext(DbContextOptions<PostContext> options)
        : base(options)
    {    
    }

    public DbSet<Post> Posts { get; set; } = null!;

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Post>(entity =>
        {
            entity.ToTable("posts");
            entity.Property(e => e.Id).HasColumnName("id").ValueGeneratedOnAdd();
            entity.Property(e => e.Title).HasColumnName("title");
            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.CreatedAt).HasColumnName("created_at").HasDefaultValueSql("NOW()");
            entity.Property(e => e.UpdatedAt).HasColumnName("updated_at").HasDefaultValueSql("NOW()");
        });    
    }
}