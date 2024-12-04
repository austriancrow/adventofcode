
cp -r day_X day_$1
mv day_$1/dayX_test.go day_$1/day$1_test.go
sed -i 's/.\/day_X/.\/day_'$1'\n\t.\/day_X/' go.work
sed -i 's/dayX/day'$1'/' day_$1/go.mod
sed -i 's/dayX/day'$1'/g' day_$1/main.go
sed -i 's/Day X/Day '$1'/g' day_$1/main.go
sed -i 's/dayX/day'$1'/g' day_$1/funcs/part1.go
sed -i 's/dayX/day'$1'/g' day_$1/funcs/part2.go
sed -i 's/dayX/day'$1'/g' day_$1/day$1_test.go
git add .
git commit -m "feat: starting day "$1
cd day_$1