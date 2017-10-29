
agenda命令设计

-------------------

```
 ./agenda
you can use this app to create or remove meetings.Also you must register a user to have the rights to use the functions.

Usage:
  agenda [command]

Available Commands:
  add         To add Participator of the meeting
  clear       clear all the meeting
  clear       To clear all the meeting created by the current user
  create      To create a new meeting
  delete      To delete your account in Agenda
  deleteM     delete meeting with the title [title]
  help        Help about any command
  login       Using UserName with PassWord to login Agenda.
  logout      To logout Agenda
  queryM      To query all the meeting have attended during [StartTime] and [EndTime]
  queryU      To query all the users' names
  quit        quit the meeting with the title [title]
  register    Register a new user
  remove      To remove Participator from the meeting

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.
```
各个命令的用法和功能：

 1. register（注册用户）
 - agenda register -u [UserName] -p [Pass] -e [Email] -t [Phone]

```
./agenda register -h
Input command register -u UserName -p PassWord -e Email -t Phone:

[Username] is the name of the new register
[PassWord] is the password to login
[Email]is the email address of the register
[Phone] is the phone of the register

Usage:
  agenda register -u [UserName] -p [Pass] -e [Email] -t [Phone] [flags]

Flags:
  -h, --help   help for register

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

2. login（用户登录）

 - agenda login -u [UserName] -p [PassWord]
```
./agenda login -h
Using UserName and PassWord to login Agenda:

attention:If the PassWord is right,you can login Agenda and use it
If forget the PassWord,you must register another one User

Usage:
  agenda login -u [UserName] -p [PassWord] [flags]

Flags:
  -h, --help   help for login

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

3. logout（当前用户登出）

 - agenda logout 
```
 ./agenda logout -h
After logouting,you can only register or login:

register -u [UserName] -p [Pass] -e [Email]
login -u [UserName] -p [PassWord]

Usage:
  agenda logout [flags]

Flags:
  -h, --help   help for logout

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

4. queryU（查看已注册的所有用户的用户名、邮箱及电话信息）

 - agenda queryU
```
 ./agenda queryU -h
You can query all the users's names who have registed.

Usage:
  agenda queryU [flags]

Flags:
  -h, --help   help for queryU

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

5. delete（删除本用户）

 - agenda delete 
```
./agenda delete -h
you can delete your account in the database of Agenda:

attention:After deleting,you will need to register a new User to login Agenda.

Usage:
  agenda delete [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

6. create（创建会议）

 - agenda create -t [Title] -p [Participator] -s [StartTime] -e [EndTime] 
```
 ./agenda create -h
To create a new meeting with:

[Title] the Title of the meeting
[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
[StartTime] the StartTime of the meeting
[EndTime] the EndTime of the meeting

Usage:
  agenda create -t [Title] -p [Participator] -s [StartTime] -e [EndTime] [flags]

Flags:
  -h, --help   help for create

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

7. add（添加会议参与者）

 - agenda add -p [Participator] -t [Title]
```
/agenda$ ./agenda add -h
Add [Participator] to the meeting with the title of [Title]:

attention:If the Participator cannot attend during the time, add fail.

Usage:
  agenda add -p [Participator] -t [Title] [flags]

Flags:
  -h, --help   help for add

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

8. remove（从某会议删除某参与者）

 - agenda remove -p [Participator] -t [Title]
```
 ./agenda remove -h
remove [Participator] from the meeting with the title of [Title]:

attention:If there is no Participators in the meeting,the meeting will be deleted

Usage:
  agenda remove -p [Participator] -t [Title] [flags]

Flags:
  -h, --help   help for remove

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

9. queryM（查询会议）

 - agenda queryM -s [StartTime] -e [EndTime] 
```
./agenda queryM -h
You can query all the meeting have attended during [StartTime] and [EndTime]

Usage:
  agenda queryM -s [StartTime] -e [EndTime] [flags]

Flags:
  -h, --help   help for queryM

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

10. deleteM（根据会议名删除会议）

 - agenda delete -t [title]
```
./agenda deleteM -h
you can delete one meeting with the title [title]

Usage:
  agenda deleteM -t [title] [flags]

Flags:
  -h, --help   help for deleteM

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

11. quit（当前用户退出会议）

 - agenda quit -t [title]
```
./agenda quit -h
you can quit the meeting with the title of [title]:

attention:if there is no participators in this meeting,the meeting will be deleted

Usage:
  agenda quit -t [title] [flags]

Flags:
  -h, --help   help for quit

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

12. clear（清除当前用户创建的会议）

 - agenda clear
```
./agenda clear
You can delete all the meeting created by you

Usage:
  agenda clear [flags]

Flags:
  -h, --help   help for clear

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)


```


