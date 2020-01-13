#!/bin/bash

if [ -z $(command -v mockgen) ]; then
    echo "This utility requires mockgen"
    exit 1
fi

TARGET_ALL="*"

REPOS_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

TARGET_REPO="$1"
REPO_FOUND=false

if [ -z "$TARGET_REPO" ]; then
    TARGET_REPO=$TARGET_ALL
    REPO_FOUND=true
    echo "Generating *ALL* Mocks"
    echo ""
fi

ALL_REPOS=(
    "UsersRepo"
    "GlobalRepository"
)

for repo in "${ALL_REPOS[@]}"
do
    if [ "$TARGET_REPO" = "$repo" -o "$TARGET_REPO" == "$TARGET_ALL" ]; then

        REPO_FOUND=true
        echo ""
        echo "Generating mock for $repo"
        dest="$REPOS_ROOT/mocks/$repo.go"
        src="github.com/happilymarrieddad/learning-go-gRPC/repos"
        GO111MODULE=on mockgen -destination=$dest $src $repo

    fi
done

echo ""

if [ "$REPO_FOUND" == "false" ]; then
    echo "Unable to find repo $TARGET_REPO - please make sure it's in all repos or spelled correctly"
    exit 1
fi

echo "Finished generating mocks"
exit 0
