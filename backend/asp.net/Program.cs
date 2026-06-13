var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

// Route registration
app.MapGet("/", () => "Hello World!");

app.Run();
