#!/bin/bash

#first test
sort -nk 2 testfiles/flagK.txt > results/flagKN_sort.txt
go run ../cmd/main.go -nk 2 testfiles/flagK.txt > results/flagKN_app.txt

DIFF=$(diff results/flagKN_sort.txt results/flagKN_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -nk PASSED"
else
    echo "TEST -nk FAIL"
fi

#second test
sort -nr testfiles/flagN.txt > results/flagNR_sort.txt
go run ../cmd/main.go -nr testfiles/flagN.txt > results/flagNR_app.txt

DIFF=$(diff results/flagNR_sort.txt results/flagNR_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -nr PASSED"
else
    echo "TEST -nr FAIL"
fi

#second test
sort -uM testfiles/flagM.txt > results/flaguM_sort.txt
go run ../cmd/main.go -uM testfiles/flagM.txt > results/flaguM_app.txt

DIFF=$(diff results/flaguM_sort.txt results/flaguM_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -uM PASSED"
else
    echo "TEST -uM FAIL"
fi