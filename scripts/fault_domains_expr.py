#!/usr/bin/env python3

"""
Generates an Azure Resource Manager (ARM) expression that will evaluate at
deployment time to the number of fault domains available in a given location.
Also generates Go code referencing the expression.

Since there is no API to query the fault domain count for an Azure location,
this script parses the public documentation to find out.
"""

import os
import re
import textwrap
import urllib.request


ARM_EXPR = """\
"[
if( contains(
      split('{}', ','),
        variables('location') ),
  3,
if( equals('centraluseuap', variables('location') ),
  1,
  2
))]"\
"""

GO_CODE = """\
// armExpr is evaluated by Azure Resource Manager at deployment time:
//   if location is in the three-fault-domain list, return 3
//   else if location is "canary" (testing), return 1
//   else return 2
// NOTE: use {} to update this ARM expression.
armExpr := `{}`
// strip all whitespace
armExpr = strings.Join(strings.Fields(armExpr), "")
""".format(os.path.basename(__file__), ARM_EXPR)

SOURCE_DOC = 'https://raw.githubusercontent.com/MicrosoftDocs/azure-docs/master/includes/managed-disks-common-fault-domain-region-list.md'  # pylint: disable=line-too-long


def main():
    """
    Print an ARM expression that returns the fault domain count for an Azure
    location, as well as the associated Go code, ready for copy-and-paste.
    """
    regex = re.compile(r"""
    \|\s*             # vertical bar followed by whitespace
    ([A-Za-z0-9 ]*?)  # >= 0 location name characters (lazy, capture)
    \s+\|\s*          # >= 1 whitespace chars, vertical bar, more whitespace
    (\d)              # a single digit (capture)
    \s*\|             # whitespace followed by a vertical bar
    """, re.VERBOSE)

    markdown = urllib.request.urlopen(SOURCE_DOC).read().decode("utf8")
    # Since the canary region is hard-coded, only the regions with three
    # fault domains need to be included in the expression.
    threes = (m[0].replace(' ', '').lower() for m in regex.findall(markdown)
              if int(m[1]) == 3)
    threes = ','.join(sorted(threes))

    print("\n\033[1;36mARM expression minified:\033[0;0m \n")
    print("".join(ARM_EXPR.format(threes).split()))
    print("\n\033[1;34mGo code snippet:\033[0;0m \n")
    print(textwrap.indent(GO_CODE.format(threes), "\t"))


if __name__ == '__main__':
    main()
