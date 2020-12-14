from django.urls import path
from . import views
from django.conf.urls import url

urlpatterns = [
    path('', views.home, name='blog-home'), 
    url(r'^archives/(?P<year>[0-9]{4})/(?P<month>[0-9]{1,2})/$', views.archive, name='archives'),
]