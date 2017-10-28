package main

import (
	"fmt"
)

func StartAgenda() bool {
	readCurrentUser()
	if CurrentUser.Name != "" {
		return false
	}
	readFromFile()
	return true
}

func QuitAgenda() {
	writeToFile()
	writeCurrentUser()
}

/**
* check if the username match password
* @param userName the username want to login
* @param password the password user enter
* @return if success, true will be returned
*/
func UserLogIn(userName string, password string) bool{
		filter := func(u *User) bool {
		  return u.getName() == userName && u.getPassword() == password
		}

		ulist := queryUser(filter)
	
		if (len(ulist) == 0) {
			return false
		} else {
			return true
		}
	}

/**
 * regist a user
 * @param userName new user's username
 * @param password new user's password
 * @param email new user's email
 * @param phone new user's phone
 * @return if success, true will be returned
 */	
func UserRegister(userName, password, email, phone string) bool {
	filter := func(u *User) bool {
		return u.getName()== userName
	}
	ulist := queryUser(filter)

	if (len(ulist) == 0) {
		createUser(User{userName, password, email, phone})
		return true
	} else {
		return false
	}
}

/**
 * delete a user
 * @param userName user's username
 * @param password user's password
 * @return if success, true will be returned
 */	
func DeleteUser(userName string, password string) bool{
	uf := func(u *User) bool {
		return (u.getName()== userName) && (u.getPassword() == password)
	}
	mf := func(m *Meeting) bool {
		return m.getSponsor() == userName || m.isParticipator(userName)
	}
	if (deleteUser(uf) != 0) {
		deleteMeeting(mf)
		return true
	} else {
		return false
	}
}

/**
 * list all users from storage
 * @return a user list result
 */
func ListAllUsers() []User {
	filter := func(u *User) bool {
		return true
	}
	return queryUser(filter)
}

/**
 * create a meeting
 * @param userName the sponsor's userName
 * @param title the meeting's title
 * @param participator the meeting's participator
 * @param startData the meeting's start date
 * @param endData the meeting's end date
 * @return if success, true will be returned
 */
func CreateMeeting(userName, title, startDate, endDate string, participator []string) bool {
	var sd Date
	var ed Date
	sd = stringToDate(startDate)
	ed = stringToDate(endDate)

	if ((!sd.isValid()) || (!ed.isValid())) {
		return false
	}
	if (sd.GreaterOrEqual(ed)) {
		return false
	}
	/*-------------------1-------------------------*/
	uf := func(u *User) bool {
		return userName == u.getName()
	}

	ulist := queryUser(uf)
	if (len(ulist) == 0) {
		return false
	}

	/*-------------------2-------------------------*/
	uf2 := func(u *User) bool {
		for _, p := range participator {
			if p == u.getName() {
				return true
			}
		}
		return false
	}
	ulist2 := queryUser(uf2)
	if (len(ulist2) == 0) {
		return false
	}

	/*-------------------3-------------------------*/
	uf3 := func(m *Meeting) bool {
	return title == m.getTitle()
	}
	ulist3 := queryMeeting(uf3)
	if (len(ulist3) != 0) {
		return false
	}
	/*-------------------4--------------------------*/
	uf4 := func(m *Meeting) bool {
		if (!(userName == m.getSponsor() || m.isParticipator(userName))) {
			return false
		}
		if ((userName == m.getSponsor() || m.isParticipator(userName)) &&
			(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
			ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
			return false
		} else {
			return true
		}
	}
	ulist4 := queryMeeting(uf4)
	if (len(ulist4) != 0) {
		return false
	}

	/*-------------------5--------------------------*/
	uf5 := func(m *Meeting) bool {
	for _, p := range participator {
		if ( !(p == m.getSponsor() || m.isParticipator(p)) ) {
				return false
			}
			if ((p == m.getSponsor() || m.isParticipator(p)) &&
			(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
			ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
				return false
			} else {
				return true
			}
		}
		return false
	}
	ulist5 := queryMeeting(uf5)
	if (len(ulist5) != 0) {
		return false
	}
	/*-------------------6--------------------------*/

	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if (participator[i] == participator[j]) {
				return false
			}
		}
	}

		/*-------------------７--------------------------*/
	for _, p := range participator {
		if (userName == p) {
			return false
		}
	}

	if (len(participator) == 0) {
		return false
	}


	createMeeting(Meeting{userName, startDate, endDate, title, participator})
	return true
}

/**
* search a meeting by username, time interval (user as sponsor or participator)
* @param uesrName the user's userName
* @param startDate time interval's start date
* @param endDate time interval's end date
* @return a meeting list result
*/
func MeetingQuery(name, startDate, endDate string) []Meeting {
	var ttt []Meeting
	sd := stringToDate(startDate)
	ed := stringToDate(endDate)
	if (sd.isMoreThan(ed) || !sd.isValid() || !ed.isValid()) {
		return ttt //此时a为空
	}

	filter := func(a *Meeting) bool {
		if ((a.Sponsor == name || a.isParticipator(name)) &&
		 (stringToDate(a.getEndDate()).GreaterOrEqual(sd)&&stringToDate(a.getStartDate()).SmallerOrEqual(sd))) {
			return true
		}
		if ((a.Sponsor == name || a.isParticipator(name)) && 
		(stringToDate(a.getStartDate()).SmallerOrEqual(ed)) && stringToDate(a.getStartDate()).GreaterOrEqual(sd)) {
			return true
		}
		return false
	}
	return queryMeeting(filter)
}
/**
* list all meetings the user take part in
* @param userName user's username
* @return a meeting list result
*/
func ListAllMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return a.Sponsor == name || a.isParticipator(name)
	}
	return queryMeeting(filter)
}

/**
* delete a meeting by title and its sponsor
* @param userName sponsor's username
* @param title meeting's title
* @return if success, true will be returned
*/
func ListAllSponsorMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return  name == a.Sponsor
	}
	return queryMeeting(filter)
}

/**
* list all meetings the user take part in and sponsor by other
* @param userName user's username
* @return a meeting list result
*/
func ListAllParticipateMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return a.isParticipator(name)
	}
	return queryMeeting(filter)
}

/**
* delete a meeting by title and its sponsor
* @param userName sponsor's username
* @param title meeting's title
* @return if success, true will be returned
*/
func DeleteMeeting(name, title string) bool {
	filter := func(a *Meeting) bool {
		return a.Title == title && a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}
/**
* delete all meetings by sponsor
* @param userName sponsor's username
* @return if success, true will be returned
*/
func DeleteAllMeetings(name string) bool {
	filter := func(a *Meeting) bool {
		return  a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}

func main() {
	a := "0"
	fmt.Println(a)
}