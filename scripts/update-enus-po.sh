#!/bin/bash

scripts/update-translation.sh -l en_US -p

mv aksengine.po translations/en_US/LC_MESSAGES/aksengine.po
rm aksengine.pot