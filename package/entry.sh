#!/bin/bash
set -e

for module in $(find . -name 'go.mod' | sed 's/\/go.mod//'); do
    pushd ${module} 2>/dev/null 1>/dev/null
    go mod download

    popd 2>/dev/null 1>/dev/null
done

root_dir=$(pwd)
for module in $(find . -name 'go.mod' | sed 's/\/go.mod//'); do
    echo "Running tests for MODULE='${module}'..."
    echo ""
    pushd ${module} 2>/dev/null 1>/dev/null
    if [[ -z ${PLUGIN_MODE} ]]; then
        gotestsum -- -count=1 -cover ./...
    else
        gotestsum --jsonfile=${root_dir}/go_test_${module}.json -- -count=1 -cover ./...
    fi
    popd 2>/dev/null 1>/dev/null
    echo ""
done

if [[ -n ${PLUGIN_MODE} ]]; then
    # Ensures output matches Sonobuoy plugin expectations
    # ref: https://sonobuoy.io/docs/v0.51.0/plugins/
    gotestsum --junitfile=report.xml --raw-command -- cat go_test_*.json
    mkdir -p /tmp/results
    echo "$(pwd)/report.xml" > /tmp/results/done
fi

echo "Finished running tests"
echo ""