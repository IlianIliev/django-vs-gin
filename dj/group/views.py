from django.http import JsonResponse, Http404

from group.models import Member


def member_to_json(member):
    return {
        "name": member.name,
        "role": member.role,
    }


def list_group_members(request):
    members = Member.objects.all()

    return JsonResponse({"members": [member_to_json(member) for member in members]})


def get_group_member(request, name):
    if not name.isalpha():
        return JsonResponse({"error": "Name must be alphabetic."}, status=400)

    try:
        member = Member.objects.get(name=name)
    except Member.DoesNotExist:
        raise Http404()

    return JsonResponse(member_to_json(member))
