"""Names each operation's request type, and its parameters' enums, after its
operationId.

    python scripts/name_requests.py build/openapi.sdk.json

Fern names a request type after the method and its group, so retrieveTicker
would take a RetrieveTickersRequest and listCreditRatios a ListCreditRequest.
Named after the operationId, they are RetrieveTickerRequest and
ListCreditRatiosRequest, as the API reference calls the operations. A
parameter's own enum is named the same way: SearchTickersSearchBy, not
SearchTickersRequestSearchBy.
"""

from __future__ import annotations

import json
import sys


def main(path: str) -> None:
    with open(path) as f:
        doc = json.load(f)
    names = set()
    for item in doc["paths"].values():
        for op in item.values():
            if not isinstance(op, dict):
                continue
            operation_id = op.get("operationId")
            if not operation_id:
                sys.exit(f"name_requests: an operation has no operationId: {op.get('summary')}")
            name = operation_id[0].upper() + operation_id[1:] + "Request"
            if name in names:
                sys.exit(f"name_requests: two operations would take a {name}")
            names.add(name)
            op["x-fern-request-name"] = name
            for param in op.get("parameters", []):
                schema = param.get("schema", {})
                for variant in [schema, *schema.get("anyOf", [])]:
                    if isinstance(variant, dict) and "enum" in variant:
                        variant["x-fern-type-name"] = name.removesuffix("Request") + pascal(param["name"])
    with open(path, "w") as f:
        json.dump(doc, f, indent=2)
        f.write("\n")


def pascal(name: str) -> str:
    return "".join(part[:1].upper() + part[1:] for part in name.replace(".", "_").split("_"))


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit(__doc__)
    main(sys.argv[1])
