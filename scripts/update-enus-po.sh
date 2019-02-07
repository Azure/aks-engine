#!/bin/bash

scripts/update-translation.sh -l en_US -p

mv acsengine.po translations/en_US/LC_MESSAGES/acsengine.po
rm acsengine.pot