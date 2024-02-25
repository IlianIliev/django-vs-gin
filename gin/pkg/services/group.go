package services

import (
	"errors"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/entities"
	"sort"
)

var ErrMemberNotFound = errors.New("member not found")

type Group interface {
	ListMembers() ([]entities.Member, error)
	GetMember(name string) (entities.Member, error)
	AddMember(member entities.Member) error
}

type inMemoryGroupService struct {
	members map[string]entities.Member
}

func (s *inMemoryGroupService) ListMembers() ([]entities.Member, error) {
	var members []entities.Member
	for _, member := range s.members {
		members = append(members, member)
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].Name < members[j].Name
	})

	return members, nil
}

func (s *inMemoryGroupService) GetMember(name string) (entities.Member, error) {
	member, ok := s.members[name]
	if !ok {
		return entities.Member{}, ErrMemberNotFound
	}

	return member, nil
}

func (s *inMemoryGroupService) AddMember(member entities.Member) error {
	s.members[member.Name] = member
	return nil
}

func NewInMemoryGroupService() Group {
	members := map[string]entities.Member{
		"John": {Name: "John", Role: "Leader"},
		"Jane": {Name: "Jane", Role: "Member"},
		"Jack": {Name: "Jack", Role: "Member"},
	}

	return &inMemoryGroupService{members: members}
}
