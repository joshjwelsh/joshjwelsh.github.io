from django.shortcuts import render
from . models import Post
import datetime

def home(request):
    posts = Post.objects.all()
    context = {
        'posts' : posts[:5],
    }
    return render(request, 'blog/blog.html', context)

def archive(request, month, year):

    post_list = Post.objects.filter(created_on__year=year,
                                    created_on__month=month
                                    ).order_by('-created_on')
    return render(request, 'blog/blog.html', context={'posts': post_list})
                            
    