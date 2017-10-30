package entity
import (
    "testing"
    "reflect"
    "fmt"
)
var t_users = []User {
        {"a", "a", "a", "a"},
        {"b", "b", "b", "b"},
        {"c", "c", "c", "c"},
      }
var t_meetings = []Meeting {
        {"a", "a_b", "2000-02-02/00:00", "2001-02-02/00:00", []string{"b"}},
        {"b", "b_c", "2002-02-02/00:00", "2003-02-02/00:00", []string{"c"}},
        {"c", "c", "2004-02-02/00:00", "2005-02-02/00:00", []string{}},
        
}
func init() {
  userlist = nil
  meetinglist = nil
}
func TestCreateUser(t *testing.T) {
   
    cases := []struct {
        in User
        want []User
    }{
        {t_users[0], t_users[:1]},
        {t_users[1], t_users[:2]},
        {t_users[2], t_users[:3]},
    }
    for _, c := range cases {
        createUser(c.in)
        fmt.Printf("userlist: %v\n", userlist)
        if got := userlist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("CreateUser(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
func TestCreateMeeting(t *testing.T) {
    cases := []struct {
        in Meeting
        want []Meeting
    }{
        {t_meetings[0], t_meetings[:1]},
        {t_meetings[1], t_meetings[:2]},
        {t_meetings[2], t_meetings[:3]},
    }
    for _, c := range cases {
        createMeeting(c.in)
        fmt.Printf("meetinglist: %v\n", meetinglist)
        if got := meetinglist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("CreateMeeting(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
func TestQueryMeeting(t *testing.T) {
    cases := []struct {
        in mFilter
        want []Meeting
    }{
        {func (m *Meeting) bool { return m.getSponsor() == "a"}, []Meeting{t_meetings[0]}},
        {func (m *Meeting) bool { return m.getSponsor() == "b"}, []Meeting{t_meetings[1]}},
        {func (m *Meeting) bool { return m.getSponsor() == "c"}, []Meeting{t_meetings[2]}},
    }
    for _, c := range cases {
        got := queryMeeting(c.in)
        if !reflect.DeepEqual(got, c.want)  {
            t.Errorf("queryMeeting(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}


