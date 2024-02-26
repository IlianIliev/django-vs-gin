from django.urls import path

from .views import list_group_members, get_group_member


urlpatterns = [
    path("", list_group_members, name="list_group_members"),
    path("<str:name>", get_group_member, name="get_group_member"),
]
