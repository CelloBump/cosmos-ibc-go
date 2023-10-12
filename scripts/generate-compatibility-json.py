#!/usr/bin/python3

import argparse
import re
import requests
from typing import List
import semver
import json

FROM_VERSION = "from_version"
RELEASES_URL = "https://api.github.com/repos/cosmos/ibc-go/releases"
CHAIN_A = "chain-a"
CHAIN_B = "chain-b"
HERMES = "hermes"
RLY = "rly"


def parse_args() -> argparse.Namespace:
    """Parse command line arguments."""
    parser = argparse.ArgumentParser(description="Generate Compatibility JSON.")
    parser.add_argument(
        "--file",
        help="The test file to look at.",
    )
    parser.add_argument(
        "--version",
        help="The version to run tests for.",
    )
    parser.add_argument(
        "--release_version",
        help="The release tag",
    )
    parser.add_argument(
        "--chain",
        choices=[CHAIN_A, CHAIN_A],
        default=CHAIN_A,
        help="Specify chain-a or chain-b for use with the json files.",
    )
    parser.add_argument(
        "--relayer",
        choices=[HERMES, RLY],
        default=HERMES,
        help=f"Specify relayer, either {HERMES} or {RLY}",
    )

    return parser.parse_args()


def main():
    args = parse_args()
    file_lines = _load_file_lines(args.file)
    test_suite_name = _extract_test_suite_function(file_lines)
    test_functions = _extract_test_functions_to_run(args.version, file_lines)

    release_versions = [args.release_version]

    tags = _get_ibc_go_releases(args.version)

    other_versions = tags

    # if we are specifying chain B, we invert the versions.
    if args.chain == CHAIN_B:
        release_versions, other_versions = other_versions, release_versions

    compatibility_json = {
        "chain-a": release_versions,
        "chain-b": other_versions,
        "entrypoint": [test_suite_name],
        "test": test_functions,
        "relayer-type": [args.relayer]
    }

    # output the json on a single line. This ensures the output is directly passable to a github workflow.
    print(json.dumps(compatibility_json), end="")


def _get_ibc_go_releases(from_version: str) -> List[str]:
    releases = []

    without_v = from_version
    if without_v.startswith("v"):
        without_v = without_v[1:]
    from_version_semver = semver.Version.parse(without_v)

    resp = requests.get(RELEASES_URL)
    resp.raise_for_status()

    response_body = resp.json()

    all_tags = [release["tag_name"] for release in response_body]
    for tag in all_tags:
        without_v = tag
        if without_v.startswith("v"):
            without_v = without_v[1:]

        # skip alphas, betas and rcs
        if any(c in without_v for c in ["beta", "rc", "alpha"]):
            continue
        try:
            semver_tag = semver.Version.parse(without_v)
        except ValueError: # skip any non semver tags.
            continue
        if semver_tag >= from_version_semver:
            releases.append(tag)

    return releases


def _extract_test_functions_to_run(version: str, file_lines: List[str]) -> List[str]:
    """creates a list of all test functions that should be run in the compatibility tests based on the version provided"""
    test_function_names: List[str] = []
    for i, line in enumerate(file_lines):
        line = line.strip()

        if not line.startswith("//"):
            continue

        if FROM_VERSION in line:
            # TODO: do semver check instead of specific version match.
            if not re.match(fr"//\sfrom_version:\s({version})", line):
                continue

            # TODO: look for the name instead of assuming it's on the next line.
            idx_of_test_declaration = i + 1

            if idx_of_test_declaration >= len(file_lines):
                raise ValueError(
                    "index out of bounds, did not find a function associated with the 'from_version' annotation",
                )

            fn_name_line = file_lines[i + 1].strip()
            test_function_names.append(_extract_function_name_from_line(fn_name_line))

    return test_function_names


def _extract_function_name_from_line(line: str) -> str:
    """extract the name of the go test function from the line of source code provided."""
    return re.search(r".*(Test.*)\(\)", line).group(1)


def _extract_test_suite_function(file_lines: List[str]) -> str:
    """extracts the name of the test suite function in the file. It is assumed there is exactly one test suite defined"""
    for line in file_lines:
        line = line.strip()
        if "(t *testing.T)" in line:
            return re.search(r"func\s+(.*)\(", line).group(1)
    raise ValueError("unable to find test suite in file lines")


def _load_file_lines(file_name: str) -> List[str]:
    with open(file_name, "r") as f:
        return f.readlines()


if __name__ == "__main__":
    main()
