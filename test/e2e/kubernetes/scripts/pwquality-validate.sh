#!/bin/bash

set -x

# an invalid password should fail
echo "tooshort1@" | sudo pwscore && exit 1
echo "password123456@J" | sudo pwscore && exit 1
echo "passSDWword@@@@J" | sudo pwscore && exit 1
echo "passSDWword1111J" | sudo pwscore && exit 1
echo "lowerrrr12case@" | sudo pwscore && exit 1
echo "UPPERRR12CASE@" | sudo pwscore && exit 1

# a valid password should succeed
echo "passSDWword1232rdw#@" | sudo pwscore || exit 1