package entity

import (
	"fmt"
)


//不要在login后调用StartAgenda
func StartAgenda() bool {
	ReadFromFile()
	ReadCurrentUser()
	if CurrentUser.Name == "" {
		return false
	}
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
//登录命令不需要调用StartAgenda,但需要调用QuitAgenda来保存登录信息
func UserLogIn(userName string, password string) bool{
		ReadFromFile()
		if (CurrentUser.Name != "") {
			return false
		}
		filter := func(u *User) bool {
		  return u.getName() == userName && u.getPassword() == password
		}

		ulist := queryUser(filter)
		if (len(ulist) == 0) {
			return false
		} else {
			//当前用户信息
			CurrentUser = ulist[0]
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
func DeleteUser(userName string, password string) bool {
	uf := func(u *User) bool {
		return (u.getName()== userName) && (u.getPassword() == password)
	}
	mf := func(m *Meeting) bool {
		return m.getSponsor() == userName || m.isParticipator(userName)
	}
	if (deleteUser(uf) != 0) {
		deleteMeeting(mf)
		if (userName == CurrentUser.Name) {
			CurrentUser.InitUser("", "", "", "")
		}
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
		fmt.Println("日期不合法")
		return false
	}
	if (sd.GreaterOrEqual(ed)) {
		fmt.Println("开始日期不可大于结束日期")
		return false
	}
	
	/*-------------------1-------------------------*/
	uf := func(u *User) bool {
		return userName == u.getName()
	}

	ulist := queryUser(uf)
	if (len(ulist) == 0) {
		fmt.Println("发起人未注册")
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
	if (len(ulist2) != len(participator)) {
		fmt.Println("存在参与者未注册")
		return false
	}
	
	/*-------------------3-------------------------*/
	uf3 := func(m *Meeting) bool {
	return title == m.getTitle()
	}
	ulist3 := queryMeeting(uf3)
	if (len(ulist3) != 0) {
		fmt.Println("会议主题已存在")
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
		fmt.Println("与发起人其他会议冲突")
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
		fmt.Println("与发起人其他会议冲突")
		return false
	}
	
	/*-------------------6--------------------------*/

	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if (participator[i] == participator[j]) {
				fmt.Println("参与者不能重复的人")
				return false
			}
		}
	}
	
	/*-------------------７--------------------------*/
	for _, p := range participator {
		if (userName == p) {
			fmt.Println("参与者不能有发起者")
			return false
		}
	}

	/*-------------------8--------------------------*/
	if (len(participator) == 0) {
		fmt.Println("参与者不能为空")
		return false
	}

	createMeeting(Meeting{userName, title, startDate, endDate, participator})
	return true
}

/**
* search a meeting by username, time interval (user as sponsor or participator)
* @param uesrName the user's userName
* @param startDate time interval's start date
* @param endDate time interval's end date
* @return a meeting list result
*/
func MeetingQuery(sponsor, startDate, endDate string) []Meeting {
	var ttt []Meeting
	sd := stringToDate(startDate)
	ed := stringToDate(endDate)
	if (sd.isMoreThan(ed) || !sd.isValid() || !ed.isValid()) {
		fmt.Println("日期不合法")
		return ttt //此时a为空
	}

	filter := func(a *Meeting) bool {
		if ((a.Sponsor == sponsor || a.isParticipator(sponsor)) &&
		 (stringToDate(a.getEndDate()).GreaterOrEqual(sd)&&stringToDate(a.getStartDate()).SmallerOrEqual(sd))) {
			return true
		}
		if ((a.Sponsor == sponsor || a.isParticipator(sponsor)) && 
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

