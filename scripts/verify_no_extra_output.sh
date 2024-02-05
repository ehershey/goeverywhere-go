#!/bin/bash
#
# Verify that when the executable is called with --completion-script-bash there's nothing extra
# coming out that will break shell startups
# scripts/verify_no_extra_output.sh
#
tempfile=`mktemp`
EXEC=./goe
$EXEC --completion-script-bash 2>$tempfile >/dev/null
if test -s "$tempfile" ; then
  echo "unclean STDERR from $EXEC">&2
  echo "saved to $tempfile"
  exit 1
fi
if ! bash <($EXEC --completion-script-bash) 2>$tempfile
then
  echo "non-zero exit code from bash running completion script">&2
  exit 1
fi

if test -s "$tempfile" ; then
  echo "unclean STDERR from $EXEC">&2
  echo "saved to $tempfile"
  exit 1
fi
rm $tempfile
