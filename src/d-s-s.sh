#!/bin/bash

cleanup(){
	echo "hello: cleanup is called on exit"
	rm -f /tmp/lock1
	rm -f /tmp/lock2
}

main(){
if [ $(id -u) != 0 ];then
	echo "Run 'd-s-s' as root user"
	exit 
fi
lockfile1=/tmp/lock1
lockfile2=/tmp/lock2
echo "hello d-s-s"
mkdir /tmp/foo
mount -o bind /tmp/foo /tmp/foo
set -o noclobber
echo $$ > $lockfile1
echo $$ > $lockfile2
unshare -m /home/smahajan/Downloads/defDeletion.sh $lockfile1 $lockfile2 &
echo "AFTER UNSHARE"
while true;
do 
	if (echo "$$" > "$lockfile1") 2> /dev/null; then
		echo "Unmounting and removing foo"
		umount /tmp/foo
		rmdir /tmp/foo
		break
	else	
		echo "lock exists"
		continue
	fi
done
rm $lockfile2
trap cleanup EXIT
exit 0
}

main "$@"

