from django.shortcuts import render
from django.http import HttpResponse


def index(request):
    context = {'data': list()}
    return render(request, 'core/home.html', context)

def code(request):
    return render(request, 'core/code.html')

def consulting(request):
    return render(request, 'core/consulting.html')