#!/bin/bash

SUCCESS=0
FAIL=0
DIFF_RES=""
RES=""
EXE="./../../myWc"

go build compareFS/main.go

for flag in -m -l -w
do
    for var in docs/1.txt docs/2.txt docs/3.txt docs/4.txt docs/5.txt docs/6.txt docs/7.txt docs/8.txt
    do
        TEST1="$flag $var"
        $EXE $TEST1 | sed 's/[[:blank:]]*//g' > myWc.txt
        wc $TEST1 | grep -v total | sed 's/[[:blank:]]*//g' > wc.txt
        DIFF_RES="$(./main -old myWc.txt -new wc.txt)"
        if [ "$DIFF_RES" != "$RES" ]
            then
                echo "$TEST1"
                echo "$DIFF_RES"
                cat myWc.txt wc.txt
                echo "\n"
                FAIL=$((FAIL+1))
            else
                SUCCESS=$((SUCCESS+1))
        fi
        rm myWc.txt wc.txt
    done
done

# for flag in -m -l -w
# do
#     for var1 in docs/1.txt docs/2.txt docs/3.txt docs/4.txt docs/5.txt docs/6.txt docs/7.txt docs/8.txt
#     do
#         for var2 in docs/1.txt docs/2.txt
#         do
#             for var3 in docs/2.txt docs/3.txt
#             do
#                 for var4 in docs/3.txt docs/4.txt
#                 do
#                     for var5 in docs/4.txt docs/5.txt
#                     do
#                         for var6 in docs/5.txt docs/6.txt
#                         do
#                             for var7 in docs/6.txt docs/7.txt
#                             do
#                                 for var8 in docs/7.txt docs/8.txt
#                                 do
#                                     TEST1="$flag $var1 $var2 $var3 $var4 $var5 $var6 $var7 $var8"
#                                     $EXE $TEST1 | sed 's/[[:blank:]]*//g' > myWc.txt
#                                     wc $TEST1 | grep -v total | sed 's/[[:blank:]]*//g' > wc.txt
#                                     DIFF_RES="$(./main -old myWc.txt -new wc.txt)"
#                                     if [ "$DIFF_RES" != "$RES" ]
#                                         then
#                                             echo "$TEST1"
#                                             echo "$DIFF_RES"
#                                             cat myWc.txt wc.txt
#                                             echo "\n"
#                                             FAIL=$((FAIL+1))
#                                         else
#                                             SUCCESS=$((SUCCESS+1))
#                                     fi
#                                     rm myWc.txt wc.txt
#                                 done
#                             done
#                         done
#                     done
#                 done
#             done
#         done
#     done
# done

rm main

echo "SUCCESS: $SUCCESS"
echo "FAIL: $FAIL"
