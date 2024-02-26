from django.test import TestCase

from group.models import Member


# Create your tests here.


class TestGetGroupMember(TestCase):
    def test_get_group_member(self):
        Member.objects.create(name="John", role="Developer")

        response = self.client.get("/group/John")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), {"name": "John", "role": "Developer"})

    def test_get_group_member_not_found(self):
        response = self.client.get("/group/Unknown")
        self.assertEqual(response.status_code, 404)

    def test_try_with_invalid_name(self):
        response = self.client.get("/group/1nval1d")
        self.assertEqual(response.status_code, 400)
        self.assertEqual(response.json(), {"error": "Name must be alphabetic."})
