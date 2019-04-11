#!/bin/bash

assignRootPW() {
  grep '^root:[!*]:' /etc/shadow
  if [ $? -eq '0' ] ; then
    SALT=`openssl rand -base64 5`
    SECRET=`openssl rand -base64 37`
    CMD="import crypt, getpass, pwd; print crypt.crypt('$SECRET', '\$6\$$SALT\$')"
    HASH=`python -c "$CMD"`

    echo 'root:'$HASH | /usr/sbin/chpasswd -e 2>/dev/null;
  fi
}
applyCIS() {
  assignRootPW
}