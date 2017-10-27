package main

type Meeting struct {
	Sponsor       string
	Participators string
	StartDate     string
	EndDate       string
	Title         string
}

func (m *Meeting) initMeeting(t_sponsor, t_participator, t_startDate, t_endDate, t_title string) {
	m.Sponsor = t_sponsor
	m.Participators = t_participator
	m.StartDate = t_startDate
	m.EndDate = t_endDate
	m.Title = t_title
}

func (m *Meeting) copyMeeting(t Meeting) {
	m.Sponsor = t.Sponsor
	m.Participators = t.Participators
	m.StartDate = t.StartDate
	m.EndDate = t.EndDate
	m.Title = t.Title
}

func (m Meeting) getSponsor() string {
	return m.Sponsor
}

func (m *Meeting) setSponsor(s string) {
	m.Sponsor = s
}

func (m Meeting) getParticipators() string {
	return m.Participators
}

func (m *Meeting) setParticipators(p string) {
	m.Participators = p
}
func (m Meeting) getStartDate() string {
	return m.StartDate
}

func (m *Meeting) setStartDate(s string) {
	m.StartDate = s
}
func (m Meeting) getEndDate() string {
	return m.EndDate
}

func (m *Meeting) setEndDate(e string) {
	m.EndDate = e
}
func (m Meeting) getTitle() string {
	return m.Title
}

func (m *Meeting) setTitle(t string) {
	m.Title = t
}
