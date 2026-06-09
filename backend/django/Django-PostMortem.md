# Django Framework Post Mortem

It's been awhile since I wrote any Python, the last time it was for data analysis shenanigans. Python really feels different from PHP where classes and functions are defined by indentation rather than using curly brackets `{}`. Luckily, this was not enough to stop me from making an API application with Django. However, I failed realise one thing when working on this project: 

#### Django is a full-stack web framework, not a backend framework.

Which means, it has no natural way of handling APIs. The solution (or to streamline the process of building APIs), is by implementing Django REST Framework(DRF). Imagine, a framework on top of a framework. This can be implemented by including it in the `requirements.txt` and `core/settings.py` (then rebuild the container if you're using Docker).

Once this was sorted out, I just did my thing. However, instead of taking the Object-Oriented Programming approach which is the intended approach when using DRF, I tried the Data-Oriented Programming approach. Which means making structs using @dataclass and no access to inheritance, so `APIController` for shared API logic becomes a middleware. I works as I expected it to be but in an event where this project scales (unlikely, this is just a toy project), I would be fighting against DRF conventions and not taking advantages of this feature. Therefore, in the future:

#### When making an API application using Django REST Framework, please take the OOP approach.

Other than that, I can confidently say that I didn't encounter any difficulty trying out Django, except that it won't run if there are errors in the application itself, unlike Laravel where if you import a nonexistent controller, it would just falls silently, Django will immediately throw an error. I can say this type of behavior is a matter of taste, one person would like this feature while another will find it annoying or agitating, I'm neutral on this of course.

Lastly, I would like to touch on programming in Python itself. After talking to a good friend I met online who happens to be expert in Python backend development, Jerry, he suggested using `uv` instead of `pip` as the package manager, `ruff` for linting and `ty` for type checking. After reading `uv` documentations, they claim `uv` is faster, way way faster than `pip`. Definitely will try it out when working on FastAPI. Having a proper linter would also help, especially since I quite pedantic of small inconsistencies in my codebase. Finally, type checking with `ty`. Oh you have no idea how I love having things like this around. I strongly hate dealing with bugs spawning from type issues so being able to statically type Python really makes me a happy programmer.