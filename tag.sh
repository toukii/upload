#!/bin/sh
git add -A
read tags
git commit -m $tags
#git push origin master:master
# git push origin devm:devm
git tag -a $tags -m $tags
git push origin --tag $tags
sleep 50
git push origin --tag :$tags
git tag -d $tags