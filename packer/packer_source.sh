#!/bin/bash

ERR_PACKER_COPY_FILE=113 # Error writing a file to disk during VHD CI

copyPackerFiles() {
  SYSCTL_CONFIG_SRC=/home/packer/sysctl-d-60-CIS.conf
  SYSCTL_CONFIG_DEST=/etc/sysctl.d/60-CIS.conf
  RSYSLOG_CONFIG_SRC=/home/packer/rsyslog-d-60-CIS.conf
  RSYSLOG_CONFIG_DEST=/etc/rsyslog.d/60-CIS.conf
  ETC_ISSUE_CONFIG_SRC=/home/packer/etc-issue
  ETC_ISSUE_CONFIG_DEST=/etc/issue
  ETC_ISSUE_NET_CONFIG_SRC=/home/packer/etc-issue.net
  ETC_ISSUE_NET_CONFIG_DEST=/etc/issue.net
  SSHD_CONFIG_SRC=/home/packer/sshd_config
  SSHD_CONFIG_DEST=/etc/ssh/sshd_config
  MODPROBE_CIS_SRC=/home/packer/modprobe-CIS.conf
  MODPROBE_CIS_DEST=/etc/modprobe.d/CIS.conf
  PWQUALITY_CONF_SRC=/home/packer/pwquality-CIS.conf
  PWQUALITY_CONF_DEST=/etc/security/pwquality.conf
  PAM_D_SU_SRC=/home/packer/pam-d-su
  PAM_D_SU_DEST=/etc/pam.d/su
  PROFILE_D_CIS_SH_SRC=/home/packer/profile-d-cis.sh
  PROFILE_D_CIS_SH_DEST=/etc/profile.d/CIS.sh
  AUDITD_RULES_SRC=/home/packer/auditd-rules
  AUDITD_RULES_DEST=/etc/audit/rules.d/CIS.rules
  LABEL_NODES_SRC=/home/packer/label-nodes.sh
  LABEL_NODES_DEST=/opt/azure/containers/label-nodes.sh
  LABEL_NODES_SERVICE_SRC=/home/packer/label-nodes.service
  LABEL_NODES_SERVICE_DEST=/etc/systemd/system/label-nodes.service
  CIS_SRC=/home/packer/cis.sh
  CIS_DEST=/opt/azure/containers/provision_cis.sh
  APT_PREFERENCES_SRC=/home/packer/apt-preferences
  APT_PREFERENCES_DEST=/etc/apt/preferences
  KMS_SERVICE_SRC=/home/packer/kms.service
  KMS_SERVICE_DEST=/etc/systemd/system/kms.service
  if [[ ${UBUNTU_RELEASE} == "16.04" ]]; then
    SSHD_CONFIG_SRC=/home/packer/sshd_config_1604
  fi
  DIR=$(dirname "$SYSCTL_CONFIG_DEST") && mkdir -p ${DIR} && cp $SYSCTL_CONFIG_SRC $SYSCTL_CONFIG_DEST && chmod 644 $SYSCTL_CONFIG_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$RSYSLOG_CONFIG_DEST") && mkdir -p ${DIR} && cp $RSYSLOG_CONFIG_SRC $RSYSLOG_CONFIG_DEST && chmod 644 $RSYSLOG_CONFIG_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$ETC_ISSUE_CONFIG_DEST") && mkdir -p ${DIR} && cp $ETC_ISSUE_CONFIG_SRC $ETC_ISSUE_CONFIG_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$ETC_ISSUE_NET_CONFIG_DEST") && mkdir -p ${DIR} && cp $ETC_ISSUE_NET_CONFIG_SRC $ETC_ISSUE_NET_CONFIG_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$SSHD_CONFIG_DEST") && mkdir -p ${DIR} && cp $SSHD_CONFIG_SRC $SSHD_CONFIG_DEST && chmod 644 $SSHD_CONFIG_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$MODPROBE_CIS_DEST") && mkdir -p ${DIR} && cp $MODPROBE_CIS_SRC $MODPROBE_CIS_DEST && chmod 644 $MODPROBE_CIS_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$PWQUALITY_CONF_DEST") && mkdir -p ${DIR} && cp $PWQUALITY_CONF_SRC $PWQUALITY_CONF_DEST && chmod 600 $PWQUALITY_CONF_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$PAM_D_SU_DEST") && mkdir -p ${DIR} && cp $PAM_D_SU_SRC $PAM_D_SU_DEST && chmod 644 $PAM_D_SU_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$PROFILE_D_CIS_SH_DEST") && mkdir -p ${DIR} && cp $PROFILE_D_CIS_SH_SRC $PROFILE_D_CIS_SH_DEST && chmod 755 $PROFILE_D_CIS_SH_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$AUDITD_RULES_DEST") && mkdir -p ${DIR} && cp $AUDITD_RULES_SRC $AUDITD_RULES_DEST && chmod 640 $AUDITD_RULES_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$LABEL_NODES_DEST") && mkdir -p ${DIR} && cp $LABEL_NODES_SRC $LABEL_NODES_DEST && chmod 744 $LABEL_NODES_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$LABEL_NODES_SERVICE_DEST") && mkdir -p ${DIR} && cp $LABEL_NODES_SERVICE_SRC $LABEL_NODES_SERVICE_DEST && chmod 644 $LABEL_NODES_SERVICE_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$CIS_DEST") && mkdir -p ${DIR} && cp $CIS_SRC $CIS_DEST && chmod 744 $CIS_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$APT_PREFERENCES_DEST") && mkdir -p ${DIR} && cp $APT_PREFERENCES_SRC $APT_PREFERENCES_DEST && chmod 644 $APT_PREFERENCES_DEST || exit $ERR_PACKER_COPY_FILE
  DIR=$(dirname "$KMS_SERVICE_DEST") && mkdir -p ${DIR} && cp $KMS_SERVICE_SRC $KMS_SERVICE_DEST && chmod 644 $KMS_SERVICE_DEST || exit $ERR_PACKER_COPY_FILE
}
