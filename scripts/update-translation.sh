#!/bin/bash

GO_SOURCE="pkg/engine/*.go pkg/engine/transform/*.go pkg/api/*.go pkg/operations/*.go pkg/operations/kubernetesupgrade/*.go"
LANGUAGE="en_US"
DOMAIN="aksengine"
generate_po="false"
generate_mo="false"

while getopts "hl:pm" opt; do
  case $opt in
    h)
      echo "$0 [-l language] [-p] [-m]"
      echo " -l <language>: Language to translate"
      echo " -p extract strings and generate PO file"
      echo " -m generate MO file"
      exit 0
      ;;
    l)
      LANGUAGE="${OPTARG}"
      ;;
    p)
      generate_po="true"
      ;;
    m)
      generate_mo="true"
      ;;
    \?)
      echo "$0 [-l language] [-p] [-m]"
      exit 1
      ;;
  esac
done

if ! which go-xgettext > /dev/null; then
  echo 'Can not find go-xgettext, install with:'
  echo 'go get github.com/JiangtianLi/gettext/go-xgettext'
  exit 1
fi

if ! which msginit > /dev/null; then
  echo 'Can not find msginit, install with:'
  echo 'apt-get install gettext'
  exit 1
fi

if [[ "${generate_po}" == "true" ]]; then
  echo "Extract strings and generate PO files..."
  go-xgettext -o ${DOMAIN}.pot --keyword=Translator.Errorf --keyword-plural=Translator.NErrorf --msgid-bugs-address="" --sort-output ${GO_SOURCE}
  msginit -l ${LANGUAGE} -o ${DOMAIN}.po -i ${DOMAIN}.pot
fi

if [[ "${generate_mo}" == "true" ]]; then
  echo "Generate MO file..."
  if [ ! -f ${DOMAIN}.po ]; then
    echo "${DOMAIN}.po not found!"
	exit 1
  fi
  msgfmt -c -v -o ${DOMAIN}.mo ${DOMAIN}.po
fi
