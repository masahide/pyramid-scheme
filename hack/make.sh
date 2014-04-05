#!/usr/bin/env bash
set -e

set -o pipefail

echo

NAME=pyramid-scheme

DATE=$(date +%y%m%d.%H%M)

VERSION=$(cat ./VERSION)
if command -v git &> /dev/null && git rev-parse &> /dev/null; then
	GITCOMMIT=$(git rev-parse --short HEAD)
	if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
		GITCOMMIT="$GITCOMMIT-dirty"
	fi
elif [ "$HASH" ]; then
	GITCOMMIT="$HASH"
else
	echo >&2 'error: .git directory missing and HASH not specified'
	echo >&2 '  Please either build with the .git directory accessible, or specify the'
	echo >&2 '  exact (--short) commit hash you are building using HASH for'
	echo >&2 '  future accountability in diagnosing build issues.  Thanks!'
	exit 1
fi


if [ ! "$GOPATH" ]; then
	echo >&2 'error: missing GOPATH; please see http://golang.org/doc/code.html#GOPATH'
	exit 1
fi

# Use these flags when compiling the tests and final binary
LDFLAGS='
			-w
			-X github.com/masahide/'$NAME'/version.GITCOMMIT "'$GITCOMMIT'"
			-X github.com/masahide/'$NAME'/version.VERSION "'$VERSION.$DATE'"
'
LDFLAGS_STATIC='-linkmode external'
EXTLDFLAGS_STATIC='-static'

EXTLDFLAGS_STATIC_CUSTOM="$EXTLDFLAGS_STATIC -lpthread -Wl,--unresolved-symbols=ignore-in-object-files"
LDFLAGS_STATIC_CUSTOM="
$LDFLAGS_STATIC
-extldflags \"$EXTLDFLAGS_STATIC_CUSTOM\"
"

HAVE_GO_TEST_COVER=
if \
	go help testflag | grep -- -cover > /dev/null \
	&& go tool -n cover > /dev/null 2>&1 \
	; then
HAVE_GO_TEST_COVER=1
fi


bundle() {
	dir=bin
	echo "---> Making binary: (in bin/ )"
	mkdir -p $dir
	binary $(pwd)/$dir
}


binary() {
	DEST=$1
	echo go build -o $DEST/$NAME -ldflags "$LDFLAGS"
	go build -o $DEST/$NAME -ldflags "$LDFLAGS" &&
		echo $VERSION.$DATE-$GITCOMMIT >$DEST/VERSION &&
		cp $DEST/VERSION $DEST/version-$VERSION.$DATE-$GITCOMMIT.txt
	echo "Created binary: $DEST/$NAME"
}

main() {
	# We want this to fail if the bin already exist and cannot be removed.
	# This is to avoid mixing bin from different versions of the code.
	mkdir -p bin
	if [ -e "bin" ]; then
		echo "bin already exists. Removing."
		rm -fr bin && mkdir bin || exit 1
		echo
	fi
	bundle
	echo
}

main "$@"
