#!/usr/bin/env bash

cp ./testdata/triceCheck.c.txt ./testdata/triceCheck.c
cp ./testdata/triceCheck.c.txt ../examples/exampleData/triceCheck.c

# create til.json && li.json if not existing
touch ./testdata/til.json ./testdata/li.json

#  insert IDs into source code
trice insert -i ./testdata/til.json -li ./testdata/li.json -liPathIsRelative -IDMin 1000 -IDMax 7999

# The file cgoPackage.go is the same in all cgo test packages, but must be inside the folders.
# os agnostic links would be better.
CGOTESTDIRS="
_ringB_di_xtea_cobs_rtt32__de_tcobs/
_stackB_di_xtea_cobs_rtt8/
_staticB_di_nopf_aux32/
_staticB_di_xtea_cobs_rtt32/
100_dblB_de_multi_nopf/
100_dblB_de_nopf/
100_ringB_de_nopf/
100_ringB_di_xtea_cobs_rtt32__de_xtea_cobs/
dblB_de_cobs/
dblB_de_multi_cobs/
dblB_de_multi_tcobs/
dblB_de_multi_xtea_cobs/
dblB_de_multi_xtea_tcobs/
dblB_de_tcobs/
dblB_de_xtea_cobs/
dblB_de_xtea_tcobs/
dblB_di_nopf_rtt32__de_cobs/
dblB_di_nopf_rtt32__de_multi_cobs/
dblB_di_nopf_rtt32__de_multi_tcobs/
dblB_di_nopf_rtt32__de_tcobs/
dblB_di_nopf_rtt8__de_cobs/
dblB_di_nopf_rtt8__de_multi_cobs/
dblB_di_nopf_rtt8__de_multi_tcobs/
dblB_di_nopf_rtt8__de_tcobs/
ringB_de_cobs/
ringB_de_tcobs/
ringB_de_xtea_cobs/
ringB_de_xtea_tcobs/
ringB_di_cobs_rtt32__de_tcobs/
ringB_di_cobs_rtt8__de_tcobs/
ringB_di_nopf_aux8/
ringB_di_nopf_rtt32__de_tcobs/
ringB_di_nopf_rtt32__de_xtea_cobs/
ringB_di_nopf_rtt8__de_tcobs/
ringB_di_tcobs_rtt32__de_tcobs/
stackB_di_nopf_aux8/
stackB_di_nopf_rtt32/
stackB_di_nopf_rtt8/
staticB_di_nopf_aux8/
staticB_di_nopf_rtt32/
staticB_di_nopf_rtt8/
staticB_di_rtt8_tcobs/
staticB_di_tcobs_rtt32/
"
for d in $CGOTESTDIRS
do
cp ./testdata/cgoPackage.go ./$d/generated_cgoPackage.go
done

go clean -cache
