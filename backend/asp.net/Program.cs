using Microsoft.EntityFrameworkCore;
using BlogAPI.Services;
using BlogAPI.Models;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

var host = Environment.GetEnvironmentVariable("DB_HOST");
var port = Environment.GetEnvironmentVariable("DB_PORT");
var name = Environment.GetEnvironmentVariable("DB_NAME");
var user = Environment.GetEnvironmentVariable("DB_USER");
var pass = Environment.GetEnvironmentVariable("DB_PASSWORD");
var connectionString = $"Host={host};Port={port};Database={name};Username={user};Password={pass}";

builder.Services.AddDbContext<PostContext>(opt =>
    opt.UseNpgsql(connectionString));
builder.Services.AddScoped<PostService>();

var app = builder.Build();

app.MapControllers();

app.Run();
