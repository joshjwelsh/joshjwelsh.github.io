from django.db import models

 
class Post(models.Model):
    title = models.CharField(max_length=100)
    created_on = models.DateField(auto_now_add=True)
    updated_on = models.DateTimeField(auto_now= True)
    content = models.TextField()
    post_refs = models.TextField(null=True)
    class Meta:
        ordering = ['-created_on']

    def __str__(self):
        return self.title


    