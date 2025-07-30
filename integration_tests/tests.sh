#!/bin/bash

#first test
sort -k 2 testfiles/flagK.txt > results/flagK_sort.txt
go run ../cmd/main.go -k 2 testfiles/flagK.txt > results/flagK_app.txt

DIFF=$(diff results/flagK_sort.txt results/flagK_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -k PASSED"
else
    echo "TEST -k FAIL"
fi

#second test
sort -n testfiles/flagN.txt > results/flagN_sort.txt
go run ../cmd/main.go -n testfiles/flagN.txt > results/flagN_app.txt

DIFF=$(diff results/flagN_sort.txt results/flagN_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -n PASSED"
else
    echo "TEST -n FAIL"
fi

#third test
sort -r testfiles/flagK.txt > results/flagR_sort.txt
go run ../cmd/main.go -r testfiles/flagK.txt > results/flagR_app.txt

DIFF=$(diff results/flagR_sort.txt results/flagR_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -r PASSED"
else
    echo "TEST -r FAIL"
fi

#fourth test
sort -u testfiles/flagU.txt > results/flagU_sort.txt
go run ../cmd/main.go -u testfiles/flagU.txt > results/flagU_app.txt

DIFF=$(diff results/flagU_sort.txt results/flagU_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -u PASSED"
else
    echo "TEST -u FAIL"
fi


#fifth test
sort -M testfiles/flagM.txt > results/flagM_sort.txt
go run ../cmd/main.go -M testfiles/flagM.txt > results/flagM_app.txt

DIFF=$(diff results/flagM_sort.txt results/flagM_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -M PASSED"
else
    echo "TEST -M FAIL"
fi

#sixth test
sort -b testfiles/flagB.txt > results/flagB_sort.txt
go run ../cmd/main.go -b testfiles/flagB.txt > results/flagB_app.txt

DIFF=$(diff results/flagB_sort.txt results/flagB_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -b PASSED"
else
    echo "TEST -b FAIL"
fi

#seventh test
sort -c testfiles/flagC.txt 2> results/flagC_sort.txt
go run ../cmd/main.go -c testfiles/flagC.txt > results/flagC_app.txt

DIFF=$(diff results/flagC_sort.txt results/flagC_app.txt) 
if [ "$DIFF" = "" ] 
then
    echo "TEST -c PASSED"
else
    echo "TEST -c FAIL"
fi