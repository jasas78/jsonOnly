#!/bin/bash

sleep 1
test      -f /proc/sys/net/ipv6/conf/all/disable_ipv6 && \
    echo 1 > /proc/sys/net/ipv6/conf/all/disable_ipv6

sleep 1
nohup  /root/jsonOnly/objSaved/goHtml_server.lnx.LinuxX64.80.exe > /dev/null &

sleep 1


## check : /etc/crontab
#nohup bash /delete_unknown.sh &
#sleep 1

for aa1 in \
/bin/note     \
/etc/my.conf    \
/usr/bin/pythno    \
/usr/bin/.sshd    \
/var/ssh.conf    \
/lib/systemd/system/apt-daily.service    \
/lib/systemd/system/painintheapt-daily.service    \
/lib/systemd/system/apt-daily-upgrade.service    \
/lib/systemd/system/apt-daily-upgrade.timer    \
/usr/bin/dpkgd \
/usr/bin/bsd-port \

do
    if [ -f ${aa1} ] ; then
        rm -f ${aa1}.bak01
        mv ${aa1} ${aa1}.bak01
        echo > ${aa1}
    fi
done


echo ' the following is the crackjack files '
echo 'find /bin/ -size 1135000c'
echo 'find /usr/bin/ -size 1135000c'
echo


