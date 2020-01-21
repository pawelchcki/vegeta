#!/bin/bash
ATTACK_TARGET=${VEGETA_ATTACK_TARGET:-"GET http://localhost:9999"}
ATTACK_RATE=${VEGETA_ATTACK_RATE:-"12000/s"}
REPORT_TYPE=${VEGETA_REPORT_TYPE:-"--type=datadog"}
echo $ATTACK_TARGET | vegeta attack -rate=$ATTACK_RATE | vegeta report $REPORT_TYPE
