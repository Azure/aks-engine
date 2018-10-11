#!/usr/bin/env bash

# DON'T RUN. This script requires PythonLocalizerTool which is not published yet.
# TODO: make PythonLocalizerTool public
set -eo pipefail

langdirs=loc/*

convert_lcl_to_po() {
  for dir in $langdirs
  do
    loc_lang=`basename "$dir"`
    translation_lang=`echo $loc_lang | tr - _`
    publish/PythonLocalizerTool lcltopo $dir translations/$translation_lang/LC_MESSAGES/ translations/en_US/LC_MESSAGES/en-US/metadata aksengine ""
    msgfmt -c -v -o translations/$translation_lang/LC_MESSAGES/aksengine.mo translations/$translation_lang/LC_MESSAGES/aksengine.po
  done
}

convert_po_to_lcg() {
  publish/PythonLocalizerTool potolcg translations/en_US/LC_MESSAGES aksengine
}

case "${1-}" in
  --potolcg)
    echo "converting po to lcg..."
    convert_po_to_lcg
    ;;
  --lcltopo)
    echo "converting lcl to po..."
    convert_lcl_to_po
    ;;
  *)
    echo $"Usage: $0 {--potolcg|--lcltopo}"
    exit 1
esac
