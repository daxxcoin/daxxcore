#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ethdir="$workspace/src/github.com/daxxcoin"
if [ ! -L "$ethdir/daxxcore" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. daxxcore
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
GO15VENDOREXPERIMENT=1
export GOPATH GO15VENDOREXPERIMENT

# Run the command inside the workspace.
cd "$ethdir/daxxcore"
PWD="$ethdir/daxxcore"

# Launch the arguments with the configured environment.
exec "$@"
