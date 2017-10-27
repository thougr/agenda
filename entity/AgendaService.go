package main


func UserLogIn(userName string, password string) bool{

    var filter uFilter
    filter = func(u *User) bool {
      return u.getName() == userName && u.getPassword() == password
    }
    ulist := queryUser(filter)

    if (len(ulist) == 0) {
        return false
    } else {
        return true
    }
}

func UserRegister(userName string, password string,
   email string, phone string) bool {
    var filter uFilter
    filter = func(u *User) bool {
      return u.getName()== userName
    }
    ulist := queryUser(filter)

    if (len(ulist) == 0) {
        user := User {
          Name     : userName,
        	Password : password,
        	Email    : email,
        	Phone    : phone,
        }
        createUser(user)
        return false
    } else {
        return true
    }
}

func DeleteUser(userName string, password string) bool{
    var uf uFilter
    var mf mFilter
    uf = func(u *User) bool {
      return (u.getName()== userName) && (u.getPassword() == password)
    }
    mf = func(m *Meeting) bool {
      return m.getSponsor() == userName || m.isParticipator(userName)
    }
    if (deleteUser(uf) != 0) {
        deleteMeeting(mf)
        return true
    } else {
        return false
    }
}
func ListAllUsers() []User {
    var filter uFilter
    filter = func(u *User) bool {
      return true
    }
    return queryUser(filter)
}

func CreateMeeting(userName string, title string,
                     startDate string, endDate string, participator []string) bool {
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
       var uf uFilter = func(u *User) bool {
         return userName == a.getName()
       }

       var ulist []User = queryUser(uf)
       if (len(ulist) == 0) {
           return false
       }

      /*-------------------2-------------------------*/
       var uf2 uFilter = func(u *User) bool {
         for _, p := range participator {
           if p == u.getName() {
             return true
           }
         }
         return false
       }
       var ulist2 []User = queryUser(uf2)
       if (len(ulist2) == 0) {
           return false
       }

      /*-------------------3-------------------------*/
      var uf3 mFilter = func(m *Meeting) bool {
        return title == m.getTitle()
      }
     var ulist3 []Meeting = queryMeeting(uf3)
     if (len(ulist3) != 0) {
         return false
     }
     /*-------------------4--------------------------*/
     var uf4 mFilter = func(m *Meeting) bool {
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

     /*-------------------5--------------------------*/
     var uf5 mFilter = func(m *Meeting) bool {
       for _, p := range participator {
         if ( !(p == m.getSponsor() || m.isParticipator(p)) ) {
                  return false
              }
              if ((p == a.getSponsor() || m.isParticipator(p)) &&
              (sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
              ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
                  return false
              } else {
                  return true
              }
         }
         return false
     }
     var ulist5 []Meeting = queryMeeting(uf5)
     if (len(ulist5) != 0) {
         return false
     }
     /*-------------------6--------------------------*/

     for i, p := range participator {
       for _, pp := range participator[i+1:] {
         if (p == pp) {
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
    var mt Meeting = Meeting {
      Sponsor       :userName,
      Participators :participator,
      StartDate     :startDate,
      EndDate       :endDate,
      Title         :title,
    }
    createMeeting(mt)
    return true
}


func main() {
     UserLogIn("5525", "5545")
}
