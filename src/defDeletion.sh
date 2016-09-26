#!/bin/bash

echo "hello deferred deletion"
echo $$
cat $1
rm $1
while true;
do
	if (set -o noclobber; echo "$$" > "$lockfile2") 2> /dev/null; then
		break
	else
		continue
	fi
done
exit 0

