# 服务计算作业3——agenda(测试)

测试agenda项目


-------------------
##**agenda界面**：

```
 ./agenda

```

界面如下：


```
you can use this app to create or remove meetings.Also you must register a user to have the rights to use the functions.

Usage:
  agenda [command]

Available Commands:
  add         To add Participator of the meeting
  clear       clear all the meeting created by the current user
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

##**1.测试register**：
###创建三个用户：Alice、Bob、Cobra

```
./agenda register -u [UserName] -p [Pass] -e [Email] -t [Phone] 
```

```
./agenda register -u Alice -p Alice -e Alice@163.com -t 123456
Register new user successfully
```
```
./agenda register -u Bob -p Bob -e Bob@163.com -t 123456
Register new user successfully
```
```
./agenda register -u Cobra -p Cobra -e Cobra@163.com -t 123456
Register new user successfully
```
###注册成功:
```
Register new user successfully
```

##**2.测试login**：
###登录Alice
用法：

```
./agenda login -u [UserName] -p [PassWord]
```

```
./agenda login -u Alice -p Alice
```
登录成功：

```
Log in successfully
```
##**3.测试queryU**：
###用Alice的账号查询其他两位用户的信息
用法：

```
./agenda queryU
```
查询结果：

```
1. Alice Alice@163.com 123456
2. Bob Bob@163.com 123456
3. Cobra Cobra@163.com 123456
```
（只能查到名字、邮箱、电话，而不能查询密码）

##**4.测试create**：
###用Alice的账号创建两个会议：
####Alice_Bob:    2000-01-01/00:00  2001-01-01/00:00
####Alice_Cobra:2002-01-01/00:00  2003-01-01/00:00
用法：
```
 ./agenda create -t Alice_Bob -p Bob -s 2000-01-01/00:00 -e 2001-01-01/00:00
Create meeting successfully
 ./agenda create -t Alice_Cobra -p Cobra -s 2002-01-01/00:00 -e 2003-01-01/00:00
Create meeting successfully

```
创建成功。

##**5.测试queryM**:
###用Alice的账号查询她参加的所有会议
用法：
```
 ./agenda queryM -s [StartTime] -e [EndTime]
```
```
./agenda queryM -s 2000-01-01/00:00 -e 2003-01-01/00:00
```

查询结果：
```
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-01-01/00:00 2001-01-01/00:00 [Bob]}
2. {Alice Alice_Cobra 2002-01-01/00:00 2003-01-01/00:00 [Cobra]}

```
查询成功。
创建失败的例子：（Bob不能分身参加两个会议）

```
 ./agenda create -t Bob_Cobra -p Bob -s 2000-01-01/00:00 -e 2002-01-01/00:00
```
```
与发起人或者参与者其他会议冲突
Fail to create meeting
```
##**6.测试deleteM**:
###用Alice的账号删除Alice_Cobra会议
用法：

```
agenda deleteM -t [title]
```

```
 ./agenda deleteM -t Alice_Cobra
 Delete meeting successfully
```
查询会议结果如下：（只剩下一个会议Alice_Bob）

```
queryM -s 2000-01-01/00:00 -e 2003-01-01/00:00
```
```
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-01-01/00:00 2001-01-01/00:00 [Bob]}
```
###删除会议成功。


##**7.测试delete**：
###登出Alice，登入Cobra,测试delete:
```
./agenda logout
Log out successfully
```
```
./agenda login -u Cobra -p Cobra
Log in successfully
```
```
./agenda delete
Delete this account successfully
```
###再次登入Alice查询用户结果：
```
./agenda login -u Alice -p Alice
Log in successfully
```
```
./agenda queryU
Name  Email Telephone
1. Alice Alice@163.com 123456
2. Bob Bob@163.com 123456
```
###Cobra注销成功。

##**8.测试add**：
###重新注册Cobra，登入Alice，添加Cobra为会议Alice_Bob参与者：

```
 ./agenda add -p Cobra -t Alice_Bob
Add participators successfully
```
```
./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-02-02/00:00 2001-02-02/00:00 [Bob Cobra]}

```
###Cobra加入成功。

##**9.测试quit**：
###登入Bob，使用quit推出Alice_Bob会议：
用法：

```
./agenda quit -t [title]
```

```
./agenda login -u Bob -p Bob
Log in successfully

 ./agenda quit -t Alice_Bob
Quit meeting successfully

```
###登入Alice查看会议的参加人员：

```
./agenda logout
Log out successfully

 ./agenda login -u Alice -p Alice
Log in successfully

./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-02-02/00:00 2001-02-02/00:00 [Cobra]}

```
###Bob退出会议成功。

##**10.测试remove**：
###再次把Bob加入Alice_Bob会议中:

```
./agenda add -p Bob -t Alice_Bob
Add participators successfully

./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-02-02/00:00 2001-02-02/00:00 [Cobra Bob]}
```
remove 用法:
```
 ./agenda remove -p [Participator] -t [Title]
```
###移出Cobra：
```
 ./agenda remove -p Cobra -t Alice_Bob
Remove participators successfully
```
```
./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-02-02/00:00 2001-02-02/00:00 [Bob]}

```
###Cobra成功退出Alice_Bob会议。

##**11.测试clear**：
###登入Bob建立会议Bob_Alice会议：

```
 ./agenda login -u Bob -p Bob
Log in successfully
 ./agenda create -t Bob_Alice -p Alice -s 2002-02-02/00:00 -e 2003-02-02/00:00
```
###登录Alice查看会议：
```
./agenda logout
Log out successfully

./agenda login -u Alice -p Alice
Log in successfully

./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Alice Alice_Bob 2000-02-02/00:00 2001-02-02/00:00 [Bob]}
2. {Bob Bob_Alice 2002-02-02/00:00 2003-02-02/00:00 [Alice]}
```
###使用clear：
```
./agenda clear
Clear meeting successfully
r ./agenda queryM -s 2000-01-01/00:00 -e 2005-01-01/00:00
Query meeting successfully

Sponsor Title StartDate EndDate Participators
1. {Bob Bob_Alice 2002-02-02/00:00 2003-02-02/00:00 [Alice]}
```
###清空发起会议成功。