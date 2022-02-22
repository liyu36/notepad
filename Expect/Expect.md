# Expect 笔记

## 常用关键字
``` bash
$argv0           # 脚本名字 等价于$0
[lindex $argv 0] # 位置变量 等价于$1  
$argc            # 变量个数 等价于$@
puts             # 输出内容 等价于echo
exit             # 退出 exit 1 
spawn            # 启动交互进程
expect           # 接收字符串
send             # 发送字符串
interact         # 等待用户交互
exp_continue     # 让expect程序继续匹配
```

## 流程控制
``` bash
if {} {
    set path "/mnt"
} elseif {$flag == "c"} {
    set path "/tmp"
} else { 
    puts "Usage: cmd <host> <m|c>"
    exit 2
}
```

## 使用例子

### 1. 自动登录
``` bash
#!/usr/bin/expect

if {$argc < 2} {
    puts "Usage: $argv0 <host> <flag>"
    exit 1
}

set timeout 5
set hostname [lindex $argv 0]
set password "toor"
set flag [lindex $argv 1]

if {$flag == "m"} {
    set path "/mnt"
} elseif {$flag == "t"} {
    set path "/tmp" 
} else {
    puts "Usage: $argv0 $hostname <flag>"
    exit 2
}

spawn ssh $hostname 
expect {
    "*yes/no*" { send "yes\r"; exp_continue }
    "*assword*" { send "$password\r" }
} 
expect "*~$ " { send "sudo su - \r" }
expect "\[sudo]*" { send "$password\r" } 
expect "*~# " { send "cd $path;pwd\r" }

interact
```

### 2.shell脚本自动登录执行命令
``` bash
#!/usr/bin/env bash
__Author__="liy"

: '
parameter1: username
parameter2: hostname
parameter3: password
output: None
return: None
'
function autorun(){
    local username="$1"
    local hostname="$2"
    local password="$3"

    /usr/bin/expect <<-EOF
    set timeout 5
    spawn ssh $username@$hostname
    expect {
        "*yes/no*" { send "yes\r"; exp_continue }
        "*password*" { send "$password\r" }
    }
    expect "*~*" { send "echo hello\r" }
    expect "*~*" { send "exit\r" }
    expect eof
EOF
}
```
