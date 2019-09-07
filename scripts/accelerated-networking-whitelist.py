#!/usr/bin/env python3

import json
import subprocess
from collections import OrderedDict


def get_accelerated_skus():
    """
    Returns a dict of SKU name to AcceleratedNetworkingEnabled value.
    """
    pinned = [
        ("AZAP_Performance_ComputeV17C", True),
        ("SQLGL", True),
        ("SQLGLCore", True),
    ]
    query = r"[? starts_with(name, `Standard`) && !ends_with(name, `Promo`)].{name: name, caps: capabilities[? name==`AcceleratedNetworkingEnabled`]}[].{sku: name, acceleratedNetworking: caps[0].value}[? acceleratedNetworking != null]"
    results = json.loads(
        subprocess.check_output(
            ["az", "vm", "list-skus", "-o", "json", "--query", query]
        ).decode("utf-8")
    )
    return OrderedDict(
        pinned + [(i["sku"], bool(i["acceleratedNetworking"])) for i in results]
    )


if __name__ == "__main__":
    # TODO: print a data structure that can be reused in tests
    print(get_accelerated_skus())
