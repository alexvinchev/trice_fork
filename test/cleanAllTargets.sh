#!/usr/bin/env bash

# The file cgoPackage.go is the same in all cgo test packages, but must be inside the folders.
# os agnostic links would be better.
VSCODETARGETDIRS="
Nucleo-STM32L432LC_generated
vsCode_Nucleo-G0B1_generated
vsCode_Nucleo-G0B1_instrumented
vsCode_Nucleo-L432KC_generated
vsCode_Nucleo-L432KC_instrumented
"

for d in $VSCODETARGETDIRS
do
echo ______ $d ______
cd $d
make clean
cd -
echo -e
done

ARMKEILTARGETDIRS="
MDK-ARM_STM32F030R8_generated
MDK-ARM_STM32F030R8_instrumented
"

for d in $ARMKEILTARGETDIRS
do
echo ______ $d ______
cd $d/MDK-ARM
rm -rf RTE/ MDK-ARM_STM32F030R8/ DebugConfig/
cd -
echo -e
done

# create til.json && li.json if not existing
touch ./testdata/til.json ./testdata/li.json

# remove IDs from source code
trice clean -src ../src -src . -i ./testdata/til.json -li ./testdata/li.json -IDMin 1000 -IDMax 7999
