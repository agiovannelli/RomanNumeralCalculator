#!/bin/bash
P=./part1/Part1_agiovannelli

r1="$($P "II plus III")"
if [ "$r1" != "V" ]; then
        echo "II plus III is not"  $r1
        exit
fi

r2="$($P "II plus III times IV minus I")"
if [ "$r2" !=  "XIII" ]; then
        echo "II plus III times IV minus I is not" $r2  
        exit
fi

r3="$($P  "II times IV minus I")"
if [ "$r3" != "VII" ]; then
        echo "II times IV minus I is not" $r3   
        exit
fi

r4="$($P  "II times {IV minus I]")"
if [ "$r4" != "VI" ]; then
        echo  "II times {IV minus I] is not" $r4        
        exit
fi

r5="$($P  "[V minus {VI minus (III minus {II minus I]}])")"
if [ "$r5" != "I" ]; then
        echo  "[V minus {VI minus (III minus {II minus I]}]) is not" $r5        
        exit
fi

r6="$($P  "III plus {IV times II] power II")"
if [ "$r6" != "LXVII" ]; then
        echo "III plus {IV times II] power II is not" $r6       
        exit
fi

r7="$($P  "{MCMXCVIII divide III divide VI minus XI) divide X")"
if [ "$r7" != "X" ]; then
        echo  "{MCMXCVIII divide III divide VI minus XI) divide X is not" $r7   
        exit
fi

r8="$($P  "{MCMXCVIII divide III divide VI minus XI) divide X power II")"
if [ "$r8" != "I" ]; then
        echo "{MCMXCVIII divide III divide VI minus XI) divide X power II is not" $r8   
        exit
fi

r9="$($P  "III plu {IV times II] power II")"
if [ "$r9" != "\
III plu {IV times II] power II
    ^
Quid dicis? You offend Caesar with your sloppy lexical habits!" ]; then
        echo  "III plu {IV times II] power II is not" $r9       
        exit
fi

r10="$($P  "I plus III minus VX times VI")"
if [ "$r10" != "\
I plus III minus VX times VI
                 ^
Quid dicis? You offend Caesar with your sloppy lexical habits!" ]; then
        echo  "I plus III minus VX times VI is not" $r10        
        exit
fi

r11="$($P  "III plus {IV times II power II")"
if [ "$r11" != "\
III plus {IV times II power II
                              ^
Quid dicis? True Romans would not understand your syntax!" ]; then
        echo "III plus {IV times II power II is not" $r11       
        exit
fi

r12="$($P  "II times (I plus II minus III)")"
if [ "$r12" != "\
II times (I plus II minus III)
                    ^
Quid dicis? Arab merchants haven't left for India yet!" ]; then
        echo "II times (I plus II minus III) is not" $r12
        exit
fi

r13="$($P  "II plus III divide IV")"
if [ "$r13" != "\
II plus III divide IV
            ^
Quid dicis? Arab merchants haven't left for India yet!" ]; then
        echo "II plus III divide IV is not" $r13        
        exit
fi

r14="$($P  "II plus I times III minus VI")"
if [ "$r14" != "\
II plus I times III minus VI
                    ^
Quid dicis? Caesar demands positive thoughts!" ]; then
        echo  "II plus I times III minus VI is not" $r14        
        exit
fi

r15="$($P  "II power III power II")"
if [ "$r15" != "DXII" ]; then
        echo "II power III power II is not" $r15        
        exit
fi

echo Congratulations, you have passed all tests in this suite!
